package uniswap

import (
	"testing"

	"github.com/angulo66/subgraph-pool-service/config"
)

func TestUniswap(t *testing.T) {
	t.Log("TestUniswap")
	SetupTest(t)
}

func SetupTest(t *testing.T) {
	t.Log("SetupTest")
	NewUniswapPoolClient(config.PoolConfiguration{IndexedNetwork: "mainnet", QueryURI: "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2", SubgraphID: "uniswap", HTTPRequestTimeout: 10})
}
