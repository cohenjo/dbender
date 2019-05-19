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

	channel := flag.String("channel", "@cohenjo", "channel to send message to")

	flag.Parse()
	
	//config.Read("conf/bender.conf.json")
	
	bag := make(map[string]string)
	bag["Type"] = "Computer"
	bag["Reason"] = "All vowel keys aren't working."

	msg :=  "You have a new request:\n*<google.com|Fred Enriquez - Time Off request>*"
	body := "*<fakeLink.toUserProfiles.com|Iris / Zelda 1-1>*\nTuesday, January 21 4:00-4:30pm\nBuilding 2 - Havarti Cheese (3)\n2 guests"

	// Set up a connection to the server.
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Error().Err(err).Msg("did not connect")
	}
	defer conn.Close()
	
	c:= pb.NewMessangerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := c.SendMessage(ctx, &pb.MessageRequest{Msg: msg, Body: body, Channel: *channel, Bag: bag})
	if err != nil {
		log.Error().Err(err).Msg("could not send message")
	}
	log.Printf("Greeting: %s", r.Message)
}