package random

import (
	"crypto/rand"
	"encoding/hex"
	math "math/rand"
	"time"
)

const (
	_attemptsToReadRandomData = 3
)

func init() {
	math.Seed(time.Now().Unix())
}

func generateRandomData(bitSize int) []byte {
	buffer := make([]byte, bitSize)

	for i := 0; i < _attemptsToReadRandomData; i++ {
		_, err := rand.Read(buffer)
		if err == nil {
			break
		}
	}

	return buffer
}

func generateRandomHexString(bitSize int) string {
	return hex.EncodeToString(generateRandomData(bitSize))
}

// GenerateID generates a ID with the given prefix
func GenerateID(prefix string) string {
	return prefix + generateRandomHexString(16)
}
