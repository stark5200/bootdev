package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/binary"
  "bytes"
  "strings"
  "errors"
  "math"
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

func getHexString(b []byte) string {
	output := ""
	if len(b) > 0 {
		for i := 0; i < len(b); i++ { 
	        output += fmt.Sprintf("%02x", b[i]) 
			if (i < len(b) - 1) {
				output += ":"
			}
	    } 
	}
	return output
}

func getBinaryString(b []byte) string {
	output := ""
	if len(b) > 0 {
		for i := 0; i < len(b); i++ { 
	        output += fmt.Sprintf("%08b", b[i]) 
			if (i < len(b) - 1) {
				output += ":"
			}
	    }
	}
	return output
}

func getHexBytes(s string) ([]byte, error) {
	values := strings.Join(strings.Split(s, ":"), "")
	values2, err := hex.DecodeString(values)
	return values2, err
}

func intToBytesCustom(n int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, int32(n)) // Use int32 or int64 as needed
	if err != nil {
		fmt.Println("Error:", err)
	}
	return buf.Bytes()
}

func alphabetSize(numBits int) float64 {
	return math.Pow(2, float64(numBits))
}

// Helper function: crypt performs XOR-based encryption/decryption
func crypt(dat, key []byte) []byte {
	final := []byte{}
	for i, d := range dat {
		final = append(final, d^key[i])
	}
	return final
}

// Helper function: intToBytes converts an integer to a 3-byte slice (little-endian)
func intToBytes(num int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int64(num))
	if err != nil {
		return nil
	}
	bs := buf.Bytes()
	if len(bs) > 3 {
		return bs[:3]
	}
	return bs
}

func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	limit := 1 << 24;
	for i := 0; i < limit; i++ {
		currentKey := intToBytes(i);
		if decrypted == string(crypt(encrypted, currentKey)) {
			return currentKey, nil
		}
	}
	return nil, errors.New("something went wrong")
}


func encrypt(plaintext string, key int) string {
	return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return crypt(ciphertext, -key)
}

func crypt(text string, key int) string {
	cryptText := ""
	for _, c := range text {
		cryptText += getOffsetChar(c, key)
	}
	return cryptText
}

func getOffsetChar(c rune, offset int) string {

	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	alphabetLength := len(alphabet)	
	char := string(c)
	
	if strings.Contains(alphabet, char) {
		originalIndex := strings.Index(alphabet, char)

		newIndex := (originalIndex + offset) % alphabetLength
		
		if newIndex < 0 {
			newIndex += alphabetLength
		}

		return string(alphabet[newIndex])
	}
	return char
}

func crypt(plaintext, key []byte) []byte {
	result := []byte{}
	for i, v := range plaintext {
		result = append(result, v ^ key[i]) // same as plaintext[i] ^ key[i], ^ is xor
	}
	return result
}
