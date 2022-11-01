package config

import "github.com/rs/zerolog/log"

func (c *PoolConfiguration) Validate() {
	if c.IndexedNetwork == "" {
		log.Fatal().Str("field", "indexedNetwork").Msg("missing required field")
	}
	if c.QueryURI == "" {
		log.Fatal().Str("field", "queryURI").Msg("missing required field")
	}
	if c.SubgraphID == "" {
		log.Fatal().Str("field", "subgraphID").Msg("missing required field")
	}
	if c.IndexedNetwork != "mainnet" && c.IndexedNetwork != "xdai" {
		log.Fatal().Str("field", "indexedNetwork").Msg("invalid value")
	}
	if c.HTTPRequestTimeout == 0 {
		log.Debug().Msg("using default HTTP request timeout")
		c.HTTPRequestTimeout = 10
	}
}

type PoolConfiguration struct {
	IndexedNetwork     string `mapstructure:"indexed_network"`      // The network that the subgraph is indexed on
	QueryURI           string `mapstructure:"query_uri"`            // The URI to query the subgraph
	SubgraphID         string `mapstructure:"subgraph_id"`          // The ID of the subgraph to query
	HTTPRequestTimeout int    `mapstructure:"http_request_timeout"` // The timeout for HTTP requests to the subgraph
}
