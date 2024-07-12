package main

import "fmt"

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
     // represents a Unicode code point

// float32 float64

// complex64 complex128

func main() {


	// initialize variables here
	var smsSendingLimit int = 0 
	var costPerSMS float64
	var hasPermission bool = false
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	messagesFromDoris := []string{
		"You doing anything later??",
		"Did you get my last message?",
		"Don't leave me hanging...",
		"Please respond I'm lonely!",
	}
	numMessages := float64(len(messagesFromDoris))
	costPerMessage := .02

	// don't touch above this line

	totalCost := costPerMessage * numMessages

	// don't touch below this line

	fmt.Printf("Doris spent %.2f on text messages today\n", totalCost)

	// declare here
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"
	result := true

	// don't edit below this line
	fmt.Println(messageStart, age, messageEnd, result)
}



func second() {
	penniesPerText := 2.0

	// don't edit below this line
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

	accountAge := 2.6

	// create a new "accountAgeInt" here
	// it should be the result of casting "accountAge" to an integer

	fmt.Println("Your account has existed for", accountAge, "years")

	// fmt.Sprintf() returns
	// %v variable, %s string, %d number, %.2f float with precision, %t boolean ...look for more
	const name = "Saul Goodman"
	const openRate = 30.5

	msg := fmt.Sprintf("Hi %v, your openrate is %.1f percent\n", name, openRate)

	// don't edit below this line

	fmt.Print(msg)
	//function examoples
  //func addToDatabase(hp, damage int) {
  // ...}
	//func addToDatabase(hp, damage int, name string) {
  // ?}
	//func addToDatabase(hp, damage int, name string, level int) {
  // ?}

	
}

