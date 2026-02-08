// package utils

// import (
// 	"context"
// 	"log"
// 	"strconv"
// 	"time"

// 	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
// 	"github.com/redis/go-redis/v9"
// )

// var (
// 	RedisClient *redis.Client
// 	ctx = context.Background()
// )

// func ConnectRedis()  {
// 	db, err := strconv.Atoi(config.AppConfig.REDIS_DB)
// 	if err != nil {
// 		log.Fatalf("couldn't convert redis db in int")
// 	}

// 	// create redis client
// 	RedisClient = redis.NewClient(&redis.Options{
// 		Addr: config.AppConfig.REDIS_HOST,
// 		Password: config.AppConfig.REDIS_PASS,
// 		DB:  db,
// 	})

// 	// ping redis
// 	pingCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

// 	if err := RedisClient.Ping(pingCtx).Err(); err != nil {
// 		log.Fatalf("Couldn't ping redis")
// 	}

// 	log.Printf("Redis Connected✅")


// }

// func RedisSetKey(key string, value string, ttl time.Duration) error {
// 	return  RedisClient.Set(ctx, key, value, ttl).Err()
// }

// func RedisGetKey(key string) (string, error) {
// 	return RedisClient.Get(ctx, key).Result()
// }
// func RedisDelKey(key string) error {
// 	return RedisClient.Del(ctx, key).Err()
// }

package utils

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	ctx         = context.Background()
)

func ConnectRedis()  {
	db, err := strconv.Atoi(config.AppConfig.REDIS_DB)
	if err != nil {
		log.Fatalf("invalid datatype of redis db")
	}

	// redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.REDIS_HOST,
		Password: config.AppConfig.REDIS_PASS,
		DB:       db,
	})

	// ping redis
	pingCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := RedisClient.Ping(pingCtx).Err(); err != nil {
		log.Fatalf("Couldn't ping redis")
	}

	log.Printf("Redis Connected✅")
	
}

func RedisSetKey(key string, value string, ttl time.Duration) error {
	return RedisClient.Set(ctx, key, value, ttl).Err()
}

func RedisGetKey(key string) (string, error){
	return  RedisClient.Get(ctx, key).Result()
}

func RedisDelKey(key string )error {
	return RedisClient.Del(ctx, key).Err()
}
