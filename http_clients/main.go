package main

import (
	"fmt"
	"log"
	"io"
	"bytes"
	"encoding/json"
	"net/http/"
	"net/url"
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

/*
DNS
A "domain name" or "hostname" is just one portion of a URL. We'll get to the other parts of a URL later.

For example, the URL https://homestarrunner.com/toons has a hostname of homestarrunner.com. The https:// and /toons portions aren't part of the domain name -> IP address mapping that we've been talking about.

 The net/url package
The net/url package is part of Go's standard library. You can instantiate a URL struct using url.Parse:

parsedURL, err := url.Parse("https://homestarrunner.com/toons")
if err != nil {
	fmt.Println("error parsing url:", err)
	return
}
Copy icon
And then you can extract just the hostname:

parsedURL.Hostname()
// homestarrunner.com
Copy icon
 Assignment
Complete the getDomainNameFromURL function. Given a full URL, it should return the domain (or host) name. Simply return any potential errors.
*/

func getIPAddress2(domain string) (string, error) {
	url := fmt.Sprintf("https://cloudflare-dns.com/dns-query?name=%s&type=A", domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var dnsRes DNSResponse
	if err := json.Unmarshal(body, &dnsRes); err != nil {
		return "", fmt.Errorf("error unmarshalling json: %w", err)
	}

	if len(dnsRes.Answer) == 0 {
		return "", fmt.Errorf("no answer found")
	}

	return dnsRes.Answer[0].Data, nil
}

func getDomainNameFromURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("error parsing url:", err)
		return "", err
	}
	return parsedURL.Hostname(), err
}

// creating a new request
req, err := http.NewRequest("GET", "https://api.example.com/users", nil)
if err != nil {
	fmt.Println("error creating request: ", err)
	return
}

// setting a header on the new request
req.Header.Set("x-api-key", "123456789")

// making the request
client := http.Client{}
res, err := client.Do(req)
if err != nil {
	fmt.Println("error making request: ", err)
	return
}
defer res.Body.Close()

// reading a header from the response
header := res.Header.Get("last-modified")
fmt.Println("last modified: ", header)

// deleting a header from the response
res.Header.Del("last-modified")

func getContentType(res *http.Response) string {
	return res.Header.Get("Content-Type")
}


func getUsers(url string) ([]User, error) {
	fullURL := url + "?sort=experience"
	res, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var users []User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
