package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

type transport struct {
	http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	resp, err = t.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	//b = bytes.Replace(b, []byte("server"), []byte("schmerver"), -1)
	body := ioutil.NopCloser(bytes.NewReader(b))
	resp.Body = body
	resp.ContentLength = int64(len(b))
	resp.Header.Set("Content-Length", strconv.Itoa(len(b)))

	logger.Println("------------------------ begin : -------------")
	// Loop over header names
	for name, values := range req.Header {
		// Loop over all values for the name.
		for _, value := range values {
			logger.Println(name, value)
		}
	}
	result := map[string]interface{}{}
	//_ = json.Unmarshal(req.Body, &result)
	//fmt.Println(result)
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&result)
	logger.Println(result)
	if err != nil {
		panic(err)
	}
	logger.Println("------------------------ end. -------------")
	return resp, nil
}


func main() {
	initLogger()
	target, err := url.Parse("http://192.168.5.109:9200")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = &transport{http.DefaultTransport}

	http.Handle("/", proxy)
	http.ListenAndServe(":9200", nil)
}
