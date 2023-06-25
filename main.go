package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"os"
)

var connection *ethclient.Client
var currentConfig Config

func setUp() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	currentConfig = getTestConfig()
	getConnection()
}

func main() {
	setUp()
	// Replace the wallet addresses and private key with the actual values
	privateKey := "YOUR_PRIVATE_KEY"
	fromAddress := "0x1234567890abcdef1234567890abcdef12345678"
	toAddress := "0xabcdef1234567890abcdef1234567890abcdef12"

	// Get the balance of the from address
	balance, err := getBalance(fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Balance of %s: %s ETH", fromAddress, balance)

	// Transfer all funds from the from address to the to address
	err = transferFunds(privateKey, toAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Funds transferred from %s to %s", fromAddress, toAddress)
}

func getConnection() {
	client, err := ethclient.Dial("https://" + currentConfig.NetType + ".infura.io/v3/8cfea7aaa1384f1a87b6d5aa65119ea3")
	if err != nil {
		log.Fatal(err)
	}
	connection = client
}