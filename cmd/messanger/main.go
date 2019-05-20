package main

import (
	"github.com/cohenjo/dbender/pkg/messanger"
	"github.com/cohenjo/dbender/pkg/config"
)



func main() {
	config.LoadConfiguration()
	messanger.Serve()
}