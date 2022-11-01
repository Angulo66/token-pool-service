package balancer

import (
	"testing"

	"github.com/angulo66/subgraph-pool-service/config"
)

func TestBalancer(t *testing.T) {
	t.Log("TestBalancer")
	SetupTest(t)
}

func SetupTest(t *testing.T) {
	t.Log("SetupTest")
	NewBalancerPoolClient(config.PoolConfiguration{IndexedNetwork: "mainnet", QueryURI: "https://api.thegraph.com/subgraphs/name/balancer-labs/balancer-v2", SubgraphID: "balancer", HTTPRequestTimeout: 10})
}
