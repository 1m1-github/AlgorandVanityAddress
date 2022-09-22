package main

import (
	"fmt"
	"strings"
	"os"
	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

func main() {
	for {
		account := crypto.GenerateAccount()

		// break condition
		if strings.HasPrefix(account.Address.String(), os.Args[1]) {
			passphrase, _ := mnemonic.FromPrivateKey(account.PrivateKey)
			fmt.Printf("My address: %s\n", account.Address)
			fmt.Printf("My passphrase: %s\n", passphrase)
			break;
		}
	}
}
