package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"math/rand"
	"log"
)

func debugEncryptDecrypt(masterKey, iv, password string) (string, string) {
	encryptedPassword := encrypt(password, masterKey, iv)
	decryptedPassword := decrypt(encryptedPassword, masterKey, iv)
	return encryptedPassword, decryptedPassword
}

// don't touch below this line

func keyToCipher(key string) (cipher.Block, error) {
	keyBytes := []byte(key)
	return aes.NewCipher(keyBytes)
}

// parts of this function are depricated, look at chapter 2 lesson 11 to use new versions
func generateRandomKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	return fmt.Sprintf("%x", bytes), err
}

func base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
	nBinary := int(bits)
	char := ""
	if (0 <= nBinary && nBinary < 8) {
		char = string(base8Alphabet[nBinary])
	}
	return char
}

func encrypt(plainText, key, iv string) string {
	bytes := []byte(plainText)
	blockCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Println(err)
		return ""
	}
	
	stream := cipher.NewCTR(blockCipher, []byte(iv))
	stream.XORKeyStream(bytes, bytes)
	return fmt.Sprintf("%x", bytes)
}

func decrypt(cipherText, key, iv string) string {
	blockCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Println(err)
		return ""
	}
	stream := cipher.NewCTR(blockCipher, []byte(iv))
	bytes, err := hex.DecodeString(cipherText)
	if err != nil {
		log.Println(err)
		return ""
	}
	stream.XORKeyStream(bytes, bytes)
	return string(bytes)
}

