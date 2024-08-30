package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// Don't touch below this line

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func main() {
	test("Hello there Kaladin!")
	test("Hi there Shallan!")
	test("Hey there Dalinar!")
}


/// Channels

/*
Channels
Channels are a typed, thread-safe queue. Channels allow different goroutines to communicate with each other.

Create a channel
Like maps and slices, channels must be created before use. They also use the same make keyword:

ch := make(chan int)
Copy icon
Send data to a channel
ch <- 69
Copy icon
The <- operator is called the channel operator. Data flows in the direction of the arrow. This operation will block until another goroutine is ready to receive the value.

Receive data from a channel
v := <-ch
Copy icon
This reads and removes a value from the channel and saves it into the variable v. This operation will block until there is a value in the channel to be read.

Blocking and deadlocks
A deadlock is when a group of goroutines are all blocking so none of them can continue. This is a common bug that you need to watch out for in concurrent programming.

Assignment
Run the program. You'll see that it deadlocks and never exits. The sendIsOld function is trying to send on a channel, but no other goroutines are running that can accept the value from the channel.

Fix the deadlock by spawning a goroutine to send the "is old" values.
*/

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

// don't touch below this line

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

/*
Channels 2
Empty structs are often used as a unary value. Sometimes, we don't care what is passed through a channel. We care when and if it is passed.

We can block and wait until something is sent on a channel using the following syntax

<-ch
Copy icon
This will block until it pops a single item off the channel, then continue, discarding the item.

Assignment
Our Mailio server isn't able to boot up until it receives the signal that its databases are all online, and it learns about them being online by waiting for tokens (empty structs) on a channel.

Run the code. It never exits! The channel passed to waitForDBs stays blocked.

Fix the waitForDBs function. It should pause execution until it receives a token for every database from the dbChan channel. Each time waitForDBs reads a token, the getDBsChannel goroutine will print a message to the console for you. The succinctly named numDBs input is the total number of databases. Look at the test code to see how these functions are used so you can understand the control flow.

*/


func waitForDBs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

// don't touch below this line

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}
