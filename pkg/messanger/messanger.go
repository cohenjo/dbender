
package messanger


//go:generate protoc -I ./messanger --go_out=plugins=grpc:./messanger ./messanger/messanger.proto


import (
	"context"
	"github.com/rs/zerolog/log"
	"net"

	"google.golang.org/grpc"
	"github.com/nlopes/slack"
	"google.golang.org/grpc/reflection"
	pb "github.com/cohenjo/dbender/pkg/messanger/messanger"
	"github.com/cohenjo/dbender/pkg/config"

)

const (
	port = ":50051"
)


// server is used to implement messanger.GreeterServer.
type server struct{
	slackAPI *slack.Client
}

// SendMessage implements helloworld.GreeterServer
func (s *server) SendMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	
	log.Info().Msgf("Received: %v", in.Msg)
	
	attachment := slack.Attachment{
		Pretext: in.Msg,
		Text:    in.Body,
		Color: "danger",
		AuthorName: "Lonely Sheep",
		
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "fears",
				Value: "BIG WOLF",
			},
			slack.AttachmentField{
				Title: "severity",
				Value: "Urgent",
			},
			slack.AttachmentField{
				Title: "Host",
				Value: "lonely.sheep.big.wolf.dinner",
			},
		},
		MarkdownIn: []string{"*This is Bold*"},
	}

	mso := slack.MsgOptionCompose(
		// slack.MsgOptionText(in.Msg, false),
		slack.MsgOptionUsername("dbender"),
		// slack.MsgOptionIconEmoji("alert"),
		slack.MsgOptionAttachments(attachment),
	)
	s.slackAPI.PostMessage(in.Channel, mso)
	return &pb.MessageReply{Message: "Hello " + in.Msg}, nil
}

func Serve() {
	lis, err := net.Listen("tcp", port)
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

