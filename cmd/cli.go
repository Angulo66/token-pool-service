package main

import (
	"encoding/json"
	"fmt"

	cfg "github.com/angulo66/subgraph-pool-service/config"
	rconf "github.com/angulo66/subgraph-pool-service/config/redis"
	bal "github.com/angulo66/subgraph-pool-service/pkg/balancer"
	grp "github.com/angulo66/subgraph-pool-service/pkg/redis-graph"
)

func main() {
	blclient := bal.NewBalancerPoolClient(cfg.PoolConfiguration{IndexedNetwork: "mainnet", QueryURI: "https://api.thegraph.com/subgraphs/name/balancer-labs/balancer-v2", SubgraphID: "balancer", HTTPRequestTimeout: 10})
	graph := grp.NewRedisClient(rconf.RedisConfig{Address: "localhost:6379", Network: "tcp"})
	pools := blclient.GetPoolBase()
	for _, pool := range pools {
		poolMap := map[string]interface{}{"address": pool.Address}
		newpool := pool.Address
		j, _ := json.Marshal(newpool)
		json.Unmarshal(j, &poolMap)
		fmt.Println("newpool", newpool)
		fmt.Println("poolmap", poolMap)
		graph.CreateNode(poolMap)
		fmt.Println("node added!")
	}
	gr := graph.GetGraph()
	gr.Commit()
}
