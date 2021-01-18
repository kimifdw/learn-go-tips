package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

//generateKeyPair: 生成key对
func GenerateKeyPair2(bites int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bites)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return privateKey, &privateKey.PublicKey
}

func main() {
	privateKey, publicKey := GenerateKeyPair2(2048)
	message := []byte("Hello, Emily")
	cipher, _ := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey, message, nil)
	fmt.Printf("Encrypted: %v\n", cipher)

	decMessage, _ := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		cipher,
		nil,
	)
	fmt.Printf("Original: %s\n", string(decMessage))
}
