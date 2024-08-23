package main

/*
ARRAYS IN GO
Arrays are fixed-size groups of variables of the same type.

The type [n]T is an array of n values of type T

To declare an array of 10 integers:

var myInts [10]int
Copy icon
or to declare an initialized literal:

primes := [6]int{2, 3, 5, 7, 11, 13}
Copy icon
ASSIGNMENT
When our clients don't respond to a message, they can be reminded with up to 2 additional messages. getMessageWithRetries returns:

An array of 3 strings
An array of 3 integers
The strings are just the original messages structured as an array. The first is the primary message, the second is the first reminder, and the third is the last reminder. The integers represent the cost of sending each message.

The cost of a message is equal to the length of the message, plus the cost of any previous messages. For example:

"hello" costs 5
"world" costs 5, adding "hello" makes total cost 10 (5 + 5)
"!" costs 1, adding the previous messages makes total cost 11 (5 + 5 + 1)
*/

func getMessageWithRetries(primary, secondary, tertiary string) (messages [3]string, cost [3]int) {
	messages = [3]string{primary, secondary, tertiary}
	cost = [3]int{len(primary), len(primary)+len(secondary), len(primary)+len(secondary)+len(tertiary)}
	return
}

const (
	planFree = "free"
	planPro  = "pro"
)

type MyError struct{}

func (m *MyError) Error() string {
	return "unsupported plan"
}

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan ==  planPro {
		return messages[:], nil
	}
	if plan ==  planFree {
		return messages[:2], nil
	}
	return nil, &MyError{}
}

/*
Make
Most of the time we don't need to think about the underlying array of a slice. We can create a new slice using the make function:

// func make([]T, len, cap) []T
mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
mySlice := make([]int, 5)
Copy icon
Slices created with make will be filled with the zero value of the type.

If we want to create a slice with a specific set of values, we can use a slice literal:

mySlice := []string{"I", "love", "go"}
Copy icon
Notice the square brackets do not have a 3 in them. If they did, you'd have an array instead of a slice.

Length
The length of a slice is simply the number of elements it contains. It is accessed using the built-in len() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(len(mySlice)) // 3
Copy icon
Capacity
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice. It is accessed using the built-in cap() function:

mySlice := []string{"I", "love", "go"}
fmt.Println(cap(mySlice)) // 3
Copy icon
Generally speaking, unless you're hyper-optimizing the memory usage of your program, you don't need to worry about the capacity of a slice because it will automatically grow as needed.

Assignment
We send a lot of text messages at Textio, and our API is getting slow and unresponsive.

If we know the rough size of a slice before we fill it up, we can make our program faster by creating the slice with that size ahead of time so that the Go runtime doesn't need to continuously allocate new underlying arrays of larger and larger sizes. By setting the length, the slice can still be resized later, but it means we can avoid all the expensive resizing since we know what we'll need.

Complete the getMessageCosts() function. It takes a slice of messages and returns a slice of message costs.

Preallocate a slice for the message costs of the same length as the messages slice.
Fill the costs slice with costs for each message. The cost in the cost slice should correspond to the message in the messages slice at the same index. The cost of a message is the length of the message multiplied by 0.01.
*/

func getMessageCosts(messages []string) []float64 {
	messagesLength := len(messages)
	costSlice := make([]float64, messagesLength)
	for i := 0; i < messagesLength; i++ {
		costSlice[i] = float64(len(messages[i]))*0.01
	} 
	return costSlice
}

