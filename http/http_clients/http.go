package main

/// URL Uniform Resource Locator
/// a URL represents a piece of information on some computer somewhere. We can get access to it by making a request, and reading the response that the server replies with.
/// Other communication protocols use URLs as well, (hence "Uniform Resource Locator"). That's why we need to be specific when we're making HTTP requests by prefixing the URL with http://

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getItemData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return data, nil
}

func getItems(url string) ([]Item, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var items []Item
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&items)
	if err != nil {
		fmt.Println("error decoding response body")
		return items, err
	}

	return items, err
}

func marshalAll[T any](items []T) ([][]byte, error) {
	var allData [][]byte
	for _, item := range items {
		data, err := json.Marshal(item)
		if err != nil {
		    return nil, fmt.Errorf("error creating request: %w", err)
		}
		allData = append(allData, data)
	}
	return allData, nil
}


func prettify(data string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(data), "", "  ")
	if err != nil {
		return "", fmt.Errorf("error indenting JSON: %w", err)
	}
	return prettyJSON.String(), nil
}

/*
net/http
In this course, we'll be using Go's standard net/http package and the http.Client to make HTTP requests. In fact, we've already been using it! The http.Get function uses the http.DefaultClient under the hood.

Making a Request
import (
	"fmt"
	"io"
	"net/http"
)

func getSpells() ([]byte, error) {
	res, err := http.Get("https://api.fantasyquest.com/spell")
	if err != nil {
		return []byte{}, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading response: %w", err)
	}
	return data, nil
}
Copy icon
We'll go in-depth on the various things happening here later, but let's cover some basics for now.

http.Get uses the http.DefaultClient to make a request to the given url
res is the HTTP response that comes back from the server
defer res.Body.Close() ensures that the response body is properly closed after reading. Not doing so can cause memory issues.
io.ReadAll reads the response body into a slice of bytes []byte called data

#Assignment
There is a bug in the getItemData function! It's returning the entire http.Response instead of the data from the body (a slice of bytes). Fix it so that it returns []byte.

Use io.ReadAll to read the .Body of the response.
Return the resulting []byte
*/

/// DNS



func getIPAddress(domain string) (string, error) {
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


func getResources(path string) []map[string]any {
	fullURL := "https://api.boot.dev"+path
	fmt.Println("Requesting URL:", fullURL)
	validUrl := string(path[0]) == "/"
	if !validUrl {
	    fmt.Printf("Invalid path: %s. Ensure it starts with '/'.\n", path)
	    return nil
	}

	res, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	defer res.Body.Close()

	var resources []map[string]any
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resources)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}

	return resources

	
}

func errIfNotHTTPS(URL string) error {
	url, err := url.Parse(URL)
	if err != nil {
		return err
	}
	if url.Scheme != "https" {
		return fmt.Errorf("URL scheme is not HTTPS: %s", URL)
	}
	return nil
}