// // package utils

// // import (
// // 	"context"
// // 	"log"
// // 	"strconv"
// // 	"time"

// // 	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
// // 	"github.com/redis/go-redis/v9"
// // )

// // var (
// // 	RedisClient *redis.Client
// // 	Ctx         = context.Background()
// // )

// // func ConnectRedis() {
// // 	db, err := strconv.Atoi(config.AppConfig.Redis_DB)
// // 	if err != nil {
// // 		log.Fatal("REDIS_DB must be number")
// // 	}

// // 	RedisClient = redis.NewClient(&redis.Options{
// // 		Addr:     config.AppConfig.Redis_Host,
// // 		Password: config.AppConfig.Redis_Pass,
// // 		DB:       db,
// // 	})

// // 	if err := RedisClient.Ping(Ctx).Err(); err != nil {
// // 		log.Fatalf("Redis connect failed: %v", err)
// // 	}

// // 	log.Println("Redis connected ✅")
// // }

// // // SET
// // func RedisSet(key string, value string, ttl time.Duration) error {
// // 	return RedisClient.Set(Ctx, key, value, ttl).Err()
// // }

// // // GET
// // func RedisGet(key string) (string, error) {
// // 	return RedisClient.Get(Ctx, key).Result()
// // }

// // // DELETE
// //
// //	func RedisDel(key string) error {
// //		return RedisClient.Del(Ctx, key).Err()
// //	}

// package utils

// import (
// 	"context"
// 	"log"
// 	"strconv"
// 	"time"

// 	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
// 	"github.com/redis/go-redis/v9"
// )

// var (
// 	RedisClient *redis.Client
// 	ctx = context.Background()
// )

// func ConnectRedis(){
// 	db, err := strconv.Atoi(config.AppConfig.Redis_DB)
// 	if err != nil {
// 		log.Fatalf("Err %s", err)
// 		return
// 	}
// 	RedisClient = redis.NewClient(&redis.Options{
// 		Addr: config.AppConfig.Redis_Host,
// 		Password: config.AppConfig.Redis_Pass,
// 		DB: db,
// 	})

// 	// ping redis
// 	if err := RedisClient.Ping(ctx).Err(); err != nil {
// 		log.Fatalf("err %s", err)
// 		return
// 	}

// 	log.Printf("Redis Connected✅")
// }

// func RedisSetKey(key string, value string, ttl time.Duration) error {
// 	return  RedisClient.Set(ctx, key, value, ttl).Err()
// }

// func RedisGetKey(key string) (string, error) {
// 	return  RedisClient.Get(ctx, key).Result()
// }

// func RedisDelKey(key string) error {
// 	return  RedisClient.Del(ctx, key).Err()
// }

package utils

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/AbdulRahman-04/FullStackWebDev_2026/08_Go_Gin_PostgreSQL/server/internal/config"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	ctx = context.Background()
)

func ConnectRedis() {
  
	// connect to redis 
	db, err := strconv.Atoi(config.AppConfig.Redis_DB)
	if err != nil {
		log.Println(err)
		return
	}

	// make redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr: config.AppConfig.Redis_Host,
		Password: config.AppConfig.Redis_Pass,
		DB: db,
	})

	// ping 
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Println(err)
		return
	}

	log.Println("Redis Connected✅")


}

func RedisSetKey(key string, value string, ttl time.Duration) error {
	return RedisClient.Set(ctx, key, value, ttl).Err()
}

func RedisGetKey(key string) (string, error){
	return  RedisClient.Get(ctx, key).Result()
}
func RedisDelKey(key string) error {
	return RedisClient.Del(ctx, key).Err()
}