package main

import "fmt"

func main() {
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
	penniesPerText := 2

	// don't edit below this line
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

	accountAge := 2.6

	// create a new "accountAgeInt" here
	// it should be the result of casting "accountAge" to an integer

	fmt.Println("Your account has existed for", accountAge, "years")
}

