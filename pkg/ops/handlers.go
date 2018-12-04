package ops

import (
	"bytes"
	"fmt"
	"io"
	"time"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
	"github.com/openark/golib/log"
	"golang.org/x/net/context"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func ReplicationIntentHandler(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent) {
	log.Infof("default replication are you handler \n")
	log.Infof("################################################################### \n")

	log.Infof("################################################################### \n")
	bot.Reply(msg, "A bit tired. You get it? A bit?", slackbot.WithTyping)

}

func LocksIntentHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {

	// masterHost, err := GetClusterMaster("billing")
	masterHost, err := GetClusterMaster("localhost")
	if err != nil {
		log.Error("failed to get key", err)
	}
	locks, err := CheckLocks(masterHost)
	if err != nil {
		log.Error("failed to get locks", err)
	}

	txt := ""
	// Get column names
	for _, lock := range locks {
		txt += fmt.Sprintf("lock: %s\n", lock)
	}

	attachment := slack.Attachment{
		Pretext: "We bring you locks. :sunglasses: :lock:",
		Title:   "Locks we found on: " + masterHost,
		// TitleLink: "https://beepboophq.com/",
		Text:     txt,
		Fallback: txt,
		ImageURL: "https://storage.googleapis.com/beepboophq/_assets/bot-1.22f6fb.png",
		Color:    "#7CD197",
	}

	// supports multiple attachments
	attachments := []slack.Attachment{attachment}
	ReplyToThreadWithAttachments(ctx, bot, evt, attachments, slackbot.WithTyping)
}

func HelloHandler(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent) {

	// bot.Reply(msg, "Oh hello!", slackbot.WithTyping)
	// if slackbot.WithTyping {
	// 	bot.Type(msg, "Oh hello!")
	// }

	ReplyToThread(ctx, bot, msg, "hello to you")
}

func HowAreYouHandler(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent) {
	log.Infof("how are you handler \n")
	witMessage := WitFromContext(ctx)
	if witMessage != nil {
		log.Infof("we got from context : %v \n", witMessage)
	}
	bot.Reply(msg, "A bit tired. You get it? A bit?", slackbot.WithTyping)
}

func AttachmentsHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	txt := "Beep Beep Boop is a ridiculously simple hosting platform for your Slackbots."
	attachment := slack.Attachment{
		Pretext:   "We bring bots to life. :sunglasses: :thumbsup:",
		Title:     "Host, deploy and share your bot in seconds.",
		TitleLink: "https://beepboophq.com/",
		Text:      txt,
		Fallback:  txt,
		ImageURL:  "https://storage.googleapis.com/beepboophq/_assets/bot-1.22f6fb.png",
		Color:     "#7CD197",
	}

	// supports multiple attachments
	attachments := []slack.Attachment{attachment}
	ReplyToThreadWithAttachments(ctx, bot, evt, attachments, slackbot.WithTyping)
	// bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
}

func ReportHandler(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent) {

	// if we are in a Handler the session already exists.
	session := SessionFromContext(ctx)
	log.Infof("sessions intent: %s, cluster: %s\n", session.Intent, session.Cluster)
	if session.Cluster == "" {
		bot.Reply(msg, "What do you want meatbag? which cluster?", slackbot.WithTyping)
	} else {
		// types.Sessions.
		log.Infof("File handler  \n")
		masterHost, err := GetClusterMaster(session.Cluster)
		if err != nil {
			log.Error("failed to get key", err)
			bot.Reply(msg, "What do you want meatbag? which cluster?", slackbot.WithTyping)
		}
		rr, w := io.Pipe()
		log.Infof("File handler - send to report \n")
		go Report(masterHost, w)
		log.Infof("File handler - send to upload \n")
		buf := new(bytes.Buffer)
		buf.ReadFrom(rr)
		output := blackfriday.Run(buf.Bytes())
		r := bytes.NewReader(output)
		ReplyToThreadWithFile(ctx, bot, msg, slackbot.WithTyping, r)
		// ReplyToThreadWithPDFFile(ctx, bot, msg, slackbot.WithTyping, r)
		rr.Close()
	}
	session.Updated = time.Now()
}

func ConfusedHandler(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent) {
	bot.Reply(msg, "I don't understand 😰", slackbot.WithTyping)
}
