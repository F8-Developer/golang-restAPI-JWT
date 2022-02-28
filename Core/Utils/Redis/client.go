package Redis

import (
	"github.com/go-redis/redis"
)

var Client *redis.Client
var Err error