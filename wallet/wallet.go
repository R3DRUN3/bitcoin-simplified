package wallet

import "fmt"

type Wallet struct {
	Balance float64
}

func NewWallet() *Wallet {
	return &Wallet{Balance: 0}
}

func (w *Wallet) SendFunds(amount float64, recipient *Wallet) {
	// Check the sender wallet's balance
	if w.Balance >= amount {
		// Create a new transaction and add it to the blockchain
		// Update the balances of the sender and recipient
		w.Balance -= amount
		recipient.Balance += amount
		fmt.Printf("Yes !!! Successfully sent %.2f BTC to recipient 💰 🤑 💰\n", amount)
	} else {
		fmt.Println("NO !!! Insufficient balance to send funds 😔 😔 😔 ")
	}
}

func (w *Wallet) PrintBalance() {
	// Print the wallet's balance
	fmt.Printf("Current balance: %.2f BTC\n", w.Balance)
}

func (w *Wallet) MineReward() {
	// Reward the miner by adding funds to the wallet
	reward := 10.0 // Fixed reward for mining a block
	w.Balance += reward
	fmt.Printf("Successfully mined %.2f BTC, sending to miner wallet!\n", reward)
}
