package Redis

import (
	log "github.com/Sirupsen/logrus"
	"intrajasa-merchant-api-gateway/Config"
	"github.com/go-redis/redis"
)

// NewClient :
// new connect object for redis.
func NewClient() *redis.Client {
	redis_host := Config.GoDotEnvVariable("REDIS_HOST")
	redis_port := Config.GoDotEnvVariable("REDIS_PORT")
	redis_password := Config.GoDotEnvVariable("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr: redis_host+":"+redis_port,
		Password: redis_password,
		DB: 0,  // use default DB
	})

	pong, err := client.Ping().Result()
	log.Info(pong, err)
	// Output: PONG <nil>
	return client
}
