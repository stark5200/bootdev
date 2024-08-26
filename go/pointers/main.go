package main

import "fmt"

/*
Introduction to Pointers
As we have learned, a variable is a named location in memory that stores a value. We can manipulate the value of a variable by assigning a new value to it or by performing operations on it. When we assign a value to a variable, we are storing that value in a specific location in memory.

x := 42
// "x" is the name of a location in memory. That location is storing the integer value of 42
Copy icon
A pointer is a variable
A pointer is a variable that stores the memory address of another variable. This means that a pointer "points to" the location of where the data is stored NOT the actual data itself.

The * syntax defines a pointer:

var p *int
Copy icon
The & operator generates a pointer to its operand.

myString := "hello"
myStringPtr := &myString
*/


type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	return fmt.Sprintf(`
To: %v
Message: %v
`, &m.Recipient, &m.Text)
}
