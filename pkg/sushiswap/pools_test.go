package sushiswap

import (
	"testing"

	"github.com/angulo66/subgraph-pool-service/config"
)

func TestSushiSwap(t *testing.T) {
	t.Log("TestSushiSwap")
	SetupTest(t)
}

func SetupTest(t *testing.T) {
	t.Log("SetupTest")
	NewSushiswapPoolClient(config.PoolConfiguration{IndexedNetwork: "mainnet", QueryURI: "https://api.thegraph.com/subgraphs/name/sushiswap/exchange", SubgraphID: "sushiswap", HTTPRequestTimeout: 10})
}
