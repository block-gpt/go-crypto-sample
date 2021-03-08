package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/fudorec/geth-sample/post-contract/contract"
	"github.com/fudorec/geth-sample/query-erc20/token"
)

func main() {
	client, err := ethclient.Dial("https://kovan.infura.io/v3/2cc44ba2af1d4142b166509013cf3b05")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("TEST3_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("address:", address.Hex())

	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	tokenAddress := common.HexToAddress("0x75b3267f7C769E36412B129A26F59Fe144c2C8A1")
	tokenContract, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	bal, err := tokenContract.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bal in wei:", bal)

	contractAddress := common.HexToAddress("0xcEa856877E28BfdC1Bd72F929DD8B1e474dc1f64")
	instance, err := contract.NewContract(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.Claim(auth)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	// bal, err = tokenContract.BalanceOf(&bind.CallOpts{}, address)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Bal in wei:", bal)
}
