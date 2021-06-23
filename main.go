package main

import (
	"bytes"
	"fmt"
	"github.com/gesquive/cli"
	"io"
	"log"
	"net/http"
	"regexp"
)


func getPage(url string) (contents string, err error) {
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

func getFastToken() (token string) {
	baseURL := "https://fast.com"

	fastBody, _ := getPage(baseURL)

	// Extract the app script url
	re := regexp.MustCompile("app-.*\\.js")
	scriptNames := re.FindAllString(fastBody, 1)

	scriptURL := fmt.Sprintf("%s/%s", baseURL, scriptNames[0])
	cli.Debug("trying to get fast api token from %s", scriptURL)

	// Extract the token
	scriptBody, _ := getPage(scriptURL)

	re = regexp.MustCompile("token:\"[[:alpha:]]*\"")
	tokens := re.FindAllString(scriptBody, 1)

	if len(tokens) > 0 {
		token = tokens[0][7 : len(tokens[0])-1]
		cli.Debug("token found: %s", token)
	} else {
		cli.Warn("no token found")
	}

	return
}

func main()  {
	mytoken := getFastToken()
	log.Println(mytoken)
}