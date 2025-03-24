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
