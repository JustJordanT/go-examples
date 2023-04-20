package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

const secretKey = "passphrasewhichneedstobe32bytes2"
const secretKey1 = "passphrasewhichneedstobe32bytes3"

func main() {
	text := "My Super Secret Code Stuff"
	encryptedText, err := encryptUsingGcm(text, secretKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Encrypted Text: %v\n", encryptedText)

	decryptedText, err := Decrypt(encryptedText)
	if err != nil {
		fmt.Println("Error decrypting your classified text: ", err)
	}
	fmt.Printf("Decrypted Text: %v\n", string(decryptedText))
}

func encryptUsingGcm(secret string, key string) ([]byte, error) {
	fmt.Println("Encryption Program v0.01")
	text := []byte("My Super Secret Code Stuff")
	secretKey := []byte(key)

	c, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	encryptedText := gcm.Seal(nonce, nonce, text, nil)
	return encryptedText, nil
}

func Decrypt(encryptedText []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(secretKey1))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := encryptedText[:nonceSize], encryptedText[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
