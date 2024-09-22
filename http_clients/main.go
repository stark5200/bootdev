package main

import (
	"fmt"
	"log"
	"bytes"
	"encoding/json"
)

const itemURL = "https://api.boot.dev/v1/courses_rest_api/learn-http/items"

func main() {
	items, err := getItemData(itemURL)
	if err != nil {
		log.Fatalf("error getting item data: %v", err)
	}
	prettyData, err := prettify(string(items))
	if err != nil {
		log.Fatalf("error prettifying data: %v", err)
	}
	fmt.Println(prettyData)
}