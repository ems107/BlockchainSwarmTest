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
)

const processAddress = "http://localhost:3335"
const key = "{\"address\":\"9c632d3b82943a2b482e70cbab0abaa8952f7e12\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"60418ccf585ca16ec26da3f718d02a0185951c5b729e2611bcf6c9fe1e474d14\",\"cipherparams\":{\"iv\":\"77889e19c3eac002c91850a93092ff7c\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"a0c8a0fc9691996dc2325fdbb58b5fd489a2de95e30279bcd2008f8ecafa1ed7\"},\"mac\":\"f7d09f8b481abcd1aaa9103774665abbf407443c57161650ead355aea05d62ca\"},\"id\":\"ae232407-001d-45a5-b3a1-5ba18dd8b3eb\",\"version\":3}"
const contractAddress = "0x49aaaabfa8711D04AfE8e3682598B97BCfEfeC64"
const accountPubKey = "0x49aaaabfa8711D04AfE8e3682598B97BCfEfeC64"
const passphrase = "miguelangel"

func main() {
	conn, err := ethclient.Dial(processAddress)
	if err != nil {
		log.Fatalf("Failed to connect to client: %v", err)
	}

	contract, err := NewRobotSwarm(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a contract: %v", err)
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), passphrase, big.NewInt(1515))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	for {
		randInt := rand.Intn(10)
		_, err := contract.Increment(&bind.TransactOpts{Signer: auth.Signer, From: common.HexToAddress(accountPubKey)}, big.NewInt(int64(randInt)))
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
	}
}
