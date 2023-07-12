package main

import (
	"fmt"

	"github.com/r3drun3/bitcoin-simplified/blockchain"
	"github.com/r3drun3/bitcoin-simplified/mining"
	"github.com/r3drun3/bitcoin-simplified/wallet"
)

func main() {
	// Create a new blockchain
	bc := blockchain.NewBlockchain()

	// Add blocks to the blockchain
	bc.AddBlock("Transaction 1")
	bc.AddBlock("Transaction 2")

	// Get the entire blockchain
	chain := bc.GetBlockchain()
	fmt.Println("Blockchain:")
	for _, block := range chain {
		fmt.Printf("Data: %s\nHash: %s\nPrevious: %s\n\n", block.Data, block.Hash, block.Previous)
	}

	// Proof of work difficulty: number of leading zeros in hash
	var difficulty = 1

	// Create a new miner with the specified difficulty
	m := mining.NewMiner(difficulty)

	// Generate two wallets
	wallet1 := wallet.NewWallet()
	wallet2 := wallet.NewWallet()

	// Print the initial balances
	fmt.Println("\nWallet 1 Initial Balance:")
	wallet1.PrintBalance()
	fmt.Println("\nWallet 2 Initial Balance:")
	wallet2.PrintBalance()

	fmt.Println("\nCan Wallet 1 send 5.0 BTC to wallet 2?")
	// Send funds from wallet1 to wallet2
	wallet1.SendFunds(5.0, wallet2)

	// Mine a new block
	fmt.Println("\nMining Block:")

	minedHash := m.Mine("Mining Block Data")
	fmt.Println("Mined Block Hash:", minedHash)

	// Mine a reward for wallet1
	wallet1.MineReward()

	// Print the updated balances
	fmt.Println("\nWallet 1 Updated Balance:")
	wallet1.PrintBalance()
	fmt.Println("\nWallet 2 Updated Balance:")
	wallet2.PrintBalance()

	fmt.Println("\nCan Wallet 1 send 5.0 BTC to wallet 2?")
	// Send funds from wallet1 to wallet2
	wallet1.SendFunds(5.0, wallet2)

	// Print the final balances
	fmt.Println("\nWallet 1 Final Balance:")
	wallet1.PrintBalance()
	fmt.Println("\nWallet 2 Final Balance:")
	wallet2.PrintBalance()
}
