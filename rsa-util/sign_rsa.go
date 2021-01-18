package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

//generateKeyPair3: 生成key对
func generateKeyPair3(bites int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bites)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return privateKey, &privateKey.PublicKey
}

func main() {
	privateKey, publicKey := generateKeyPair3(2048)
	message := []byte("super secret message")
	// We actually sign the hashed message
	msgHash := sha256.New()
	msgHash.Write(message)
	msgHashSum := msgHash.Sum(nil)
	// We have to provide a random reader, so every time
	// we sign, we have a different signature
	signature, _ := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	fmt.Printf("Signature: %v\n", signature)

	err := rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("Verification failed: ", err)
	} else {
		fmt.Println("Message verified.")
	}
}
