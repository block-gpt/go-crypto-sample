package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	// client, err := ethclient.Dial("https://kovan.infura.io/v3/2cc44ba2af1d4142b166509013cf3b05")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// tokenAddress := common.HexToAddress("0x18Edd1E7B1e906B0223DceE6983bb664a9E0feFa")
	// instance, err := token.NewToken(tokenAddress, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	privateKey, err := crypto.HexToECDSA(os.Getenv("TEST3_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("priv key:", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("pub key:", hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address:", address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("address:", hexutil.Encode(hash.Sum(nil)[12:]))
}
