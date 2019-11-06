package v1

import (
	"fmt"
)

var (
	NewSource = new
)

// +kubebuilder:validation:Enum=S3;Consul
type SourceType string

// Source defines the contract for config source.
type Source interface {
	GetData() (map[string]string, error)
}

type SourceConfig struct {

	// +comment BasePath url/host to be used as base
	BasePath string `json:"basePath"`

	// +comment Prefix is the uri base path to the object
	Prefix string `json:"prefix"`

	// +comment File/object name
	// +optional
	Object string `json:"object"`
}

func new(t SourceType, c SourceConfig) (Source, error) {

	var s Source
	var err error
	switch t {
	case "S3":
		s, err = NewS3Source(c)
		break
	case "Consul":
		s, err = NewConsulSource(c)
		break
	default:
		err = fmt.Errorf("no such source type")
	}

	return s, err
}
