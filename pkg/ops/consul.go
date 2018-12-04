package ops

import (
	"fmt"

	"github.com/cohenjo/dbender/pkg/config"
	"github.com/hashicorp/consul/api"
)

func GetClusterMaster(clusterName string) (string, error) {
	if clusterName == "localhost" {
		return "127.0.0.1", nil
	}
	consulConfig := api.DefaultConfig()
	consulConfig.Address = config.Config.ConsulAddress
	client, err := api.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}
	kv := client.KV()
	// Lookup the pair
	key := fmt.Sprintf("%s/mysql_%s/hostname", config.Config.KVConsulPrefix, clusterName)
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return "failed to get key", err
	}
	return string(pair.Value), nil
}
