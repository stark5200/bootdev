package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
	"log"
)

type User struct {
	Password string `json:"password"`
	EncryptedPassword string `json:"encryptedPassword"`
	PrivateKey string `json:"privateKey"`
}


func genRsaKeys() (pubKey *rsa.PublicKey, privKey *rsa.PrivateKey, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return &privateKey.PublicKey, privateKey, nil
}

//encryption using public key
func encryptWithPublicKey(publicKey *rsa.PublicKey , text []byte) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, text, nil)
	if err != nil {
		log.Fatal("Error: unable to encrypt with Rsa.")
		return nil, err
	}

	return encryptedBytes, nil
}

// decrypt password using private key
func decryptWithPrivateKey(privKey *rsa.PrivateKey, encoded []byte) ([]byte, error) {
	// get raw data from encoded
	
	decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, encoded, nil)
	if err != nil {
		log.Fatal("Error: unable to decrypt data.")
		return nil, err
	}
	return decryptedBytes, nil
}

// PrivateKeyToString converts an RSA private key to a PEM string
func PrivateKeyToString(privateKey *rsa.PrivateKey) (string, error) {
	// Convert the private key to DER format
	der := x509.MarshalPKCS1PrivateKey(privateKey)

	// Create a PEM block
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: der,
	}

	// Encode the PEM block as a string
	return string(pem.EncodeToMemory(pemBlock)), nil
}

// StringToPrivateKey converts a PEM string back into an RSA private key
func StringToPrivateKey(pemString string) (*rsa.PrivateKey, error) {
	// Decode the PEM block
	block, _ := pem.Decode([]byte(pemString))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	// Parse the DER-encoded key
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA private key: %v", err)
	}

	return privateKey, nil
}


func encrypt(normalText string) (encrypted string, pvKey *rsa.PrivateKey, err error)  {
	pubKey, privKey, genErr := genRsaKeys()
	if genErr != nil {
		log.Fatal((genErr))
		return "", nil, genErr
	}
	fmt.Printf("pubkey:\t %v \n privatekey:\t %v \n", pubKey, privKey)

	encoded, enErr := encryptWithPublicKey(pubKey, []byte(normalText))
	if err != nil {
		log.Fatal((err))
		return "", nil, enErr
	}
	fmt.Printf("encoded password:\t %v \n", encoded)
	return string(encoded), privKey, nil
}

func decrypt(encryptedText []byte, pemPrivKey string) string {

	// Deserialize the string back to a private key
		deserializedKey, err := StringToPrivateKey(pemPrivKey)
		if err != nil {
			log.Fatal((err))
			return ""
		}

	unwrapped, err := decryptWithPrivateKey(deserializedKey, encryptedText)
	if err != nil {
		log.Fatal((err))
		return ""
	}
	return string(unwrapped)
}

func main() {
	passwords := []string {"passwordForTwitch123", "passwordForDiscord456", "passwordForTwitter789", "passwordForFFIV012", "passwordForPayPal345", "passwordForMyCompany678", "passwordForExpo901"}

	encryptedPasswords := []string {}

	privateKeys := []string {}

	for _, password := range passwords {
		encrypted, privateKey, _ := encrypt(password)
		encryptedPasswords = append(encryptedPasswords, encrypted)

		// Serialize the private key to a string
		pemString, err := PrivateKeyToString(privateKey)
		if err != nil {
			fmt.Println("Error converting private key to string:", err)
			return
		}

		privateKeys = append(privateKeys, pemString)
	}

	fmt.Println("Elements in encrypted passwords:")
	for _, str := range encryptedPasswords {
		fmt.Println(str)
	}

	fmt.Println("/////////////////////////////////////////////////////////////")

	fmt.Println("Elements in private keys:")
	for _, str := range privateKeys {
		fmt.Println(str)
	}

	data := make(map[string]User)

	for index, str := range passwords {
		u := new(User)
    u.Password = str
		u.EncryptedPassword = encryptedPasswords[index]
		u.PrivateKey = privateKeys[index]
		data[string(index)] = User{Password: u.Password, EncryptedPassword: u.EncryptedPassword, PrivateKey: u.PrivateKey}

	}

	// Convert the data to JSON format
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error creating JSON:", err)
	}

	// Save JSON to a file
	savePath := "users.json"
	err = os.WriteFile(savePath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Printf("User credentials saved to %s\n", savePath)

}

