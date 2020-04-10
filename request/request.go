package request

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var client = http.Client{
	Timeout: 5 * time.Second,
}

// GetJSON get json url
func GetJSON(url string) io.ReadCloser {
	r, httperr := client.Get(url)
	if httperr != nil {
		time.Sleep(1 * time.Minute)
		return GetJSON(url)
	}

	return r.Body
}

// Get gets a url and error checks, loops request on error
func Get(url string) (string, int) {
	r, httperr := client.Get(url)
	if httperr != nil {
		time.Sleep(3 * time.Second)
		return Get(url)
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return Get(url)
	}

	return string(body), r.StatusCode
}
