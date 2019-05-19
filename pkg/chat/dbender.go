package chat

import (
	"bytes"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/cohenjo/dbender/pkg/config"
	"github.com/cohenjo/dbender/pkg/types"
	"github.com/mitchellh/mapstructure"
	"github.com/nlopes/slack"
	"github.com/openark/golib/log"
	wit "github.com/wit-ai/wit-go"
	"golang.org/x/net/context"
)

func GetNewBot() *slackbot.Bot {
	bot := slackbot.New(config.Config.SalckToken)
	return bot
}

func InitBotRoutes(bot *slackbot.Bot) {

	bot.Hear("yes").MessageHandler(YesNoHandler)
	bot.Hear("no").MessageHandler(YesNoHandler)
	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Preprocess(WitPreprocess).Subrouter()
	// toMe.AddMatcher(&IntentMatcher{intent: "default_intent"}).MessageHandler(DefaultIntentHandler)
	toMe.AddMatcher(&IntentMatcher{intent: "replication"}).MessageHandler(ReplicationIntentHandler)
	bot.AddMatcher(&IntentMatcher{intent: "report"}).MessageHandler(ReportHandler)
	// toMe.AddMatcher(&IntentMatcher{intent: "slowness"}).MessageHandler(DefaultIntentHandler)
	// toMe.AddMatcher(&IntentMatcher{intent: "problem_resolution"}).MessageHandler(DefaultIntentHandler)
	// bot.AddMatcher(&IntentMatcher{intent: "locks"}).MessageHandler(LocksIntentHandler)
	toMe.MessageHandler(ConfusedHandler)
	bot.Hear("(?i)who can help(.*)").MessageHandler(HelloHandler)
	bot.Hear("(mongo)|(mysql)").MessageHandler(ClusterHandler)
	// bot.Hear("(?)attachment").MessageHandler(AttachmentsHandler)
	// bot.Hear("(?)file").MessageHandler(ReportHandler)
}

func ReplyToThread(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent, response string) {

	// bot.Reply(msg, "Oh hello!", slackbot.WithTyping)
	if slackbot.WithTyping {
		bot.Type(msg, "Oh hello!")
	}
	newMsg := bot.RTM.NewOutgoingMessage(response, msg.Channel)
	newMsg.ThreadTimestamp = msg.Timestamp
	// newMsg.ThreadBroadcast = false
	bot.RTM.SendMessage(newMsg)
}

func ReplyToThreadWithAttachments(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent, attachments []slack.Attachment, typing bool) {
	params := slack.PostMessageParameters{AsUser: true}
	params.Attachments = attachments
	params.ThreadTimestamp = msg.Timestamp
	bot.Client.PostMessage(msg.Msg.Channel, "", params)
}

func ReplyToThreadWithFile(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent, typing bool, r *bytes.Reader) {
	log.Infof("File upload  \n")
	fileParams := slack.FileUploadParameters{}
	fileParams.Filename = "test.html"
	fileParams.Filetype = "html"
	// fileParams.File = "/tmp/test"

	fileParams.Reader = r
	fileParams.ThreadTimestamp = msg.Timestamp
	fileParams.Channels = append(fileParams.Channels, msg.Msg.Channel)
	uploadedFile, err := bot.Client.UploadFile(fileParams)
	if err != nil {
		log.Errorf("Bot failed to upload file\n, %v", err)
		return
	}
	log.Infof("File uploaded - check permlink: %s\n", uploadedFile.Permalink)

}

func ReplyToThreadWithPDFFile(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent, typing bool, r *bytes.Reader) {
	log.Infof("File upload  \n")
	fileParams := slack.FileUploadParameters{}
	fileParams.Filename = "test.pdf"
	fileParams.Filetype = "pdf"
	// fileParams.File = "/tmp/test"

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Errore(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(true)

	page := wkhtmltopdf.NewPageReader(r)
	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)

	// Add to document
	pdfg.AddPage(page)
	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Errore(err)
	}
	fileParams.Content = string(pdfg.Bytes())

	fileParams.ThreadTimestamp = msg.Timestamp
	fileParams.Channels = append(fileParams.Channels, msg.Msg.Channel)
	uploadedFile, err := bot.Client.UploadFile(fileParams)
	if err != nil {
		log.Errorf("Bot failed to upload file\n", err)
		return
	}
	log.Infof("File uploaded - check permlink: %s\n", uploadedFile.Permalink)

}

