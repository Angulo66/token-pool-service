package config

type RedisConfig struct {
	Network string `mapstructure:"network"`
	Address string `mapstructure:"address"`
}
