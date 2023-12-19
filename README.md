# Redis Bitfield Example In Golang

This Go program demonstrates the usage of Redis bitfields to store and retrieve voting information for citizens. Each citizen is assigned a unique ID, and their vote is stored as a bit in a Redis bitfield.

## Prerequisites

Make sure you have Go installed on your machine. You also need a Redis server running and accessible.

```bash
go get -u github.com/go-redis/redis/v8
```

## Usage

1. Clone this repository:

```bash
git clone https://github.com/naruepanart/redis-bitfields.git
```

2. Navigate to the project directory:

```bash
cd naruepanart/redis-bitfields
```

3. Update the Redis server details in the `main.go` file:

```go
const (
	redisAddress  = "your-redis-address"
	redisPassword = "your-redis-password"
	cacheKey      = "vote"
)
```

4. Run the program:

```bash
go run main.go
```

The program will generate a random Citizen ID and Vote Number, set the corresponding bit in the Redis bitfield, and then retrieve and print the bit value.

## Configuration

- `redisAddress`: Address of the Redis server.
- `redisPassword`: Password for connecting to the Redis server.
- `cacheKey`: Redis key used for storing bitfield data.

## Functions

- `getRandomNumber(min, max int) int`: Generates a random number between `min` and `max` (inclusive).
- `setBitField(client *redis.Client, cacheKey string, position int, value int) error`: Sets the bit at the given position in the Redis bitfield.
- `getBitField(client *redis.Client, cacheKey string, position int) (int64, error)`: Retrieves the bit at the given position from the Redis bitfield.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.