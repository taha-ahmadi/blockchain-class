package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	pr, err := crypto.GenerateKey()
	if err != nil {
		return err
	}
	address := crypto.PubkeyToAddress(pr.PublicKey)
	fmt.Println(address)
	return nil
}
