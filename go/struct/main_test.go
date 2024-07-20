package main

import (
	"fmt"
	"testing"
)

func TestGetBasicAuth(t *testing.T) {
	tests := []struct {
		auth     authenticationInfo
		expected string
	}{
		{authenticationInfo{"Google", "12345"}, "Authorization: Basic Google:12345"},
		{authenticationInfo{"Bing", "98765"}, "Authorization: Basic Bing:98765"},
	}
	if withSubmit {
		tests = append(tests, struct {
			auth     authenticationInfo
			expected string
		}{authenticationInfo{"DDG", "76921"}, "Authorization: Basic DDG:76921"})
	}

	for _, test := range tests {
		if output := test.auth.getBasicAuth(); output != test.expected {
			t.Errorf(`Test Failed:
%+v
-->
  expected: %s
  actual: %s
`, test.auth, test.expected, output)
		} else {
			fmt.Printf(`Test Passed:
%+v
-->
  expected: %s
  actual: %s
`, test.auth, test.expected, output)
			fmt.Println("==============================")
		}
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