func WitPreprocess(ctx context.Context) context.Context {
	msg := slackbot.MessageFromContext(ctx)
	text := slackbot.StripDirectMention(msg.Text)

	session := types.Sessions.GetSession(msg.User)
	log.Info("check session")
	if session.User == "" {
		session.User = msg.User
		types.Sessions.AddSession(session)
		log.Info("new session")
	}
	log.Infof("session with %s\n", session.User)

	if len(text) < 10 || len(text) > 260 {
		log.Debug("message out of size for wit, len %d\n", len(text))
		return ctx
	}
	req := &wit.MessageRequest{Query: text}
	client := wit.NewClient(config.Config.WitAIToken)
	witMessage, err := client.Parse(req)
	if err != nil {
		log.Errorf("failure with Wit: %v\n", err)
		bot := slackbot.BotFromContext(ctx)
		bot.Reply(msg, "Uh oh, I seem to be out of sorts :dizzy_face", slackbot.WithTyping)
		return ctx
	}

	var benderWit types.BenderWit
	mapstructure.Decode(witMessage.Entities, &benderWit)
	for i := range benderWit.Cluster {
		if benderWit.Cluster[i].Confidence > 0.75 {
			session.Cluster = benderWit.Cluster[i].Value
		}
	}
	for i := range benderWit.Intent {
		if benderWit.Intent[i].Confidence > 0.75 {
			session.Intent = benderWit.Intent[i].Value
			log.Infof("session intent: %s\n", session.Intent)
		}
	}
	for i := range benderWit.Sentiment {
		if benderWit.Sentiment[i].Confidence > 0.75 {
			session.Sentiment = benderWit.Sentiment[i].Value
		}
	}

	ctx = AddWitToContext(ctx, witMessage)
	ctx = AddSessionToContext(ctx, &session)
	return ctx
}

type IntentMatcher struct {
	slackbot.RegexpMatcher
	intent     string
	confidence float32
}

func (it *IntentMatcher) Match(ctx context.Context) (bool, context.Context) {
	// default confidence to 50%
	if it.confidence == 0 {
		it.confidence = 0.5
	}
	session := SessionFromContext(ctx)
	// witMessage := WitFromContext(ctx)
	if session != nil {
		log.Debugf("################################################################\n")
		log.Infof("session intent: %s, matcher intent: %s\n", session.Intent, it.intent)
		if session.Intent == it.intent {
			log.Debugf("###########################TRUE#################################\n")
			return true, ctx

		}
	}
	log.Debugf("############################FALSE########################\n")
	return false, ctx
}

func WitFromContext(ctx context.Context) *wit.MessageResponse {
	if result, ok := ctx.Value(types.BenderContextKey("__WIT__")).(*wit.MessageResponse); ok {
		return result
	}
	return nil
}

// AddLoggerToContext sets the logger and returns the newly derived context
func AddWitToContext(ctx context.Context, witMessage *wit.MessageResponse) context.Context {
	k := types.BenderContextKey("__WIT__")
	return context.WithValue(ctx, k, witMessage)
}

func SessionFromContext(ctx context.Context) *types.UserSession {
	if result, ok := ctx.Value(types.BenderContextKey("__SESSION__")).(*types.UserSession); ok {
		return result
	}
	return nil
}

// AddLoggerToContext sets the logger and returns the newly derived context
func AddSessionToContext(ctx context.Context, session *types.UserSession) context.Context {
	k := types.BenderContextKey("__SESSION__")
	return context.WithValue(ctx, k, session)
}
