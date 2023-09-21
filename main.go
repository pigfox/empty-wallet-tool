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
	privateKey := "47acca3128a6e46469fa867a2d7d01a3b7fdb1349c6dba74967a669df2ad1de9"
	fromAddress := "0x0c003e08f98748cb4ff3fe3abe88b11ed4c64add"
	toAddress := "0x3f00b9c2c1d3548e0d2afdc96d6cd620c1b4a7bc"

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
