package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
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
	Address        string
}

// Config is *the* configuration instance, used globally to get configuration data
var Config *Configuration

// LoadConfiguration loads configuration using viper
func LoadConfiguration() *Configuration {

	viper.SetDefault("Debug", true)
	viper.SetDefault("Address", ":50051")

	viper.SetConfigName("messanger")        // name of config file (without extension)
	viper.AddConfigPath("/etc/")            // path to look for the config file in
	viper.AddConfigPath("$HOME/.messanger") // call multiple times to add many search paths
	viper.AddConfigPath("./conf")           // optionally look for config in the working directory
	err := viper.ReadInConfig()             // Find and read the config file
	if err != nil {                         // Handle errors reading the config file
		log.Error().Err(err).Msg("Fatal error config file")
	}

	viper.WatchConfig()
	viper.OnConfigChange(reloadConfig)
	var cfg Configuration
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Error().Err(err).Msg("unable to decode into struct")
	}

	log.Debug().Msgf("configuration loaded: %+v", cfg)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	Config = &cfg
	return &cfg
}

func reloadConfig(e fsnotify.Event) {
	log.Info().Msgf("Config file changed: %v", e.Name)
	var cfg Configuration
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Error().Err(err).Msg("unable to decode into struct")
	}
	Config = &cfg
}
