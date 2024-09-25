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
/*
Web Addresses Quiz
To recap, a domain name is part of a URL. It's the part that tells the computer where the server is located on the internet by being converted into a numerical IP address.

We'll cover exactly how an IP address is used by your computer to find a path to the server in a later course. For now, it's just important to understand that an IP address is what your computer is using at a lower level to communicate on a network.

Deploying a real website to the internet is actually quite simple. It involves only a couple of steps:

Create a server that hosts your website files and connect it to the internet
Acquire a domain name
Connect the domain name to the IP address of your server
Your server is accessible via the internet!

*/