package ops

import (
	"context"
	"fmt"

	"github.com/cohenjo/dbender/pkg/config"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	"github.com/openark/golib/log"
)

// https://docs.mongodb.com/manual/reference/command/currentOp/#dbcmd.currentOp
// CurrentOps return currently running operations in mongod
func CurrentOps(host string, filter string, ns string) error {
	user := config.Config.User
	passwd := config.Config.Password
	clientURI := fmt.Sprintf("mongodb://%s:%s@%s/admin", user, passwd, host)
	client, err := mongo.NewClient(clientURI)
	if err != nil {
		return fmt.Errorf("todo: couldn't connect to mongo: %v", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		return fmt.Errorf("todo: mongo client couldn't connect with background context: %v", err)
	}
	db := client.Database("admin")
	doc := bsonx.Doc{}
	res := db.RunCommand(context.Background(), bson.D{{"currentOp", 1}})

	err = res.Decode(&doc)
	if err != nil {
		// fmt.Println(err)
		log.Error("failed to run cmd... ", err)
	}
	log.Infof("let's check doc: %v \n", doc)

	return nil
}

func CurrentOpDB(host string, db string) error {
	return CurrentOps(host, "", db)
}
