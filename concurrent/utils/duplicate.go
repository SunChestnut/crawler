package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

func DuplicateWithRedis() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.56.10:6379",
		Password: "",
		DB:       0,
	})

	key := "GoGoGo"
	value := "WaoWaoWao~"

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	val1, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	log.Printf("key: %s, value: %v\n", key, val1)

	val2, err := rdb.Get(ctx, "Coffee").Result()
	if err == redis.Nil {
		log.Fatalln("key coffee does not exist")
	} else if err != nil {
		panic(err)
	} else {
		log.Printf("key: Coffee, value: %v\n", val2)
	}
}
