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
		DB:       0,                                                                                // use default DB
	})
}

func Set(ctx context.Context, client *redis.Client, key string, value string, expiration time.Duration) error {
	if _, err := Get(ctx, client, key); err != nil && !errors.Is(err, redis.Nil) {
		return err
	}
	if err := client.Set(ctx, key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func Get(ctx context.Context, client *redis.Client, key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", err
	}
	return val, nil
}

func FetchKeysForPattern(ctx context.Context, client *redis.Client, pattern string) ([]string, error) {
	var allKeys []string
	var cursor uint64

	for {
		keys, newCursor, err := client.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error scanning keys for pattern %s: %v", pattern, err))
		}
		allKeys = append(allKeys, keys...)

		cursor = newCursor
		if cursor == 0 {
			break
		}
	}
	return allKeys, nil
}

func ResetValidatorDBSubmissions(ctx context.Context, client *redis.Client, epochID *big.Int) {
	pattern := fmt.Sprintf("%s.%d.*", ValidatorKey, epochID)

	// Use Scan to find keys matching the pattern.
	var cursor uint64
	var n int
	var keysToDelete []string
	for {
		var keys []string
		var err error
		keys, cursor, err = client.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			log.Errorf("Error scanning keys: %v", err)
		}
		keysToDelete = append(keysToDelete, keys...)
		n += len(keys)

		if cursor == 0 { // No more keys
			break
		}
	}

	// Delete the keys.
	if len(keysToDelete) > 0 {
		_, err := client.Del(ctx, keysToDelete...).Result()
		if err != nil {
			log.Errorf("Error deleting keys: %v", err)
		}
		log.Debugf("Deleted %d keys.\n", n)
	} else {
		log.Debugln("No keys found to delete.")
	}
}
