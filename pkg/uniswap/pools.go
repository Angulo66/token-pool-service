package uniswap

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

type UniswapPoolClient struct {
	logger     zerolog.Logger
	cfg        config.PoolConfiguration
	wg         *sync.WaitGroup
	poolbase   *[]etypes.Pool
	httpClient *http.Client
}

func NewUniswapPoolClient(cfg config.PoolConfiguration) *UniswapPoolClient {
	b := &UniswapPoolClient{
		logger:     log.With().Str("module", "uniswap").Logger(),
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

func (c *UniswapPoolClient) GetConfig() config.PoolConfiguration {
	return c.cfg
}

func (c *UniswapPoolClient) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(c.cfg.HTTPRequestTimeout)*time.Second)
}

func (c *UniswapPoolClient) GetPools() ([]etypes.Pool, error) {
	ctx, cancel := c.getContext()

	query := uniswapV2PoolsQuery()
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

func uniswapV2PoolsQuery() map[string]string {
	return map[string]string{
		"query": `{
			pairs(where: { trackedReserveETH_gt: 0 }) {
			  id
			  token0 {
				id
				name
				symbol
				decimals
			  }
			  token1 {
				id
				name
				symbol
				decimals
			  }
			  trackedReserveETH
			  reserveETH
			  totalSupply
			  reserveUSD
			  reserve0
			  reserve1
			}
		  }`,
	}
}
