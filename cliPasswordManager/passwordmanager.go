package main

import (
	"crypto/rand"
	"io"
	"log"

	"golang.org/x/crypto/nacl/secretbox"
)

// possible usage of ssh public and private keys for local machine
// possible considaration of embedded DB for key,value sttore vs simple .json file (redis, pebble, etc)

const (
	KeySize = 32
	NonceSize = 24
)

func generateKey() (*[KeySize]byte, error) {
	key := new([KeySize]byte)
	_, err := io.ReadFull(rand.Reader, key[:])
	if err != nil {
		log.Fatal("Error: unable to generate key, failed to read random.")
		return nil, err
	}
	return key, nil
}

func generateNonce() (*[NonceSize]byte, error) {
	nonce := new([NonceSize]byte)
	_, err := io.ReadFull(rand.Reader, nonce[:])
	if err != nil {
		log.Fatal("Error: unable to generate nonce, failed to read random.")
		return nil, err
	}
	return nonce, nil
}

//encryption using secretbox, uses symmetric key and nonce
func encryptPassword(password []byte) ([]byte, *[KeySize]byte, *[NonceSize]byte, error) {
	nonce, err := generateNonce()
	if err != nil {
		log.Fatal("Error: unable to generate nonce.")
		return nil, nil, nil, err
	}
	key, err := generateKey()
	if err != nil {
		log.Fatal("Error: unable to generate key.")
		return nil, nil, nil, err
	}

	ciphertext := make([]byte, len(nonce))
	copy(ciphertext, nonce[:])
	ciphertext = secretbox.Seal(ciphertext, password, nonce, key)
	return ciphertext, key, nonce, nil
}