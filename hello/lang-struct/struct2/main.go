package main

import "fmt"

type TrainData struct {
	sentence  string
	sentiment string
	weight    int
}

func main() {
	trainData := TrainData{
		sentence:  "a",
		sentiment: "b",
		weight:    1,
	}

	//trainData["weight"] = 2

	fmt.Println(trainData)
}
