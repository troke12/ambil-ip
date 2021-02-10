package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

func get_ip() {
	kustom := "101.0.4.6"
	url := "http://ip.zxq.co/" + kustom + "/country"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	if strings.Contains(string(body), "ID"){
		fmt.Printf("\nBenar\n")
	} else {
		fmt.Printf("\nSalah\n")
	}
	os.Exit(0)
}

func main() {
	get_ip()
}
