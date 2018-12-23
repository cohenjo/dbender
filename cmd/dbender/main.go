package main

import (
	"flag"

	"github.com/cohenjo/dbender/pkg/chat"
	"github.com/cohenjo/dbender/pkg/config"
	"github.com/openark/golib/log"
)

func main() {

	configFile := flag.String("config", "", "config file name")
	flag.Parse()

	if len(*configFile) > 0 {
		config.Read(*configFile)
	} else {
		config.Read("/etc/bender.conf.json")
	}
	log.Infof("slack token is: %s, WIT is: %s", config.Config.SalckToken, config.Config.WitAIToken)

	bot := chat.GetNewBot()
	chat.InitBotRoutes(bot)

	bot.Run()
}
