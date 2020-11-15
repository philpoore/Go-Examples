package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	key := []byte("12345678901234567890123456789012")
	block, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println("Error creating cipher")
		os.Exit(1)
	}

	// msg := []byte("Hello world     ")
	msg, _ := base64.StdEncoding.DecodeString("HpMehHx/G/vo8fMv965x6Q==")

	var enMsg = make([]byte, 16)

	fmt.Println("Block Size:", block.BlockSize())
	fmt.Println("Msg Size:", len(msg))
	// block.Encrypt(enMsg, msg)
	block.Decrypt(enMsg, msg)
	fmt.Printf("Message: %v\n", string(msg))
	fmt.Printf("Encrypted Message: %v\n", string(enMsg))
	// encoded := base64.StdEncoding.EncodeToString(enMsg)
	// fmt.Printf("Encrypted Message: %v\n", encoded)
}
