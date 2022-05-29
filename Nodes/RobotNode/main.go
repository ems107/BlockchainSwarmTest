package main

import (
	"bufio"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/melbahja/goph"
)

const key = "{\"address\":\"fbdaa5dcf6ea0323359a1abb20659868351e86a6\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"c9b411a3091e35264fa5a693dc591488e32137931507c599a83c2b1c7146c9df\",\"cipherparams\":{\"iv\":\"099540aa431b73e512a26e2c73fb066e\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"1b5cf5662dea6a656e5013b2e56b82d881d3fcbbb5b0688ca66f2a17a0df25eb\"},\"mac\":\"f6275b4f837a9d6dbf58a007f45e483f1f382332ca74e2484e8ddda802bbbf26\"},\"id\":\"9327c735-5870-451e-9b51-6771e58f8fa4\",\"version\":3}"
const passphrase = "miguelangel"

func main() {
	conn, err := ethclient.Dial("http://localhost:3334")
	if err != nil {
		log.Fatalf("Failed to connect to client: %v", err)
	}

	contract, err := NewRobotSwarm(common.HexToAddress("0x49aaaabfa8711D04AfE8e3682598B97BCfEfeC64"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a contract: %v", err)
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), passphrase, big.NewInt(1515))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	for {
		randInt := rand.Intn(10)
		_, err := contract.Increment(&bind.TransactOpts{Signer: auth.Signer, From: common.HexToAddress("0xfbdaa5dcf6ea0323359a1abb20659868351e86a6")}, big.NewInt(int64(randInt)))
		if err != nil {
			log.Fatalf("Failed to retrieve total: %v", err)
		}
		log.Println("Incrementing by: ", randInt)

		time.Sleep(time.Second * 3)
		total, err := contract.GetTotal(nil)
		log.Printf("Total: %v", total)
		// time.Sleep(time.Second * 5)
		log.Println("Press a button to continue...")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		order, err := contract.GetOrder(nil)
		if err != nil {
			log.Fatalf("Failed to retrieve order: %v", err)
		}
		log.Printf("Order: %v", order)
		// time.Sleep(time.Second * 5)

		log.Println("Executing order in robot...")
		client, err := goph.New("robot", "10.42.0.230", goph.Password("maker"))
		if err != nil {
			log.Fatalf("Failed to connect to robot: %v", err)
		}
		defer client.Close()

		out, err := client.Run("./robot " + order)
		if err != nil {
			log.Fatalf("Failed to execute order: %v", err)
		}
		log.Println(out)

		log.Println("Press a button to continue...")
		input = bufio.NewScanner(os.Stdin)
		input.Scan()
	}
}
