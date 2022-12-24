package elasticsearch

import (
	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/pkg/errors"
)

type Config struct {
	Url string
}

func NewElasticSearchClient(cfg *Config) (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.Config{
		Addresses: []string{cfg.Url},
	})
	if err != nil {
		return nil, errors.Wrap(err, "pkg.elasticsearch.NewClient")
	}

	_, err = client.Info()
	if err != nil {
		return nil, errors.Wrap(err, "pkg.elasticsearch.client.Info")
	}

	return client, nil
}
