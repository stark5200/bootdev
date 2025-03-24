/*package main

import (
	"fmt"
	"io"
	"os"
	"strings"

)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("error reading messages")
		return
	}
	defer file.Close()
	buffer := make([]byte, 8)
	var currentLine string // Stores the accumulated line

	for {
		// Read up to 8 bytes into the buffer
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break // Exit loop on end of file
			}
			fmt.Println("Error reading file:", err)
			return
		}

		// Convert the read bytes to a string and split on newlines
		parts := strings.Split(string(buffer[:n]), "\n")

		// Process all but the last part
		for i := 0; i < len(parts)-1; i++ {
			fmt.Printf("read: %s\n", currentLine+parts[i])
			currentLine = "" // Reset the current line
		}

		// The last part remains in currentLine for the next iteration
		currentLine += parts[len(parts)-1]
	}

	// Print any remaining content as the last line
	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}
}
	*/

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// getLinesChannel reads from an io.ReadCloser in 8-byte chunks and sends complete lines to a channel.
func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer f.Close()  // Close the file when done
		defer close(lines) // Close the channel when the goroutine finishes

		buffer := make([]byte, 8)
		var currentLine string

		for {
			n, err := f.Read(buffer)
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("Error reading:", err)
				return
			}

			parts := strings.Split(string(buffer[:n]), "\n")

			for i := 0; i < len(parts)-1; i++ {
				lines <- currentLine + parts[i]
				currentLine = ""
			}

			currentLine += parts[len(parts)-1]
		}

		// Send any remaining data as the last line
		if currentLine != "" {
			lines <- currentLine
		}
	}()

	return lines
}

func main() {
	// Open file
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Get lines channel
	linesChannel := getLinesChannel(file)

	// Read from the channel and print each line
	for line := range linesChannel {
		fmt.Printf("read: %s\n", line)
	}
}

/*
Files vs. Network
Files and network connections behave very similarly - that's why we started by simply reading and writing to files, then updated our code to be a bit more abstract (the getLinesChannel function) so that it can handle both. From the perspective of your code, files and network connections are both just streams of bytes that you can read from and write to.

All of a sudden, Go's io.Reader (and the very similar io.ReadCloser) and io.Writer interfaces make a lot more sense, right? They're designed to work with any type of stream, whether it's a file, a network connection, or something else entirely.

Pull vs. Push
When you read from a file, you're in control of the reading process. You decide:

When to read
How much to read
When to stop reading.
You pull data from the file.

When you read from a network connection, the data is pushed to you by the remote server. You don't have control over when the data arrives, how much arrives, or when it stops arriving. Your code has to be ready to receive it when it comes.
*/
