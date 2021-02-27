package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func dsSavePropertyBigImage(url,filename string) bool {
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Printf("dsSavePropertyBigImage -> Get Error1 (%s) Fail.\n", e)
		time.Sleep(time.Second * 3)
		response, e = http.Get(url)
		if e != nil {
			log.Printf("dsSavePropertyBigImage -> Get Error2 (%s) Fail.\n", e)
			return false
		}
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create(filename)
	if err != nil {
		log.Panicf("dsSavePropertyBigImage -> Open File Error(%s) Fail.\n", err)
		return false
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Panicf("dsSavePropertyBigImage -> Save File Error(%s) Fail.\n", err)
		return false
	}
	log.Printf("dsSavePropertyBigImage -> Save Url[%s] To File[%s] Success", url, filename)
	return true
}

func readFile(filename string) {

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	//s := string(dat)
	fmt.Printf(" file size is : %d \n", len(dat))
	//fmt.Print(s)

}

func main() {
	url := "https://img.alicdn.com/imgextra/i4/6000000003757/O1CN01XsbmmG1dckdwIH0I7_!!6000000003757-0-octopus.jpg"

	for i := 0; i < 100; i++ {
		filename := fmt.Sprintf("%d.jpg", time.Now().UnixNano()/1e6+int64(i))
		fmt.Println(filename)
		dsSavePropertyBigImage(url, filename)
		readFile(filename)
		_ = os.Remove(filename)
	}

}
