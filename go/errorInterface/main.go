package main

import (
	"fmt"
	"errors"
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

func getSMSErrorString(cost float64, recipient string) (s string) {
	s = fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent", cost, recipient)
	return
}


// Error Interface 3

/*
Because errors are just interfaces, you can build your own custom types that implement the error interface. Here's an example of a userError struct that implements the error interface:

type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}
Copy icon
It can then be used as an error:

func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) {
        return userError{name: userName}
    }
    ...
}
*/

type divideError struct {
	dividend float64
}


func (e divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero", e.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}


// Error Package


func divide2(x, y float64) (float64, error) {
	if y == 0 {
		var err error = errors.New("no dividing by 0")
		return 0, err
	}
	return x / y, nil
}

/// Panic

/*
PANIC
As we've seen, the proper way to handle errors in Go is to make use of the error interface. Pass errors up the call stack, treating them as normal values:

func enrichUser(userID string) (User, error) {
    user, err := getUser(userID)
    if err != nil {
        // fmt.Errorf is GOATed: it wraps an error with additional context
        return User{}, fmt.Errorf("failed to get user: %w", err)
    }
    return user, nil
}
Copy icon
However, there is another way to deal with errors in Go: the panic function. When a function calls panic, the program crashes and prints a stack trace.

As a general rule, do not use panic!

The panic function yeets control out of the current function and up the call stack until it reaches a function that defers a recover. If no function calls recover, the goroutine (often the entire program) crashes.

func enrichUser(userID string) User {
    user, err := getUser(userID)
    if err != nil {
        panic(err)
    }
    return user
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()

    // this panics, but the defer/recover block catches it
    // a truly astonishingly bad way to handle errors
    enrichUser("123")
}
Copy icon
Sometimes new Go developers look at panic/recover, and think, "This is like try/catch! I like this"! Don't be like them.

I use error values for all "normal" error handling, and if I have a truly unrecoverable error, I use log.Fatal to print a message and exit the program.
*/

// import ( "errors" )

func validateStatus(status string) error {
	if len(status) == 0 {
		var emptyErr error = errors.New("status cannot be empty")
		return emptyErr
	}

	if len(status) > 140 {
		var fullErr error = errors.New("status exceeds 140 characters")
		return fullErr
	}

	return nil
}
