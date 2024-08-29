package main

import "unicode"

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

/*
Len and Cap Review
The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself. The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume. Here is a function to append data to a slice. If the data exceeds the capacity, the slice is reallocated. The resulting slice is returned. The function uses the fact that len and cap are legal when applied to the nil slice, and return 0.

Referenced from Effective Go

func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]byte, (l+len(data))*2)
        // The copy function is predeclared and works for any slice type.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}
*/

/// Variadic 

func sum(nums ...int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

/*
Variadic
Many functions, especially those in the standard library, can take an arbitrary number of final arguments. This is accomplished by using the "..." syntax in the function signature.

A variadic function receives the variadic arguments as a slice.

func concat(strs ...string) string {
    final := ""
    // strs is just a slice of strings
    for i := 0; i < len(strs); i++ {
        final += strs[i]
    }
    return final
}

func main() {
    final := concat("Hello ", "there ", "friend!")
    fmt.Println(final)
    // Output: Hello there friend!
}
Copy icon
The familiar fmt.Println() and fmt.Sprintf() are variadic! fmt.Println() prints each element with space delimiters and a newline at the end.

func Println(a ...interface{}) (n int, err error)
Copy icon
Spread operator
The spread operator allows us to pass a slice into a variadic function. The spread operator consists of three dots following the slice in the function call.

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}
Copy icon
Assignment
We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.

Complete the sum function to return the sum of all inputs.

Take note of how the variadic inputs and the spread operator are used in the test suite.
*/

/*
Append
The built-in append function is used to dynamically add elements to a slice:

func append(slice []Type, elems ...Type) []Type
Copy icon
If the underlying array is not large enough, append() will create a new underlying array and point the slice to it.

Notice that append() is variadic, the following are all valid:

slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)
Copy icon
Assignment
We've been asked to total costs per day, in a given period.

Complete the getCostsByDay() function using the append() function. It accepts a slice of cost structs and returns a slice of float64s, where each float64 represents the total cost for a day.

The length of the returned slice should be inferred from the highest numbered day field. Some days may have multiple costs, while others may have none. Include all actual and implied totals in the returned slice, starting with day '0'. Use the append() function to dynamically increase the size of the returned slice when needed.

Example
Given this input:

[]cost{
    {0, 4.0},
    {1, 2.1},
    {5, 2.5},
    {1, 3.1},
}
Copy icon
We expect this result:

[]float64{
    4.0, // first day
    5.2, // 2.1 + 3.1
    0.0, // intermediate days with no costs
    0.0, // ...
    0.0, // ...
    2.5, // last day
}
*/
type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsPerDay := []float64{}
	for i := 0; i < len(costs); i++ {
		for len(costsPerDay) <= costs[i].day{
			costsPerDay = append(costsPerDay, 0.0)
		}
	}
	for j := 0; j < len(costsPerDay); j++ {
		for k := 0; k < len(costs); k ++ {
			if costs[k].day == j {
			costsPerDay[j] += costs[k].value
			}
		}
	} 
	return costsPerDay
}

/// Slice of slices

/*
Slice of slices
Slices can hold other slices, effectively creating a matrix, or a 2D slice.

rows := [][]int{}
Copy icon
Assignment
We support various visualization dashboards on Textio that display message analytics to our users. The UI for our graphs and charts is built on top of a grid system. Let's build some grid logic.

Complete the createMatrix function. It takes a number of rows and columns and returns a 2D slice of integers. The value of each cell is i * j where i and j are the indexes of the row and column respectively.

For example, a 10x10 matrix would look like this:

[0 0 0 0 0 0 0 0 0 0]
[0 1 2 3 4 5 6 7 8 9]
[0 2 4 6 8 10 12 14 16 18]
[0 3 6 9 12 15 18 21 24 27]
[0 4 8 12 16 20 24 28 32 36]
[0 5 10 15 20 25 30 35 40 45]
[0 6 12 18 24 30 36 42 48 54]
[0 7 14 21 28 35 42 49 56 63]
[0 8 16 24 32 40 48 56 64 72]
[0 9 18 27 36 45 54 63 72 81]
*/

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for i := 0; i < rows; i++ {
		matrix = append(matrix, []int{})
		for j := 0; j < cols; j++ {
			matrix[i] = append(matrix[i], i*j)
		}
	}
	return matrix
}

func contains(s []string, word string) (bool, int) {
	for index, i := range s {
			if word == i {
						return true, index
			}
	}
	return false, 0
}

func indexOfFirstBadWord(msg []string, badWords []string) int {
for i, word := range msg {
	result, _ := contains(badWords, word)
	if result == true {
		return i
	}
}
return -1
}

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

// Don't touch above this line

func filterMessages(messages []Message, filterType string) []Message {
	filtered := []Message{}
	for _, msg := range messages {
		if msg.Type() == filterType {
			filtered = append(filtered, msg)
		}
	}
	return filtered
}

//// Password strength

/*
Password Strength
As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:

At least 5 characters long but no more than 12 characters.
Contains at least one uppercase letter.
Contains at least one digit.
Assignment
Implement the isValidPassword function. Use a loop to inspect each character in the password string to check for its length, and the presence of an uppercase letter and a digit.
*/


func isValidPassword(password string) bool {
	password_runes := []rune(password)
	if len(password_runes) < 5 || len(password_runes) > 12 {
		return false
	}
	digit_count := 0
	upper_count := 0
	for _, char := range password_runes {
		if unicode.IsDigit(char) {
			digit_count++
		}
		if unicode.IsUpper(char) {
			upper_count++
		}
	}
	if digit_count > 0 && upper_count > 0 {
		return true
	}
	return false
}

/// Challenge 3 

/*
Mailio
Textio is launching a new email messaging product, "Mailio"!

Assignment
Fix the compile-time bug in the getFormattedMessages function. The function body is correct, but the function signature is not.
*/

// type fn func(string) string

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}

