package sushiswap

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/angulo66/subgraph-pool-service/config"
	etypes "github.com/angulo66/subgraph-pool-service/pkg/uniswap/types"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type SushiswapPoolClient struct {
	logger     zerolog.Logger
	cfg        config.PoolConfiguration
	wg         *sync.WaitGroup
	poolbase   *[]etypes.Pool
	httpClient *http.Client
}

func NewSushiswapPoolClient(cfg config.PoolConfiguration) *SushiswapPoolClient {
	b := &SushiswapPoolClient{
		logger:     log.With().Str("module", "sushiswap").Logger(),
		cfg:        cfg,
		wg:         &sync.WaitGroup{},
		httpClient: &http.Client{Timeout: time.Duration(cfg.HTTPRequestTimeout) * time.Second},
	}
	pools, err := b.GetPools()
	if err != nil {
		b.logger.Error().Err(err).Msg("Error getting pools")
	}
	b.poolbase = &pools
	return b
}

func (c *SushiswapPoolClient) GetConfig() config.PoolConfiguration {
	return c.cfg
}

func (c *SushiswapPoolClient) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(c.cfg.HTTPRequestTimeout)*time.Second)
}

func (c *SushiswapPoolClient) GetPools() ([]etypes.Pool, error) {
	ctx, cancel := c.getContext()

	query := sushiswapPoolsQuery()
	jsonValue, _ := json.Marshal(query)
	defer cancel()
	request, err := http.NewRequestWithContext(ctx, "POST", c.cfg.QueryURI, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	res := result["data"].(map[string]interface{})["pairs"].([]interface{})
	var pools []etypes.Pool

	b, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, &pools)
	return pools, nil
}

func sushiswapPoolsQuery() map[string]string {
	return map[string]string{
		"query": `{
			pairs(where:{trackedReserveETH_gt:0}) {
			  id
			  name
			  token0 {
				id
				symbol
				name
			  }
			  token1 {
				id
				symbol
				name
			  }
			  reserve0
			  reserve1
			  trackedReserveETH
			  token0Price
			  token1Price
			}
		  }`,
	}
}
