package main

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