package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	redisAddr = "" // Replace with your Redis server address
	redisPass = "" // Replace with your Redis password
	redisKey  = "vote"
)

func randNum(min, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(max-min+1) + min
}

func setBit(client *redis.Client, key string, pos, val int) error {
	err := client.BitField(context.Background(), key, "set", "u1", pos, val).Err()
	return err
}

func getBit(client *redis.Client, key string, pos int) (int64, error) {
	bitVal, err := client.BitField(context.Background(), key, "get", "u1", pos).Result()
	return bitVal[0], err
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
	})
	defer client.Close()

	rID := randNum(1, 100_000)
	log.Printf("Citizen ID: %d\n", rID)

	rVote := randNum(0, 1)
	log.Printf("Vote: %d\n", rVote)

	err := setBit(client, redisKey, rID, rVote)
	if err != nil {
		log.Fatalf("Error setting bit: %v", err)
	}

	log.Println("Bit set successfully")

	bitVal, err := getBit(client, redisKey, rID)
	if err != nil {
		log.Fatalf("Error getting bit: %v", err)
	}

	log.Printf("Bit value: %d\n", bitVal)
}
