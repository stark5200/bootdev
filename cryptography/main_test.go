package main

import (
	"fmt"
	"testing"
)

func TestDebugEncryptDecrypt(t *testing.T) {
	type testCase struct {
		masterKey string
		iv        string
		password  string
		expectedE string
		expectedD string
	}

	const masterKey = "kjhgfdsaqwertyuioplkjhgfdsaqwert"
	const iv = "1234567812345678"

	tests := []testCase{
		{masterKey, iv, "k33pThisPasswordSafe", encrypt("k33pThisPasswordSafe", masterKey, iv), "k33pThisPasswordSafe"},
		{masterKey, iv, "12345", encrypt("12345", masterKey, iv), "12345"},
		{masterKey, iv, "thePasswordOnMyLuggage", encrypt("thePasswordOnMyLuggage", masterKey, iv), "thePasswordOnMyLuggage"},
		{masterKey, iv, "pizza_the_HUt", encrypt("pizza_the_HUt", masterKey, iv), "pizza_the_HUt"},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{masterKey, iv, "aNewPassword", encrypt("aNewPassword", masterKey, iv), "aNewPassword"},
			{masterKey, iv, "edgeCaseTest", encrypt("edgeCaseTest", masterKey, iv), "edgeCaseTest"},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		encrypted, decrypted := debugEncryptDecrypt(test.masterKey, test.iv, test.password)
		if encrypted != test.expectedE || decrypted != test.expectedD {
			failCount++
			t.Errorf(`---------------------------------
Inputs:      masterKey: %v, iv: %v, password: %v
Expecting:   Encrypted: %v, Decrypted: %v
Actual:      Encrypted: %v, Decrypted: %v
Fail`, test.masterKey, test.iv, test.password, test.expectedE, test.expectedD, encrypted, decrypted)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:      masterKey: %v, iv: %v, password: %v
Expecting:   Encrypted: %v, Decrypted: %v
Actual:      Encrypted: %v, Decrypted: %v
Pass
`, test.masterKey, test.iv, test.password, test.expectedE, test.expectedD, encrypted, decrypted)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true