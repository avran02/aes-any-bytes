# aes-any-bytes

## Introduction

aes any bytes is a simple and lightweight go library implementing aes algorythm

## Installation
You can install it with 
`go get github.com/avran02/aes-any-bytes`

## Examples

```
package main

import (
	"fmt"
	aes "github.com/avran02/aes-any-bytes"
)

func main() {
	// Example usage
	plaintext := []byte("This is a secret message")
	key := []byte("thisisa16bytekey") // Should be 16, 24, or 32 bytes for AES 128, 192 or 256

	// Encrypt
	ciphertext, err := aes.Encript(plaintext, key)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}

	// Decrypt
	decryptedText, err := aes.Decript(ciphertext, key)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}

	fmt.Println("Original:", string(plaintext))
	fmt.Println("Decrypted:", string(decryptedText))
}

```
More examples will be added soon