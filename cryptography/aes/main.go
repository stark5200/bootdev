package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	plaintext := []byte("password123")
	fmt.Printf("Plaintext: %s\n", plaintext)
	// Generate Master Key
	masterkey, err := generateMasterKey()
	if err != nil {
		fmt.Println("Error: unable to generate Master Key :", err)
	}
	//fmt.Printf("Generated Master Key: %x\n", masterkey)
	lestring := fmt.Sprintf("%x", masterkey)
	masterKeyString := hex.EncodeToString(masterkey)
	fmt.Println(lestring)
	// Generate IV
	iv, err := generateIV()
	if err != nil {
		fmt.Println("Error: unable to generate IV :", err)
	}
	fmt.Printf("Generated IV: %x\n", iv)
	fmt.Println()

	// Encrypt	
	ciphertextWithIV := AESencrypt(plaintext, masterkey, iv)
	fmt.Printf("Ciphertext (IV + Encrypted Data in hex): %s\n", ciphertextWithIV)

	// Decrypt 
	decrypted := AESdecrypt(ciphertextWithIV, masterkey)
	fmt.Printf("Decrypted text: %s\n", decrypted)
}

//////////
// don't touch below this line

func AESencrypt(plaintext, key, iv []byte) string {
	plaintextCopy := make([]byte, len(plaintext))
	copy(plaintextCopy, plaintext)

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return ""
	}
	// use cipher in counter mode
	stream := cipher.NewCTR(blockCipher, iv)
	stream.XORKeyStream(plaintextCopy, plaintextCopy)

	// Append the IV to the ciphertext
	cipherTextWithIV := append(iv, plaintextCopy...)

	return fmt.Sprintf("%x", cipherTextWithIV)
}

func AESdecrypt(ciphertextWithIVHex string, key []byte) string {
	// Decode the hex-encoded ciphertext+IV back into raw bytes
	ciphertextWithIV, err := hex.DecodeString(ciphertextWithIVHex)
	if err != nil {
		log.Println("Error decoding hex:", err)
		return ""
	}

	// Ensure the input has enough bytes for at least an IV
	if len(ciphertextWithIV) < aes.BlockSize {
		log.Println("Error: Ciphertext too short")
		return ""
	}

	// Extract the IV
	iv := ciphertextWithIV[:aes.BlockSize]
	// Extract the actual ciphertext (remaining bytes)
	ciphertext := ciphertextWithIV[aes.BlockSize:]

	// Initialize the aes block cipher
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error creating cipher:", err)
		return ""
	}
	// use cipher in counter (CTR) mode for decryption
	stream := cipher.NewCTR(blockCipher, iv)

	// Make a copy of the ciphertext to avoid modifying the original slice
	ciphertextCopy := make([]byte, len(ciphertext))
	copy(ciphertextCopy, ciphertext)
	// Decrypt the data
	stream.XORKeyStream(ciphertextCopy, ciphertextCopy)

	// Return the decrypted text as a string
	return string(ciphertextCopy)
}

func generateMasterKey() ([]byte, error) {
	key := make([]byte, 16) // 128 bits
	_, err := rand.Read(key)
	if err != nil {
		return []byte(""), err
	}
	// Convert the key to a readable hex string
	return key, nil
}

func generateIV() ([]byte, error) {
	iv := make([]byte, 16) // AES block size (128 bits)
	_, err := rand.Read(iv)
	if err != nil {
		return []byte(""), err
	}
	// Convert the IV to a readable hex string
	return iv, nil
}
