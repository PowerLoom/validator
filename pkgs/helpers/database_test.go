package helpers

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func newTestRedisClient(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func TestSetWithSetTracking(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatalf("Failed to start miniredis: %v", err)
	}
	defer s.Close()

	client := newTestRedisClient(s.Addr())

	ctx := context.Background()
	setKey := "test-set"
	key := "test-key"
	value := "test-value"

	err = SetWithSetTracking(ctx, client, key, value, 10*time.Second, setKey)
	assert.NoError(t, err, "Error setting key with tracking")

	val, err := client.Get(ctx, key).Result()
	assert.NoError(t, err, "Error getting key from Redis")
	assert.Equal(t, value, val, "Stored value does not match")

	members, err := client.SMembers(ctx, setKey).Result()
	assert.NoError(t, err, "Error getting members of the set")
	assert.Contains(t, members, key, "Key not found in the set")
}

func TestDelWithSetTracking(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatalf("Failed to start miniredis: %v", err)
	}
	defer s.Close()

	client := newTestRedisClient(s.Addr())

	ctx := context.Background()
	setKey := "test-set"
	key := "test-key"
	value := "test-value"

	err = SetWithSetTracking(ctx, client, key, value, 10*time.Second, setKey)
	assert.NoError(t, err, "Error setting key with tracking")

	err = DelWithSetTracking(ctx, client, []string{key}, setKey)
	assert.NoError(t, err, "Error deleting key with tracking")

	_, err = client.Get(ctx, key).Result()
	assert.ErrorIs(t, err, redis.Nil, "Key should have been deleted")

	members, err := client.SMembers(ctx, setKey).Result()
	assert.NoError(t, err, "Error getting members of the set")
	assert.NotContains(t, members, key, "Key should have been removed from the set")
}

func TestFetchKeysFromSet(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		t.Fatalf("Failed to start miniredis: %v", err)
	}
	defer s.Close()

	client := newTestRedisClient(s.Addr())

	ctx := context.Background()
	setKey := "test-set"
	keys := []string{"key1", "key2", "key3"}

	for _, key := range keys {
		err = AddToSet(ctx, client, setKey, key)
		assert.NoError(t, err, "Error adding key to set")
	}

	fetchedKeys, err := FetchKeysFromSet(ctx, client, setKey)
	assert.NoError(t, err, "Error fetching keys from set")
	assert.ElementsMatch(t, keys, fetchedKeys, "Fetched keys do not match")
}
