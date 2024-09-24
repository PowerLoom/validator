package helpers

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"runtime"
)

func simulateUsage(ctx context.Context, client *redis.Client, setKey string, numKeys int) {
	for i := 0; i < numKeys; i++ {
		key := "test-key-" + strconv.Itoa(i)
		value := "test-value-" + strconv.Itoa(i)
		_ = SetWithSetTracking(ctx, client, key, value, 10*time.Second, setKey)
	}

	for i := 0; i < numKeys/2; i++ {
		key := "test-key-" + strconv.Itoa(i)
		_ = DelWithSetTracking(ctx, client, []string{key}, setKey)
	}
}

func BenchmarkRedisSetTracking(b *testing.B) {
	s, err := miniredis.Run()
	if err != nil {
		b.Fatalf("Failed to start miniredis: %v", err)
	}
	defer s.Close()

	client := newTestRedisClient(s.Addr())
	ctx := context.Background()

	setKey := "benchmark-set"
	numKeys := 1000

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		simulateUsage(ctx, client, setKey, numKeys)
	}
}

func BenchmarkMemoryUsage(b *testing.B) {
	s, err := miniredis.Run()
	if err != nil {
		b.Fatalf("Failed to start miniredis: %v", err)
	}
	defer s.Close()

	client := newTestRedisClient(s.Addr())
	ctx := context.Background()

	setKey := "benchmark-memory-set"
	numKeys := 1000

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		simulateUsage(ctx, client, setKey, numKeys)

		b.StopTimer()
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)
		b.Logf("Memory Usage After Iteration %d: Alloc = %v MiB, TotalAlloc = %v MiB, HeapAlloc = %v MiB, HeapObjects = %v",
			i+1,
			memStats.Alloc/1024/1024,
			memStats.TotalAlloc/1024/1024,
			memStats.HeapAlloc/1024/1024,
			memStats.HeapObjects)
		b.StartTimer()
	}
}
