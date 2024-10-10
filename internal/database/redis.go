package database

import (
	"log"
	"strconv"

	appConf "github.com/audryus/crispy-octo-system/configs"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis(conf appConf.Redis) {
	i, err := strconv.Atoi(conf.Db)
	if err != nil {
		log.Fatal("Redis erro ao tentar ler database", err)
	}

	Redis = redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password, // no password set
		DB:       i,             // use default DB
		Protocol: 3,             // specify 2 for RESP 2 or 3 for RESP 3
	})
	log.Print("Redis connected")
}
