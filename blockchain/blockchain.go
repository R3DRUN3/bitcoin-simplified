package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Data     string
	Hash     string
	Previous string
}

type Blockchain struct {
	Chain []Block
}

func NewBlockchain() *Blockchain {
	// Initialize the blockchain with a genesis block
	genesisBlock := Block{Data: "This is the Genesis Block!", Hash: "", Previous: ""}
	chain := []Block{genesisBlock}
	return &Blockchain{Chain: chain}
}

func (bc *Blockchain) AddBlock(data string) {
	// Add a new block to the blockchain
	previousBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := Block{Data: data, Previous: previousBlock.Hash}
	// Calculate the hash of the new block
	newBlock.Hash = calculateHash(newBlock)
	// Append the new block to the blockchain
	bc.Chain = append(bc.Chain, newBlock)
}

func (bc *Blockchain) GetBlockchain() []Block {
	// Return the entire blockchain
	return bc.Chain
}

func calculateHash(block Block) string {
	// Calculate the hash of a block
	data := block.Data + block.Previous
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
