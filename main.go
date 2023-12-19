package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	redisAddress  = ""
	redisPassword = ""
	cacheKey      = "vote"
)

// getRandomNumber generates a random number between min and max (inclusive).
func getRandomNumber(min, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(max-min+1) + min
}

// setBitField sets the bit at the given position in the Redis bitfield.
func setBitField(client *redis.Client, cacheKey string, position int, value int) error {
	err := client.BitField(context.Background(), cacheKey, "set", "u32", position, value).Err()
	return err
}

// getBitField gets the bit at the given position from the Redis bitfield.
func getBitField(client *redis.Client, cacheKey string, position int) (int64, error) {
	bitValue, err := client.BitField(context.Background(), cacheKey, "get", "u32", position).Result()
	return bitValue[0], err
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
	})
	defer client.Close()

	// Generate random citizen ID and vote number
	randomCitizenID := getRandomNumber(1, 100)
	log.Printf("Citizen ID: %d\n", randomCitizenID)

	randomVoteNumber := getRandomNumber(1, 20)
	log.Printf("Vote number: %d\n", randomVoteNumber)

	// Set the bit in the Redis bitfield
	err := setBitField(client, cacheKey, randomCitizenID, randomVoteNumber)
	if err != nil {
		log.Fatalf("Error setting bit: %v", err)
	}

	log.Println("Bit set successfully")

	// Retrieve the bit from the Redis bitfield
	bitValue, err := getBitField(client, cacheKey, randomCitizenID)
	if err != nil {
		log.Fatalf("Error getting bit: %v", err)
	}

	log.Printf("Bit value: %d\n", bitValue)
}
