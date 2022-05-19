package hydra

import (
	"net/url"

	"github.com/ory/hydra-client-go/client"
)

type Config struct {
	AdminURL string `yaml:"adminURL" envconfig:"HYDRA_ADMIN_URL"`
}

func NewClient(cfg *Config) (*client.OryHydra, error) {
	adminURL, err := url.Parse(cfg.AdminURL)
	if err != nil {
		return nil, err
	}

	client := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{Schemes: []string{adminURL.Scheme}, Host: adminURL.Host, BasePath: adminURL.Path})

	return client, nil
}
