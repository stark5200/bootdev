package main

import (
	"fmt"
	"time"
	"math/rand"
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

/*
Buffered Channels
Channels can optionally be buffered.

Creating a channel with a buffer
You can provide a buffer length as the second argument to make() to create a buffered channel:

ch := make(chan int, 100)
Copy icon
A buffer allows the channel to hold a fixed number of values before sending blocks. This means sending on a buffered channel only blocks when the buffer is full, and receiving blocks only when the buffer is empty.

Assignment
We want to be able to send emails in batches. A writing goroutine will write an entire batch of email messages to a buffered channel, and later, once the channel is full, a reading goroutine will read all of the messages from the channel and send them out to our clients.

Complete the addEmailsToQueue function. It should create a buffered channel with a buffer large enough to store all of the emails it's given. It should then write the emails to the channel in order, and finally return the channel.
*/

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

/// Closing Channels

/*
Closing channels in Go
Channels can be explicitly closed by a sender:

ch := make(chan int)

// do some stuff with the channel

close(ch)
Copy icon
Checking if a channel is closed
Similar to the ok value when accessing data in a map, receivers can check the ok value when receiving from a channel to test if a channel was closed.

v, ok := <-ch
Copy icon
ok is false if the channel is empty and closed.

Don't send on a closed channel
Sending on a closed channel will cause a panic. A panic on the main goroutine will cause the entire program to crash, and a panic in any other goroutine will cause that goroutine to crash.

Closing isn't necessary. There's nothing wrong with leaving channels open, they'll still be garbage collected if they're unused. You should close channels to indicate explicitly to a receiver that nothing else is going to come across.

Assignment
At Mailio we're all about keeping track of what our systems are up to with great logging and telemetry.

The sendReports function sends out a batch of reports to our clients and reports back how many were sent across a channel. It closes the channel when it's done.

Complete the countReports function. It should:

Use an infinite for loop to read from the channel:
If the channel is closed, break out of the loop
Otherwise, keep a running total of the number of reports sent
Return the total number of reports sent
*/

func countReports(numSentCh chan int) int {
	total := 0
	for {
		numSent, ok := <-numSentCh
		if !ok {
			break
		}
		total += numSent
	}
	return total
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}


/// Range

/*
Range
Similar to slices and maps, channels can be ranged over.

for item := range ch {
    // item is the next value received from the channel
}
Copy icon
This example will receive values over the channel (blocking at each iteration if nothing new is there) and will exit only when the channel is closed.

Assignment
It's that time again, Mailio is hiring and we've been assigned to do the interview. The Fibonacci sequence is Mailio's interview problem of choice. We've been tasked with building a small toy program we can use in the interview.

Complete the concurrentFib function. It should:

Create a new channel of ints
Call fibonacci concurrently
Use a range loop to read from the channel and append the values to a slice
Return the slice
*/

func concurrentFib(n int) []int {
	fibSlice := []int{}
	fibs := make(chan int, n)
	go fibonacci(n, fibs)
	for i := range fibs {
		fibSlice = append(fibSlice, i)
	}
	return fibSlice
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

/// Select

/*
Select
Sometimes we have a single goroutine listening to multiple channels and want to process data in the order it comes through each channel.

A select statement is used to listen to multiple channels at the same time. It is similar to a switch statement but for channels.

select {
case i, ok := <- chInts:
    fmt.Println(i)
case s, ok := <- chStrings:
    fmt.Println(s)
}
Copy icon
The first channel with a value ready to be received will fire and its body will execute. If multiple channels are ready at the same time one is chosen randomly. The ok variable in the example above refers to whether or not the channel has been closed by the sender yet.

Assignment
Complete the logMessages function.

Use an infinite for loop and a select statement to log the emails and sms messages as they come in order across the two channels. Add a condition to return from the function when one of the two channels closes, whichever is first.

Use the logSms and logEmail functions to log the messages.
*/



func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case s, ok := <- chSms:
			if !ok {
				return
			}
			logSms(s)
		case e, ok := <- chEmails:
			if !ok {
				return
			}
			logEmail(e)
		}
	}
}

// don't touch below this line

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

func test2(sms []string, emails []string) {
	fmt.Println("Starting...")

	chSms, chEmails := sendToLogger(sms, emails)

	logMessages(chEmails, chSms)
	fmt.Println("===============================")
}

func main2() {
	rand.Seed()
	test2(
		[]string{
			"hi friend",
			"What's going on?",
			"Welcome to the business",
			"I'll pay you to be my friend",
		},
		[]string{
			"Will you make your appointment?",
			"Let's be friends",
			"What are you doing?",
			"I can't believe you've done this.",
		},
	)
	test2(
		[]string{
			"this song slaps hard",
			"yooo hoooo",
			"i'm a big fan",
		},
		[]string{
			"What do you think of this song?",
			"I hate this band",
			"Can you believe this song?",
		},
	)
}

func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)
	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			done := make(chan struct{})
			s := sms[i]
			e := emails[i]
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
			go func() {
				time.Sleep(t1)
				chSms <- s
				done <- struct{}{}
			}()
			go func() {
				time.Sleep(t2)
				chEmails <- e
				done <- struct{}{}
			}()
			<-done
			<-done
			time.Sleep(10 * time.Millisecond)
		}
		close(chSms)
		close(chEmails)
	}()
	return chSms, chEmails
}

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
			return
		default:
			waitForData(logChan)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
