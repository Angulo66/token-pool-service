package graph

import (
	"testing"

	config "github.com/angulo66/subgraph-pool-service/config/redis"
)

func TestRedisGraph(t *testing.T) {
	t.Log("TestRedisGraph")
	SetupTest(t)
}

func SetupTest(t *testing.T) {
	t.Log("SetupTest")
	NewRedisClient(config.RedisConfig{Address: "localhost:6379", Network: "tcp"})
}
