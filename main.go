package main

import (
	"fmt"
	"github.com/AronMc/cryptit/encrypt"
	"github.com/AronMc/cryptit/decrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
