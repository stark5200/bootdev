package main

import (
	"fmt"
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/gob"
	"encoding/base64"
	"io"
	"log"
	"errors"
	"golang.org/x/crypto/nacl/secretbox"
)

// possible considaration of embedded DB for key,value sttore vs simple .json file (redis, pebble, etc)

const (
	KeySize = 32
	NonceSize = 24
)

type EncryptPayload struct {
	Key *[KeySize]byte
	Nonce *[NonceSize]byte
	Ciphertext []byte
}

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

func genRsaKeys() (pubKey *rsa.PublicKey, privKey *rsa.PrivateKey, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return &privateKey.PublicKey, privateKey, nil
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

//encryption using public key
func encryptWithPublicKey(publicKey *rsa.PublicKey , ciphertext bytes.Buffer) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, ciphertext.Bytes(), nil)
	if err != nil {
		log.Fatal("Error: unable to encrypt with Rsa.")
		return nil, err
	}

	return encryptedBytes, nil
}

// decrypt password using private key
func decryptWithPrivateKey(privKey *rsa.PrivateKey, encoded []byte) ([]byte, error) {
	// get raw data from encoded
	encryptedData, err := base64.StdEncoding.DecodeString(string(encoded))
	if err != nil {
		log.Fatal("Error: unable to decode encoded data.")
		return nil, err
	}
	
	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, encryptedData, nil)
	if err != nil {
		log.Fatal("Error: unable to decrypt bytes data.")
		return nil, err
	}
	return decryptedBytes, nil
}

func decryptBox(text bytes.Buffer) ([]byte, error) {
	gobs := gob.NewDecoder(&text)
	var encPayLoad EncryptPayload
	err := gobs.Decode(&encPayLoad)
	if err != nil {
		log.Fatal("Error: unable to decode payload.")
		return nil, err
	}

	var nonce [NonceSize]byte 
	copy(nonce[:], encPayLoad.Ciphertext[:NonceSize])
	
	out, ok := secretbox.Open(nil, encPayLoad.Ciphertext[:NonceSize], encPayLoad.Nonce, encPayLoad.Key)
	if !ok {
		return nil, errors.New("unable to process secret box")
	}

	return out, nil
}

func main() {
	examplePassword := "password123"
	fmt.Printf("password:\t %v \n", examplePassword)

	// round 1 encryption
	ciphertext, key, nonce, err := encryptPassword([]byte(examplePassword))
	if err != nil {
		log.Fatal((err))
	}
	fmt.Printf("ciphertext, key, nonce:\t %v \n", ciphertext, key, nonce)

	// encode ciphertext, nonce, key into single payload to use for second round encryption
	var encPayLoad bytes.Buffer
	payload := gob.NewEncoder(&encPayLoad)
	err = payload.Encode(EncryptPayload{key, nonce, ciphertext})
	if err != nil {
		log.Fatal((err))
	}
	fmt.Printf("encrypted payload:\t %v \n", encPayLoad)

	// assymetric rsa encryption (encryption round 2)
	pubKey, privKey, err := genRsaKeys()
	if err != nil {
		log.Fatal((err))
	}
	fmt.Printf("pubkey:\t %v \n privatekey:\t %v \n", pubKey, privKey)

	encoded, err := encryptWithPublicKey(pubKey, encPayLoad)
	if err != nil {
		log.Fatal((err))
	}
	fmt.Printf("encoded password:\t %v \n", encoded)
	// save format for db or json
	// db.Set([]byte(key), []byte(base64.StdEncoding.EncodeToString(encoded)))
	fmt.Printf("base 64 encoded password:\t %v \n", []byte(base64.StdEncoding.EncodeToString(encoded)))

	//for decryption now all steps in reverse:
	unwrapped, err := decryptWithPrivateKey(privKey, []byte(base64.StdEncoding.EncodeToString(encoded)))
	if err != nil {
		log.Fatal((err))
	}
	var unwrap bytes.Buffer
	unwrap.WriteString(string(unwrapped))
	decrypted, err := decryptBox(unwrap)
	if err != nil {
		log.Fatal((err))
	}
	fmt.Printf("here is your original password:\t %v \n", decrypted)

}

