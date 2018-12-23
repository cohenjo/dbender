package main

import (
	"flag"

	"github.com/openark/golib/log"

	"github.com/cohenjo/dbender/pkg/config"
	"github.com/cohenjo/dbender/pkg/ops"
)

func main() {

	configFile := flag.String("config", "", "config file name")
	flag.Parse()

	if len(*configFile) > 0 {
		config.Read(*configFile)
	} else {
		config.Read("/etc/bender.conf.json")
	}

	err := ops.CurrentOps("db-mongo-others0a.42.wixprod.net")
	if err != nil {
		log.Error("failed to get mongo ops", err)
	}

	// masterHost, err := ops.GetClusterMaster("localhost")
	// if err != nil {
	// 	log.Error("failed to get key", err)
	// }
	// // ops.Report(masterHost)
	// locks, err := ops.CheckLocks(masterHost)
	// if err != nil {
	// 	log.Error("failed to get locks", err)
	// }

	// // Get column names
	// for _, lock := range locks {
	// 	fmt.Printf("lock: %s\n", lock)
	// }

}
