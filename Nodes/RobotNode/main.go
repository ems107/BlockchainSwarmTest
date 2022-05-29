package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

const key = "{\"address\":\"fbdaa5dcf6ea0323359a1abb20659868351e86a6\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"c9b411a3091e35264fa5a693dc591488e32137931507c599a83c2b1c7146c9df\",\"cipherparams\":{\"iv\":\"099540aa431b73e512a26e2c73fb066e\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"1b5cf5662dea6a656e5013b2e56b82d881d3fcbbb5b0688ca66f2a17a0df25eb\"},\"mac\":\"f6275b4f837a9d6dbf58a007f45e483f1f382332ca74e2484e8ddda802bbbf26\"},\"id\":\"9327c735-5870-451e-9b51-6771e58f8fa4\",\"version\":3}"

func main() {
	conn, err := ethclient.Dial("http://localhost:3334")
	if err != nil {
		log.Fatalf("Failed to connect to client: %v", err)
	}

}
