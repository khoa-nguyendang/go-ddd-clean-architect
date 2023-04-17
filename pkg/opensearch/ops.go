package opensearch

import (
	config "app/core/configs"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	opensearch "github.com/opensearch-project/opensearch-go/v2"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
}

type OpenSearch struct {
	Client *opensearch.Client
	Index  string
	Alias  string
}

// OpenSearch config
type OpenSearchConfig struct {
	Server   string
	Username string
	Password string
}

func NewConfig(cfg *config.Config) (OpenSearchConfig, error) {
	if cfg == nil {
		return OpenSearchConfig{}, errors.New("invalid configs")
	}
	c := OpenSearchConfig{
		Server:   cfg.OpenSearch.Server,
		Username: cfg.OpenSearch.Username,
		Password: cfg.OpenSearch.Password,
	}
	return c, nil
}

func New(c OpenSearchConfig) (OpenSearch, error) {
	cfg := opensearch.Config{
		Addresses: []string{c.Server},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Username: c.Username,
		Password: c.Password,
		// Retry on 429 TooManyRequests statuses
		RetryOnStatus: []int{502, 503, 504, 429},
		// A simple incremental backoff function
		RetryBackoff: func(i int) time.Duration { return time.Duration(i) * 100 * time.Millisecond },
		// Retry up to 5 attempts
		MaxRetries: 5,
	}

	Client, err := opensearch.NewClient(cfg)
	if err != nil {
		return OpenSearch{}, err
	}

	return OpenSearch{Client: Client}, nil
}

func (e *OpenSearch) CreateIndex(Index string) error {
	e.Index = Index
	e.Alias = Index + "_alias"

	res, err := e.Client.Indices.Exists([]string{e.Index})
	if err != nil {
		return fmt.Errorf("cannot check Index existence: %w", err)
	}
	if res.StatusCode == 200 {
		return nil
	}
	if res.StatusCode != 404 {
		return fmt.Errorf("error in Index existence response: %s", res.String())
	}

	res, err = e.Client.Indices.Create(e.Index)
	if err != nil {
		return fmt.Errorf("cannot create Index: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("error in Index creation response: %s", res.String())
	}

	res, err = e.Client.Indices.PutAlias([]string{e.Index}, e.Alias)
	if err != nil {
		return fmt.Errorf("cannot create Index Alias: %w", err)
	}
	if res.IsError() {
		return fmt.Errorf("error in Index Alias creation response: %s", res.String())
	}

	return nil
}
