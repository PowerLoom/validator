package helpers

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"math/big"
	"time"
	"validator/config"
)

var RedisClient *redis.Client

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.SettingsObj.RedisHost, config.SettingsObj.RedisPort), // Redis server address
		Password: "",                                                                               // no password set
		DB:       config.SettingsObj.RedisDB,                                                       // use default DB
	})
}

func Get(ctx context.Context, client *redis.Client, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", err
	}
	return val, nil
}

func ResetValidatorDBSubmissions(ctx context.Context, client *redis.Client, epochID *big.Int) {
	setKey := fmt.Sprintf("%s.%d", ValidatorKey, epochID)

	keysToDelete, err := FetchKeysFromSet(ctx, client, setKey)
	if err != nil {
		log.Errorf("Error fetching keys from set: %v", err)
		return
	}

	if len(keysToDelete) > 0 {
		if err := DelWithSetTracking(ctx, client, keysToDelete, setKey); err != nil {
			log.Errorf("Error deleting keys: %v", err)
		}
		log.Debugf("Deleted %d keys.\n", len(keysToDelete))
	} else {
		log.Debugln("No keys found to delete.")
	}
}

func AddToSet(ctx context.Context, client *redis.Client, setKey, memberKey string) error {
	return client.SAdd(ctx, setKey, memberKey).Err()
}

func RemoveFromSet(ctx context.Context, client *redis.Client, setKey, memberKey string) error {
	return client.SRem(ctx, setKey, memberKey).Err()
}

func SetWithSetTracking(ctx context.Context, client *redis.Client, key, value string, expiration time.Duration, setKey string) error {
	if err := client.Set(ctx, key, value, expiration).Err(); err != nil {
		return err
	}
	if err := AddToSet(ctx, client, setKey, key); err != nil {
		return err
	}
	return nil
}

func DelWithSetTracking(ctx context.Context, client *redis.Client, keys []string, setKey string) error {
	if _, err := client.Del(ctx, keys...).Result(); err != nil {
		return err
	}
	for _, key := range keys {
		if err := RemoveFromSet(ctx, client, setKey, key); err != nil {
			return err
		}
	}
	return nil
}

func FetchKeysFromSet(ctx context.Context, client *redis.Client, setKey string) ([]string, error) {
	return client.SMembers(ctx, setKey).Result()
}
