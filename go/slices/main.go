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
