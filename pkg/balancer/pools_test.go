package balancer

import (
	"testing"

	"github.com/angulo66/subgraph-pool-service/config"
	math "github.com/angulo66/subgraph-pool-service/pkg/balancer/math"
)

func TestBalancer(t *testing.T) {
	t.Log("TestBalancer")
	SetupTest(t)
}

func SetupTest(t *testing.T) {
	t.Log("SetupTest")
	c := NewBalancerPoolClient(config.PoolConfiguration{IndexedNetwork: "mainnet", QueryURI: "https://api.thegraph.com/subgraphs/name/balancer-labs/balancer-v2", SubgraphID: "balancer", HTTPRequestTimeout: 10})

	test_pool, err := c.GetPoolByID("0x5c6ee304399dbdb9c8ef030ab642b10820db8f56000200000000000000000014")
	if err != nil {
		t.Error(err)
	}
	t.Log("test_pool", test_pool)
	pool := test_pool[0]
	swap := math.GetSwapData(1, &pool, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", "0xba100000625a3754423978a60c9317c58a424e3d")
	t.Log("swap", swap)

	// wp := math.PoolDataToWeightedPoolPairData(&pool, "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", "0xba100000625a3754423978a60c9317c58a424e3d")
	// t.Log("wp", wp)
}
