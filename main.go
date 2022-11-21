package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	shortName  string
	autoDelete uint
)

type UrlRequestBody struct {
	Url        string `json:"url"`
	Name       string `json:"name"`
	AutoDelete uint   `json:"auto_delete"`
}

type UrlResponseBody struct {
	ShortId string `json:"short_id"`
}

func main() {
	flag.StringVar(&shortName, "name", "", "--name custom-url-name")
	flag.UintVar(&autoDelete, "auto-delete", 0, "--auto-delete 30")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("Usage: zatch {long_url}")
		os.Exit(1)
	}
	url := args[0]

	body := UrlRequestBody{
		Url: url,
	}

	if shortName != "" {
		body.Name = shortName
	}
	if autoDelete > 0 {
		body.AutoDelete = autoDelete
	}

	resp, err := createShortUrl(body)
	if err != nil {
		fmt.Println("Error occured while sending request: ", err.Error())
	}

	shortUrl := fmt.Sprintf("https://zatch.ru/%s", resp.ShortId)

	fmt.Printf("Url shortened: %s\n", shortUrl)
}

func createShortUrl(body UrlRequestBody) (*UrlResponseBody, error) {
	b, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("https://zatch.ru/api/shorten", "application/json", bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data *UrlResponseBody
	err = json.Unmarshal(respData, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
