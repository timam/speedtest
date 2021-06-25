package utility

import (
	"bytes"
	"io"
	"net/http"
)

// GetPage Takes an url returns content of it's Body
func GetPage(url string) (contents string, err error) {
	// Create the string buffer
	buffer := bytes.NewBuffer(nil)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return contents, err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(buffer, resp.Body)
	if err != nil {
		return contents, err
	}
	contents = buffer.String()

	return
}

