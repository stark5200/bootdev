package main

import (
	"fmt"
	"testing"
	"errors"
)
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


/// Pointer Receiver code

/*
Pointer Receiver Code
Methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.

type circle struct {
	x int
	y int
    radius int
}

func (c *circle) grow() {
    c.radius *= 2
}

func main() {
    c := circle{
        x: 1,
        y: 2,
        radius: 4,
    }

    // notice c is not a pointer in the calling function
    // but the method still gains access to a pointer to c
    c.grow()
    fmt.Println(c.radius)
    // prints 8
}
Copy icon
Assignment
Fix the bug in the code so that setMessage sets the message field of the given email structure, and the new value persists outside the scope of the setMessage method.
*/

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

// don't edit below this line

type email struct {
	message     string
	fromAddress string
	toAddress   string
}


/// Go Pointers benchmark


//	"fmt"
//	"testing"


type data struct {
	a, b, c, d, e, f, g, h, i, j int64
}

var globalPtr *data
var globalValue data

func newDataPtr(i int) *data {
	data := &data{int64(i), int64(i + 1), int64(i + 2), int64(i + 3), int64(i + 4), int64(i + 5), int64(i + 6), int64(i + 7), int64(i + 8), int64(i + 9)}
	return data
}

func newData(i int) data {
	data := data{int64(i), int64(i + 1), int64(i + 2), int64(i + 3), int64(i + 4), int64(i + 5), int64(i + 6), int64(i + 7), int64(i + 8), int64(i + 9)}
	return data
}

func BenchmarkProcessValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		globalValue = newData(i)
	}
	// use it to avoid compiler optimization
	fmt.Println(globalValue.a)
}

func BenchmarkProcessPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		globalPtr = newDataPtr(i)
	}
	// use it to avoid compiler optimization
	fmt.Println(globalPtr.a)
}

/// pointer Challenge 1

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

	func updateBalance(cP *customer, t transaction) error {
		customerBalance := cP.balance
		if (t.transactionType != transactionDeposit) && (t.transactionType != transactionWithdrawal) {
			return errors.New("unknown transaction type")
		}
		if t.transactionType == transactionWithdrawal {
			if customerBalance - t.amount < 0 {
				return errors.New("insufficient funds")
			}
			cP.balance = customerBalance - t.amount 
		}
		if t.transactionType == transactionDeposit {
			cP.balance = customerBalance + t.amount
		}
		return nil
	}
