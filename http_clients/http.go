package main

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

func prettify(data string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(data), "", "  ")
	if err != nil {
		return "", fmt.Errorf("error indenting JSON: %w", err)
	}
	return prettyJSON.String(), nil
}

package main

import (
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
	return res, nil
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
Assignment
There is a bug in the getItemData function! It's returning the entire http.Response instead of the data from the body (a slice of bytes). Fix it so that it returns []byte.

Use io.ReadAll to read the .Body of the response.
Return the resulting []byte
*/