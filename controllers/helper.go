package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJsonRespFromUrl(Url string, target interface{}, printBody bool) error {
	resp, err := myClient.Get(Url)
	if err != nil { return err }

	defer resp.Body.Close()

    if (printBody) {
        body, err := ioutil.ReadAll(resp.Body)
        if (err != nil) {
            fmt.Print(Url, err.Error())
        }
        fmt.Print(Url, string(body))
    }

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetStringRespFromUrl(Url string, printBody bool) (string, error) {
	resp, err := myClient.Get(Url)
	if err != nil { return "", err }

	defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if (err != nil) {
        fmt.Print(Url, err.Error())
    }

    if (printBody) {
        fmt.Print(Url, string(body))
    }

	return string(body), err
}