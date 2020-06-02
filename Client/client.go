package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://13.231.145.155/hi"

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
}
