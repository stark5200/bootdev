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