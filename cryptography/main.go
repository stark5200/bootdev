package main

import (
	"crypto/aes"
  "crypto/des"
	"crypto/cipher"
  "crypto/ecdsa"
  "crypto/elliptic"
  "crypto/rsa"
	"crypto/sha256"
  "crypto/rand"
  //"math/rand"
	"encoding/hex"
	"encoding/binary"
  "bytes"
  "strings"
  "errors"
  "math"
  "math/big"
	"fmt"
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

func crypt(textCh, keyCh <-chan byte, result chan<- byte) {
	defer close(result)
	for {
		textByte, textOk := <-textCh
		keyByte, keyOk := <-keyCh
		if !textOk || !keyOk {
			// means channel is closed
			return
		}
		result <- textByte ^ keyByte
	}
}

func getBlockSize(keyLen, cipherType int) (int, error) {
	var block cipher.Block
	var err error
	
	key := make([]byte, keyLen)

	switch cipherType {
		case typeAES:
			block, err = aes.NewCipher(key) 
			if err != nil {
				return 0, err
			}
		case typeDES:
			block, err = des.NewCipher(key)
			if err != nil {
				return 0, err
			}
		default:
	        return 0, errors.New("invalid cipher type")
	}
	return block.BlockSize(), nil
}

func padWithZeros(block []byte, desiredSize int) []byte {
	num_of_zeroes := desiredSize - len(block)
	padding := make([]byte, num_of_zeroes)
	paddedBlock := append(block, padding...)
	return paddedBlock
}

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
	byteValue := byte(roundNumber)
	roundKey := [4]byte{}
	for i := range masterKey {
		roundKey[i] = masterKey[i] ^ byteValue
	}
	return roundKey
}

func xor(lhs, rhs []byte) []byte {
	res := []byte{}
	for i := range lhs {
		res = append(res, lhs[i]^rhs[i])
	}
	return res
}

func feistel(msg []byte, roundKeys [][]byte) []byte {
	
	length := len(msg)
	lhs := msg[:length/2]
	rhs := msg[length/2:]
	for _, roundKey := range roundKeys {
		h := sha256.New()
		h.Write(append(rhs, roundKey...))
		currentRhs := xor(lhs, h.Sum(nil)[:length/2])
		lhs = rhs
		rhs = currentRhs
	}
	return append(rhs, lhs...)
}

func padWithZeros(block []byte, desiredSize int) []byte {
	for len(block) < desiredSize {
		block = append(block, 0)
	}
	return block
}

func encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddedText := padMsg(plaintext, des.BlockSize)

	ciphertext := make([]byte, des.BlockSize+len(paddedText))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(ciphertext[des.BlockSize:], paddedText)

	return ciphertext, nil
}

func padMsg(plaintext []byte, blockSize int) []byte {
	padLength := len(plaintext) + (blockSize - (len(plaintext) % blockSize)) % blockSize
	return padWithZeros(plaintext, padLength)
}

func generateIV(length int) ([]byte, error) {
	iv := make([]byte, length)	
	_, err := rand.Read(iv)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return iv, err
}

func sBox(b byte) (byte, error) {
	//bits := make([]int, 4)
	//for i := range 4 {
	//	bits[i] = (b >> (3 - i)) & 1
	//}
	
	var lookupTable = map[byte]byte {
	    0b0000: 0b00,
	    0b0001: 0b10,
	    0b0010: 0b01,
	    0b0011: 0b11,
			
		0b0100: 0b10,
	    0b0101: 0b00,
	    0b0110: 0b11,
	    0b0111: 0b01,
			
		0b1000: 0b01,
	    0b1001: 0b11,
	    0b1010: 0b00,
	    0b1011: 0b10,
	
		0b1100: 0b11,
	    0b1101: 0b01,
	    0b1110: 0b10,
	    0b1111: 0b00,
	}

	output, ok := lookupTable[b]
	
	if !ok {
		return 0, errors.New("invalid input")
	}

	return output, nil
}

// aes decrypt

func decrypt(key, ciphertext, nonce []byte) (plaintext []byte, err error) {
	block, blockErr := aes.NewCipher(key)
	if err != nil {
		return []byte{}, blockErr
	}
	aesgcm, gcmErr := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, gcmErr
	}
	return aesgcm.Open(nil, nonce, ciphertext, nil)
}

func nonceStrength(nonce []byte) int {
	return int(math.Pow(2, float64(len(nonce)*8)))
}

// generate keys for assymetric encryption

func genKeys() (pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey, err error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return &privateKey.PublicKey, privateKey, err
}

func encrypt(pubKey *rsa.PublicKey, msg []byte) ([]byte, error) {
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, msg, nil)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

func generatePrivateNums(keysize int) (*big.Int, *big.Int) {
	p, _ := getBigPrime(keysize)
	q, _ := getBigPrime(keysize)
	return p, q
}

// Calculate n = p * q
func getN(p, q *big.Int) *big.Int {
	N := new(big.Int)
	N.Mul(p, q)
	return N
}

