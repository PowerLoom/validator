package redis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
	"validator/config"
	"validator/pkgs"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func NewRedisClient() *redis.Client {
	db, err := strconv.Atoi(config.SettingsObj.RedisDB)
	if err != nil {
		log.Fatalf("Incorrect redis db: %s", err.Error())
	}
	return redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.SettingsObj.RedisHost, config.SettingsObj.RedisPort), // Redis server address
		Password:     "",                                                                               // no password set
		DB:           db,
		PoolSize:     1000,
		ReadTimeout:  200 * time.Millisecond,
		WriteTimeout: 200 * time.Millisecond,
		DialTimeout:  5 * time.Second,
		IdleTimeout:  5 * time.Minute,
	})
}

func AddToSet(ctx context.Context, set string, keys ...string) error {
	if err := RedisClient.SAdd(ctx, set, keys).Err(); err != nil {
		return fmt.Errorf("unable to add to set: %s", err.Error())
	}
	return nil
}

func Get(ctx context.Context, key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		} else {
			return "", err
		}
	}
	return val, nil
}

func Set(ctx context.Context, key, value string) error {
	return RedisClient.Set(ctx, key, value, 0).Err()
}

// Use this when you want to set an expiration
func SetWithExpiration(ctx context.Context, key, value string, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

func StoreValidatorDetails(ctx context.Context, dataMarketAddress, epochID, rootHash, details string) error {
	// Store the batch roothash in the master set
	if err := AddToSet(ctx, ValidatorSet(dataMarketAddress, epochID), rootHash); err != nil {
		return fmt.Errorf("failed to add roothash %s to validator set for epoch %s, data market %s: %v", rootHash, epochID, dataMarketAddress, err)
	}

	// Store the batch details associated with the given roothash in Redis
	if err := SetWithExpiration(ctx, SnapshotSubmissionValidatorKey(dataMarketAddress, epochID, rootHash), details, pkgs.Day*3); err != nil {
		return fmt.Errorf("failed to store batch details for roothash %s, epoch %s, data market %s in Redis: %w", rootHash, epochID, dataMarketAddress, err)
	}

	return nil
}
