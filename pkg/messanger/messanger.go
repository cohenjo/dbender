
package messanger


//go:generate protoc -I ./messanger --go_out=plugins=grpc:./messanger ./messanger/messanger.proto


import (
	"context"
	"github.com/rs/zerolog/log"
	"net"
	"fmt"

	"google.golang.org/grpc"
	"github.com/nlopes/slack"
	"google.golang.org/grpc/reflection"
	pb "github.com/cohenjo/dbender/pkg/messanger/messanger"
	"github.com/cohenjo/dbender/pkg/config"

)

var (
	messages map[string]string
)


// server is used to implement messanger.GreeterServer.
type server struct{
	slackAPI *slack.Client
}

// SendMessage implements helloworld.GreeterServer
func (s *server) SendMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	var msg slack.MsgOption
	switch in.Type {
	default:
		msg = generateDefaultMessage(ctx,in)
	}
	
	cnl, rts, err := s.slackAPI.PostMessage(in.Channel, msg)
	return &pb.MessageReply{Message: "channel: " + cnl + ", ts: " + rts}, err
}

func Serve() {


	lis, err := net.Listen("tcp", config.Config.Address)
	if err != nil {
		log.Error().Err(err).Msg("failed to listen")
	}
	s := grpc.NewServer()
	api := slack.New(config.Config.SalckToken)
	
	pb.RegisterMessangerServer(s, &server{slackAPI: api})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Error().Err(err).Msg("failed to serve: %v")
	}
}

func registerMessageTemplate(name, template string) {
	if messages == nil {
		messages = make(map[string]string)
	}
	messages[name] = template
}

func generateDefaultMessage(ctx context.Context, in *pb.MessageRequest) slack.MsgOption {
	
	/* EXAMPLE - REPLACE THIS!*/
	headerText := slack.NewTextBlockObject("mrkdwn",in.Msg, false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	bodyText := slack.NewTextBlockObject("mrkdwn", in.Body, false, false)
	bodyAccessory := slack.NewImageBlockElement("https://api.slack.com/img/blocks/bkb_template_images/notifications.png", "calendar thumbnail")
	bodySection := slack.NewSectionBlock(bodyText, nil, slack.NewAccessory(bodyAccessory))

	// Fields
	fieldSlice := make([]*slack.TextBlockObject, 0)
	for k, v := range in.Bag {
		txt := fmt.Sprintf("*%s:*\n%s",k,v)
		fieldSlice = append(fieldSlice, slack.NewTextBlockObject("mrkdwn", txt, false, false))

	}
	fieldsSection := slack.NewSectionBlock(nil, fieldSlice, nil)

	msg := slack.MsgOptionCompose(
		slack.MsgOptionUsername("dbender"),
		slack.MsgOptionBlocks(
			headerSection,
			slack.NewDividerBlock(),
			bodySection,
			slack.NewDividerBlock(),
			fieldsSection,
		),
	)
	return msg
}