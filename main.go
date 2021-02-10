package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
)

type AmbilIP struct {
	IP            string `json:"ip"`
	City          string `json:"city"`
	Region        string `json:"region"`
	Country       string `json:"country"`
	CountryFull   string `json:"country_full"`
	Continent     string `json:"continent"`
	ContinentFull string `json:"continent_full"`
	Loc           string `json:"loc"`
	Postal        string `json:"postal"`
}

func perror(err error) {
	if err != nil {
		panic(err)
	}
}

func get_ip() {
	kustom := "101.0.4.6"
	url := "http://ip.zxq.co/" + kustom

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var data AmbilIP
	json.Unmarshal(body, &data)
	if strings.Contains(data.Country, "ID"){
		fmt.Printf("\nBenar\n")
	} else {
		fmt.Printf("\nSalah\n")
	}
	os.Exit(0)
}

func main() {
	get_ip()
}
