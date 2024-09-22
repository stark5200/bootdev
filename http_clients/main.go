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

func prettify2(data string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(data), "", "  ")
	if err != nil {
		return "", fmt.Errorf("error indenting JSON: %w", err)
	}
	return prettyJSON.String(), nil
}

/// JSON 

const itemList = `
[
	{
		 "id": 0, 
		 "name": "sword", 
		 "damage": 10.5, 
		 "equipped": false
	},
	{
		"id": 1, 
		"name": "shield", 
		"block": 5.5, 
		"equipped": true
	}
]
`

const playerObject = `
{ 
	"name": "Fudd", 
	"items": "spear and magic helmet", 
	"wife": "Brunhilde", 
	"power": 9000
}
`
