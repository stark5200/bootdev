package main

import (
	"fmt"
	"log"
)

const itemURL = "https://api.boot.dev/v1/courses_rest_api/learn-http/items"

func main() {
	items, err := getItemData("")
	if err != nil {
		log.Fatalf("error getting item data: %v", err)
	}
	prettyData, err := prettify(string(items))
	if err != nil {
		log.Fatalf("error prettifying data: %v", err)
	}
	fmt.Println(prettyData)
}
