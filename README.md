# BITCOIN IN A NUTSHELL
<img src="images/banner.png" width="600" height="80" />  

₿ Simplified Proof of Concept demonstrating bitcoin architecture and components interactions ₿

## Abstract
Are you interested in `Bitcoin` but never found the time and/or courage to delve into all the technical details?  
This simplified Proof of Concept implementation of the Bitcoin architecture is for you!  
The goal of this project is to provide beginners with a basic overview of Bitcoin's key functionalities and technical primitives.  

>Please note that this PoC is a simplified version and does not reflect the complete complexity of the actual Bitcoin system.  
>It serves as an educational tool to gain an introductory understanding of the key concepts involved.  


## Table of Contents 
- [BITCOIN IN A NUTSHELL](#bitcoin-in-a-nutshell)
  - [Abstract](#abstract)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Key Concepts](#key-concepts)
    - [Blockchain](#blockchain)
    - [Proof-of-Work (PoW)](#proof-of-work-pow)
    - [Wallets](#wallets)
  - [Project Structure](#project-structure)
  - [Getting Started](#getting-started)
  - [Detailed explanation](#detailed-explanation)
## Introduction

**Bitcoin** is a decentralized digital currency that operates on a peer-to-peer network **without the need for a central authority**.  
It was introduced in 2009 by an anonymous person or group of people using the pseudonym *Satoshi Nakamoto*.  
Bitcoin enables secure, fast, and low-cost transactions, making it an innovative alternative to traditional centralized financial systems.  

This PoC aims to provide a very simplified understanding of Bitcoin's architecture, focusing on the key concepts of `blockchain`, `Proof-of-Work (PoW)`, and `wallets`.  
It demonstrates how these components work together to achieve a decentralized and secure system for digital transactions.  
## Key Concepts
### Blockchain

The blockchain is at the heart of Bitcoin's architecture.  
It is a public, distributed ledger that records all transactions ever made in the Bitcoin network.  
Each transaction is grouped into blocks, and each block contains a reference to the previous block, forming a chain of blocks.  

The blockchain ensures transparency, immutability, and security of transactions.  
Once a block is added to the blockchain, it becomes extremely difficult to alter or tamper with the recorded transactions.  
This makes the blockchain a trustworthy source of information.
### Proof-of-Work (PoW)

Proof-of-Work (PoW) is one of the most importan concepts around Bitcoin. 
It is a consensus mechanism employed by Bitcoin to validate and secure transactions in a decentralized manner.  
It involves a process where miners compete to solve a complex mathematical puzzle by performing computational work.  

Miners continuously modify the data within a block and calculate its [hash](https://en.wikipedia.org/wiki/Hash_function) until they discover a hash that fulfills specific criteria, particularly a certain number of leading zeros, defined by the difficulty level.  
This adjustable difficulty level guarantees that the average time required to find a valid solution remains constant, irrespective of the computational power available within the network.  

Upon finding a valid solution, the miner broadcasts it to the network.  
Other nodes in the network can quickly verify the correctness of the solution by performing a simple hash calculation using the provided data.  
This validation process ensures the integrity of the transactions and prevents malicious actors from easily manipulating the blockchain.  

In summary, Proof-of-Work incentivizes miners to compete and invest computational resources in order to secure the Bitcoin network.  
It establishes a trustless consensus protocol, where the discovery of a valid solution serves as proof of the miner's effort and provides assurance of the validity of transactions within the blockchain.  
### Wallets

Wallets are software or hardware tools that enable users to manage their Bitcoin holdings and perform transactions.  
Each wallet has a unique address associated with it, which acts as a digital identifier.  
Wallets store private keys, which are used to sign transactions and provide proof of ownership.  

In this PoC, we simulate the creation of two wallets.  
The first wallet belongs to a miner node that mines new blocks and receives mining rewards, while the second wallet is used as the recipient of transactions.  
Wallets can send funds by creating transactions that are added to the blockchain.  
Balances are maintained by subtracting outgoing transactions and adding incoming transactions.  
## Project Structure

The project follows a simple structure to demonstrate the core functionalities: 
- `blockchain`: Contains the implementation of the blockchain and its blocks. 
- `mining`: Implements the mining process using Proof-of-Work. 
- `wallet`: Includes the wallet functionality for managing balances and transactions. 
- `main.go`: The main program that initializes the components and demonstrates the functionalities.
## Getting Started

To run this PoC, follow these steps: 
1. Clone the repository: `git clone https://github.com/r3drun3/bitcoin-simplified.git`. 
2. Navigate to the project directory: `cd bitcoin-simplified`.
3. Ensure you have Go installed (v1.16+). 
4. Build and run the project: `go run main.go`.
5. Examine the printed outputs to understand the blockchain, mining, and wallet functionalities.  
6. If you want, modify the mining difficulty in `main.go` and notice how the runtime increases exponentially!!  


## Detailed explanation
The script output demonstrates the execution and functionality of the "Bitcoin in a Nutshell" PoC:  
```console
Blockchain:
Data: This is the Genesis Block!
Hash: 
Previous: 

Data: Transaction 1
Hash: dff3b30655dc240deca00ed22fae68fdf8cf465bbe99bb2b2e24259cc1daac3a
Previous: 

Data: Transaction 2
Hash: 6552affdd08bdab9e4d071a00269fd40e4c9212f3e7cd8307d413c5aaee3ce98
Previous: dff3b30655dc240deca00ed22fae68fdf8cf465bbe99bb2b2e24259cc1daac3a


Wallet 1 Initial Balance:
Current balance: 0.00 BTC

Wallet 2 Initial Balance:
Current balance: 0.00 BTC

Can Wallet 1 send 5.0 BTC to wallet 2?
NO !!! Insufficient balance to send funds 😔 😔 😔 

Mining Block:
[===========================]
Mining Operation Time: 247.63µs
Mined Block Hash: 0341ca503accb4f312e855088036ce152a3e418b90c306ed1080651e7b7b6a9a
Successfully mined 10.00 BTC, sending to miner wallet!

Wallet 1 Updated Balance:
Current balance: 10.00 BTC

Wallet 2 Updated Balance:
Current balance: 0.00 BTC

Can Wallet 1 send 5.0 BTC to wallet 2?
Yes !!! Successfully sent 5.00 BTC to recipient 💰 🤑 💰

Wallet 1 Final Balance:
Current balance: 5.00 BTC

Wallet 2 Final Balance:
Current balance: 5.00 BTC
```   
1. The program starts by displaying the existing blockchain. It shows three blocks:
- Genesis Block: The first block in the blockchain, indicated by an empty hash and no previous block reference.
- Transaction 1 Block: Contains the transaction data "*Transaction 1*" and a corresponding hash.
- Transaction 2 Block: Contains the transaction data "*Transaction 2*" and a corresponding hash, with a reference to the previous block's hash. 
1. The initial balances of Wallet 1 and Wallet 2 are displayed as 0.00 BTC. 
1. The program checks if Wallet 1 can send 5.0 BTC to Wallet 2.  
   Since Wallet 1 has a balance of 0.00 BTC, it displays an insufficient balance error message. 
1. The *mining process* begins, indicated by the "*Mining Block*" message on the screen.  
   A progress bar is displayed, showing the mining progress.  
   After completing the mining operation, it prints the mining operation time and the mined block's hash. 
1. Wallet 1, the wallet of the miner that completed the process, receives a mining reward of 10.00 BTC.  
   The updated balance of Wallet 1 is displayed as 10.00 BTC. 
1. The program checks if Wallet 1 can now send 5.0 BTC to Wallet 2.  
   Since Wallet 1 now has a balance of 10.00 BTC, it successfully sends 5.00 BTC to Wallet 2. 
1. The updated balances of Wallet 1 and Wallet 2 are displayed as 5.00 BTC each.  

The script output demonstrates the mining process, the handling of balances and transactions, and the interaction between the wallets in the PoC. It showcases the functionality of mining rewards, verifying balances, and successful transactions.


