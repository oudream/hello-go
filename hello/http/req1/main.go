package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func helloUrl1()  {
	now := time.Now().UnixNano() / 1e6
	//url := "http://192.168.5.29:8080/web/search"
	url := "http://192.168.5.29:808/api/search"
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
	log.Println("cost time : " , time.Now().UnixNano() / 1e6 - now)
}

func main() {
	//resp, err := http.Get("http://gobyexample.com")
	helloUrl1()

	//for {
	//	helloUrl1()
	//	time.Sleep(10)
	//}
}