package main

import (
	"fmt"
	"log"
)

func main() {
	items, err := getItemData()
	if err != nil {
		log.Fatalf("error getting item data: %v", err)
	}

	// Don't edit above this line

	fmt.Println(string(items[:]))
}
