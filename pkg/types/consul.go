package types

import (
	"github.com/cohenjo/dbender/pkg/config"
	"github.com/hashicorp/consul/api"
	"github.com/outbrain/golib/log"
)

type ConsulClient struct {
	client *api.Client
	config *api.Config
}

func GetConsulClient() ConsulClient {

	consulConfig := api.DefaultConfig()
	if config.Config.ConsulAddress != "" {

		consulConfig.Address = config.Config.ConsulAddress
		// ConsulAclToken defaults to ""
		consulConfig.Token = config.Config.ConsulAclToken
	}
	client, err := api.NewClient(consulConfig)
	if err != nil {
		log.Errore(err)
	}
	ckvs := ConsulClient{
		client: client,
		config: consulConfig,
	}
	return ckvs

}

func (client ConsulClient) Set(key string, value []byte) error {
	log.Infof("storing value to KV, Key: %s\n", key)
	kv := client.client.KV()
	p := &api.KVPair{Key: key, Value: value}
	_, err := kv.Put(p, nil)
	return err

}
func (client ConsulClient) Get(key string) (string, error) {
	kv := client.client.KV()
	// Lookup the pair
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return "", err
	}
	return string(pair.Value), nil
}
func (client ConsulClient) GetCatalog() *api.Catalog {
	return client.client.Catalog()
}
