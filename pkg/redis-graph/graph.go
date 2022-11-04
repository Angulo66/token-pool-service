package graph

import (
	"fmt"

	config "github.com/angulo66/subgraph-pool-service/config/redis"
	"github.com/gomodule/redigo/redis"
	rg "github.com/redislabs/redisgraph-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type RedisGraphClient struct {
	conn   redis.Conn
	cfg    config.RedisConfig
	logger zerolog.Logger
	graph  rg.Graph
}

func NewRedisClient(cfg config.RedisConfig) *RedisGraphClient {
	conn, err := redis.Dial(cfg.Network, cfg.Address)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Redis")
	}

	c := &RedisGraphClient{
		conn:   conn,
		cfg:    cfg,
		logger: log.With().Str("module", "redis-graph").Logger(),
	}

	graph := c.CreateGraph()
	c.graph = graph

	return c
}

func (c *RedisGraphClient) GetConfig() config.RedisConfig {
	return c.cfg
}

func (c *RedisGraphClient) CreateGraph() rg.Graph {
	g := rg.GraphNew("graph", c.conn)
	_, err := g.Commit()
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to create graph")
	}
	return g
}

func (c *RedisGraphClient) GetGraph() rg.Graph {
	return c.graph
}

func (c *RedisGraphClient) CreateNode(props map[string]interface{}) {

	fmt.Println("props", props)

	poolNode := rg.NodeNew("pool", "pool", props)
	c.graph.AddNode(poolNode)
	fmt.Println("node added!")
	//c.graph.Commit()
	fmt.Println("graph commited!")
}

// func (c *RedisGraphClient) Query(query string) (interface{}, error) {
// 	c.logger.Info().Msg("Querying redis graph")
// 	return c.conn.Do("GRAPH.QUERY", c.cfg.GraphName, query)
// }
