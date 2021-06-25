package main

import (
	"fmt"
	"github.com/gesquive/cli"
	"github.com/timam/speedtest/utility"
	"regexp"
)


func getFastToken() (token string) {
	baseURL := "https://fast.com"

	fastBody, _ := utility.GetPage(baseURL)

	// Extract the app script url
	re := regexp.MustCompile("app-.*\\.js")
	scriptNames := re.FindAllString(fastBody, 1)

	scriptURL := fmt.Sprintf("%s/%s", baseURL, scriptNames[0])
	cli.Debug("trying to get fast api token from %s", scriptURL)

	// Extract the token
	scriptBody, _ := utility.GetPage(scriptURL)

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

func getDownloadURLs(token string, urlCount uint64) (urls []string) {

	url := fmt.Sprintf("https://api.fast.com/netflix/speedtest?https=true&token=%s&urlCount=%d", token, urlCount)
	jsonData, _ := utility.GetPage(url)

	re := regexp.MustCompile("(?U)\"url\":\"(.*)\"")
	reUrls := re.FindAllStringSubmatch(jsonData, -1)
	for _, arr := range reUrls{
		urls = append(urls, arr[1])
	}
	return
}

func main()  {
	token := getFastToken()
	urls := getDownloadURLs(token, 3)
	fmt.Println(urls)
}