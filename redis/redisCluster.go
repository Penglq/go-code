package main

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

var ClusterRedisClient *redis.ClusterClient

func incCluster(key string) {
	var err error
	var value int64 = 1
	var storeTime time.Duration = 10 * time.Second

	if ClusterRedisClient.Exists(key).Val() == 1 {
		log.Print("key ", key, "value ", value)
		cmd := ClusterRedisClient.IncrBy(key, value)
		err = cmd.Err()
		if err != nil {
			log.Print("inc", err)
		}
	} else {
		cmd := ClusterRedisClient.Set(key, value, storeTime)
		if err = cmd.Err(); err != nil {
			log.Print("set ", err)
		}
	}
}

func main() {
	ClusterRedisClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{},
		Password: "",
	})
	s := ClusterRedisClient.Get("")
	log.Print(s.Val())

	//incCluster("tests")
}
