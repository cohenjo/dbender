package main

import (
	"context"
	"time"
	"google.golang.org/grpc"
	pb "github.com/cohenjo/dbender/pkg/messanger/messanger"
	//"github.com/cohenjo/dbender/pkg/config"
	"flag"

	"github.com/rs/zerolog/log"
)



func main() {

	msg    := flag.String("msg", "hello", "the message to send")
	body   := flag.String("body", "hello world", "the message body to send")
	channel := flag.String("channel", "@cohenjo", "channel to send message to")

	flag.Parse()
	
	//config.Read("conf/bender.conf.json")
	
		// Set up a connection to the server.
		conn, err := grpc.Dial(":50051", grpc.WithInsecure())
		if err != nil {
			log.Error().Err(err).Msg("did not connect")
		}
		defer conn.Close()
	
	c:= pb.NewMessangerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.SendMessage(ctx, &pb.MessageRequest{Msg: *msg, Body: *body, Channel: *channel})
	if err != nil {
		log.Error().Err(err).Msg("could not send message")
	}
	log.Printf("Greeting: %s", r.Message)
}