package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func main() {
	// network := &chaincfg.MainNetParams
	network := &chaincfg.TestNet3Params

	wif, err := GenerateWIF(network)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("WIF:", wif.String())
	// fmt.Println("Priv key:", wif.PrivKey)

	fmt.Println("Pub key: ", hex.EncodeToString(wif.PrivKey.PubKey().SerializeCompressed()))

	addr, err := GetAddress(network, wif)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Addr:", addr)
	fmt.Println("Addr encode:", addr.EncodeAddress())
}

func GetAddress(network *chaincfg.Params, wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network)
}

func GenerateWIF(network *chaincfg.Params) (*btcutil.WIF, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		log.Fatal(err)
	}

	return btcutil.NewWIF(secret, network, true)
}
