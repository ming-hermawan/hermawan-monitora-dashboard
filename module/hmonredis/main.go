package hmonredis

import (
    "context"
    "fmt"
    "strconv"
    "time"
    "github.com/go-redis/redis/v8"
    "hermawan-monitora/module/hmonenv"
)

var redisClient *redis.Client
var ctx context.Context

func getRedisClient() *redis.Client {
    var redisPort int
    var redisDB int
    var err error
    redisPort, err = strconv.Atoi(hmonenv.Get("REDIS_PORT"))
    if err != nil {
        panic(fmt.Sprintf("REDIS_PORT ENV error, %s", redisPort))
    }
    redisDB, err = strconv.Atoi(hmonenv.Get("REDIS_DB"))
    if err != nil {
        panic(fmt.Sprintf("REDIS_DB ENV error, %s", redisDB))
    }
    return redis.NewClient(&redis.Options{
      Addr: fmt.Sprintf(
        "%s:%d",
        hmonenv.Get("REDIS_HOST"),
        redisPort),
        Password: hmonenv.Get("REDIS_PASSWORD"),
        DB: redisDB,
  })
}

func SetRaw(key string, data []uint8) error {
    result := redisClient.Set(ctx, key, data, 0)
    return result.Err()
}

func SetRawWithExpired(key string, data []uint8) error {
    ttl := time.Duration(60) * time.Second
    result := redisClient.Set(ctx, key, data, ttl)
    return result.Err()
}

func SetStr(key string, val string) error {
    result := redisClient.Set(ctx, key, val, 0)
    return result.Err()
}

func SetInt(key string, val int) error {
    result := redisClient.Set(ctx, key, val, 0)
    return result.Err()
}

func Get(key string) (string, error) {
    var err error
    result := redisClient.Get(ctx, key)
    err = result.Err()
    if err != nil {
        if err == redis.Nil {
            return "", nil
        }
        return "", err
    }
    var out string
    out, _ = result.Result()
    if err != nil {
        if err == redis.Nil {
            return "", nil
        }
        return "", err
    }
    return out, nil
}

func Del(key string) error {
    result := redisClient.Del(ctx, key)
    return result.Err()
}

func Subscribe(c context.Context, key string) *redis.PubSub {
    return redisClient.Subscribe(c, key)
}

func SubscriberReceiveMessage(subscriber *redis.PubSub) (*redis.Message, error) {
    return subscriber.ReceiveMessage(ctx)
}

func init() {
    redisClient = getRedisClient()
    ctx = context.Background()
}
