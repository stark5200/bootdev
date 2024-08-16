package main

import (
	"fmt"
)


/*
Error Interface explanation

type error interface {
    Error() string
}
Copy icon
When something can go wrong in a function, that function should return an error as its last return value. Any code that calls a function that can return an error should handle errors by testing whether the error is nil.

// Atoi converts a stringified number to an integer
i, err := strconv.Atoi("42b")
if err != nil {
    fmt.Println("couldn't convert:", err)
    // because "42b" isn't a valid integer, we print:
    // couldn't convert: strconv.Atoi: parsing "42b": invalid syntax
    // Note:
    // 'parsing "42b": invalid syntax' is returned by the .Error() method
    return
}
// if we get here, then
// i was converted successfully

*/

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (cost int, err error) {
	cost1, err1 := sendSMS(msgToCustomer)
	if err1 != nil {
		return 0, err1
	}
	cost2, err2 := sendSMS(msgToSpouse)
	if err2 != nil {
		return 0, err2
	}
	cost = cost1 + cost2
	return 
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

func sendSMS2(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

/*

FORMATTING STRINGS REVIEW
A convenient way to format strings in Go is by using the standard library's fmt.Sprintf() function. It's a string interpolation function, similar to JavaScript's built-in template literals. The %v substring uses the type's default formatting, which is often what you want.

DEFAULT VALUES
const name = "Kim"
const age = 22
s := fmt.Sprintf("%v is %v years old.", name, age)
// s = "Kim is 22 years old."
Copy icon
The equivalent JavaScript code:

const name = 'Kim'
const age = 22
s = `${name} is ${age} years old.`
// s = "Kim is 22 years old."
Copy icon
ROUNDING FLOATS
fmt.Printf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
fmt.Printf("I am %.2f years old", 10.523)
// I am 10.52 years old

*/


package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) (s string) {
	s = fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent", cost, recipient)
	return
}
