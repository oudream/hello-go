package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func helloUrl1()  {
	url := "http://192.168.241.32:8081/api/allpropertyinfos/0?begin=1598292000000&end=0"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(url, time.Now())
		log.Println(bodyString)
	}
}

func main() {
	//resp, err := http.Get("http://gobyexample.com")
	for {
		helloUrl1()
		time.Sleep(10)
	}
}