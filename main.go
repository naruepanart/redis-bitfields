package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	addr  = ""
	pass = ""
	key = "vote"
)

func randNum(min, max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(max-min+1) + min
}

func setBit(client *redis.Client, k string, pos string, val int) error {
	err := client.BitField(context.Background(), k, "set", "u32", pos, val).Err()
	return err
}

func getBit(client *redis.Client, k string, pos string) (int64, error) {
	bitVal, err := client.BitField(context.Background(), k, "get", "u32", pos).Result()
	return bitVal[0], err
}

func main() {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
	})
	defer c.Close()

	randID := fmt.Sprintf("#%d", randNum(1, 100_000))
	log.Printf("Citizen ID: %s\n", randID)

	randVoteNum := randNum(1, 20)
	log.Printf("Vote number: %d\n", randVoteNum)

	err := setBit(c, key, randID, randVoteNum)
	if err != nil {
		log.Fatalf("Error setting bit: %v", err)
	}

	log.Println("Bit set successfully")

	bitVal, err := getBit(c, key, randID)
	if err != nil {
		log.Fatalf("Error getting bit: %v", err)
	}

	log.Printf("Bit value: %d\n", bitVal)
}