package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	addr = ""
	pass = ""
	key  = "vote"
)

func randNum(min, max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(max-min+1) + min
}

func setBit(client *redis.Client, k string, pos int, val int) error {
	err := client.BitField(context.Background(), k, "set", "u1", pos, val).Err()
	return err
}

func getBit(client *redis.Client, k string, pos int) (int64, error) {
	bitVal, err := client.BitField(context.Background(), k, "get", "u1", pos).Result()
	return bitVal[0], err
}

func main() {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
	})
	defer c.Close()

	randID := randNum(1, 100_000)
	log.Printf("Citizen ID: %d\n", randID)

	randYesNo := randNum(0, 1)
	log.Printf("Vote: %d\n", randYesNo)

	err := setBit(c, key, randID, randYesNo)
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
