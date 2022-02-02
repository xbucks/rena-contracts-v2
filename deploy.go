package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	greeter "github.com/hirokimoto/Greeter/contracts"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/8c9a3de5eb604d56b8ea15f081d5bdd9")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA("aed7147564caaea6dd4ef501c8b880c6b2c39cb5034b9c80c9eb11314449f658")
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = big.NewInt(10000)

	address, tx, instance, err := greeter.DeployGreeter(auth, client, "Hello world!")
	if err != nil {
		panic(err)
	}

	fmt.Println(address.Hex())

	_, _ = instance, tx
}
