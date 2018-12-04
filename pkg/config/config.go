package config

import (
	"encoding/json"
	"os"

	"github.com/openark/golib/log"
)

type Configuration struct {
	Debug          bool
	SalckToken     string
	WitAIToken     string
	ConsulAddress  string // Address where Consul HTTP api is found. Example: 127.0.0.1:8500
	ConsulAclToken string // ACL token used to write to Consul KV
	KVConsulPrefix string // Prefix to use for clusters' masters entries in KV stores (internal, consul, ZK), default: "mysql/master"
	User           string
	Password       string
}

// Config is *the* configuration instance, used globally to get configuration data
var Config = newConfiguration()

func newConfiguration() *Configuration {
	return &Configuration{
		Debug:          false,
		ConsulAddress:  "dbmng-shepherd0a.42.wixprod.net:8500",
		ConsulAclToken: "",
		KVConsulPrefix: "db/mysql/master",
	}
}

// Read reads configuration from given file, or silently skips if the file does not exist.
// If the file does exist, then it is expected to be in valid JSON format or the function bails out.
func Read(fileName string) (*Configuration, error) {
	file, err := os.Open(fileName)
	if err == nil {
		decoder := json.NewDecoder(file)
		err := decoder.Decode(Config)
		if err == nil {
			log.Infof("Read config: %s", fileName)
		} else {
			log.Fatal("Cannot read config file:", fileName, err)
		}
	}
	return Config, err
}
