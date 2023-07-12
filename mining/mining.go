package mining

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
	"time"
)

type Miner struct {
	Difficulty int
}

func NewMiner(difficulty int) *Miner {
	return &Miner{Difficulty: difficulty}
}

// Execute mining
func (m *Miner) Mine(blockData string) string {
	// Perform proof-of-work mining
	fmt.Print("[")
	startTime := time.Now() // Start time of mining
	nonce := 0
	index := 0.0        // for progress reporting
	compareIndex := 0.0 // for progress reporting
	if m.Difficulty == 1 {
		compareIndex = 1.0
	} else {
		compareIndex = 100000
	}
	for {
		index++
		// Generate a random nonce
		nonceBytes := make([]byte, 8)
		rand.Read(nonceBytes)
		nonceString := hex.EncodeToString(nonceBytes)
		nonce = int(nonceString[0]) + int(nonceString[1]) + int(nonceString[2])

		// Concatenate the block data and nonce
		data := blockData + fmt.Sprintf("%d", nonce)

		// Calculate the hash of the concatenated data
		hash := calculateHash(data)
		if math.Mod(index, compareIndex) == 0 {
			fmt.Print("=")
		}

		// Check if the hash satisfies the mining criteria (e.g., number of leading zeros)
		if strings.HasPrefix(hash, strings.Repeat("0", m.Difficulty)) {
			fmt.Println("]")
			endTime := time.Now()                        // End time of mining
			totalTime := endTime.Sub(startTime).String() // Total mining operation time
			fmt.Println("Mining Operation Time:", totalTime)
			return hash
		}
	}
}

func calculateHash(data string) string {
	// Calculate the hash of a given data
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
