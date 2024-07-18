package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		message       string
		formatter     func(string) string
		formatterName string
		expected      string
	}
	tests := []testCase{
		{"hello", addExclamation, "addExclamation", "TEXTIO: hello!!!"},
		{"hello there", addPeriod, "addPeriod", "TEXTIO: hello there..."},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{"moor der ehT", reverseString, "reverseString", "TEXTIO: The red room"},
		}...)
	}

	for _, test := range tests {
		if output := reformat(test.message, test.formatter); output != test.expected {
			t.Errorf(`Test Failed: (%v, %v) =>
  expected: %v
  got: %v
`,
				test.message, test.formatterName, test.expected, output)
		} else {
			fmt.Printf(`Test Passed: (%v, %v) =>
  expected: %v
  got: %v
`,
				test.message, test.formatterName, test.expected, output)
		}
		fmt.Println("--------------------------------")
	}
}

func addPeriod(s string) string {
	return s + "."
}

func addExclamation(s string) string {
	return s + "!"
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
