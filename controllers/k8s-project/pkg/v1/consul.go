package v1

import (
	"fmt"
	"strings"

	consulAPI "github.com/hashicorp/consul/api"
)

type ConsulSource struct {
	Host   string `json:"bucket"`
	Prefix string `json:"prefix"`
	Object string `json:"object"`
}

func NewConsulSource(config SourceConfig) (Source, error) {
	src := &ConsulSource{
		Host:   config.BasePath,
		Prefix: config.Prefix,
		Object: config.Object,
	}

	return src, nil
}

func (c *ConsulSource) GetData() (map[string]string, error) {

	client, err := consulAPI.NewClient(consulAPI.DefaultConfig())

	if err != nil {
		return nil, err
	}

	// Get a handle to the KV API
	kv := client.KV()

	m := make(map[string]string)

	if c.Object != "" {
		pair, _, err := kv.Get(fmt.Sprintf("%s/%s", c.Prefix, c.Object), nil)

		if err != nil {
			return nil, err
		}

		fmt.Printf("KV: %v\n", pair)

		m[c.Object] = string(pair.Value)
	} else {

		pairs, _, err := kv.List(c.Prefix, nil)

		if err != nil {
			return nil, err
		}

		for _, p := range pairs {
			// <BASE>_<KEY>:
			k := strings.Replace(p.Key, "/", "_", -1)
			m[k] = string(p.Value)
		}
	}

	return m, nil
}
