package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

func transferFunds(fromPrivateKey, toAddress string) error {
	// Convert the private key string to a crypto.PrivateKey
	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	// Get the public key and address from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("failed to get public key from private key")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get the current nonce for the from address
	nonce, err := connection.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get the gas price
	gasPrice, err := connection.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %v", err)
	}

	// Get the current balance of the from address
	balance, err := connection.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		return fmt.Errorf("failed to get balance: %v", err)
	}

	// Set the transfer amount to the available balance
	amount := new(big.Int).Set(balance)

	// Define the gas limit
	gasLimit := uint64(21000) // Example gas limit value, adjust according to your needs

	// Convert the recipient's address to a common.Address
	toAddressObj := common.HexToAddress(toAddress)

	// Create the unsigned transaction with the updated recipient's address
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddressObj,
		Value:    amount,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		Data:     nil,
	})

	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send the transaction
	err = connection.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %v", err)
	}

	// Print the transaction hash as a confirmation
	log.Printf("Transaction sent: %s", signedTx.Hash().Hex())

	return nil
}
