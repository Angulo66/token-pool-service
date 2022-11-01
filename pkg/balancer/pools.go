package balancer

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/angulo66/subgraph-pool-service/config"
	etypes "github.com/angulo66/subgraph-pool-service/pkg/balancer/types"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type BalancerPoolClient struct {
	logger     zerolog.Logger
	cfg        config.PoolConfiguration
	wg         *sync.WaitGroup
	poolbase   *[]etypes.SubgraphPoolBase
	httpClient *http.Client
}

func NewBalancerPoolClient(cfg config.PoolConfiguration) *BalancerPoolClient {
	b := &BalancerPoolClient{
		logger:     log.With().Str("module", "balancer").Logger(),
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

func (c *BalancerPoolClient) GetConfig() config.PoolConfiguration {
	return c.cfg
}

func (c *BalancerPoolClient) getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(c.cfg.HTTPRequestTimeout)*time.Second)
}

func (c *BalancerPoolClient) GetPools() ([]etypes.SubgraphPoolBase, error) {
	ctx, cancel := c.getContext()

	query := balancerV2PoolsQuery()
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

	res := result["data"].(map[string]interface{})["pools"].([]interface{})
	var pools []etypes.SubgraphPoolBase

	b, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, &pools)
	return pools, nil
}

// TODO: remove this
func balancerV2PoolsQuery() map[string]string {
	return map[string]string{
		"query": `{
				pools(where: { totalLiquidity_gt: 0 }) {
				  id
				  address
				  poolType
				  swapFee
				  swapEnabled
				  totalShares
				  tokens {
					address
					name
					symbol
					decimals
					balance
				  }
				  tokensList
				  totalWeight
				  amp
				  expiryTime
				  unitSeconds
				  principalToken
				  baseToken
				  mainIndex
				  wrappedIndex
				  lowerTarget
				  upperTarget
			  }
			}`,
	}
}
