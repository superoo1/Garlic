package util

import "github.com/go-redis/redis/v7"

// 连接池

var RedisDB *redis.Client

func InitRedisDB() (err error) {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = RedisDB.Ping().Result()
	if err != nil {
		panic(err)
	}
	return nil
}

func CloseRedisDB() error {
	return RedisDB.Close()
}
func init() {
	InitRedisDB()
}
