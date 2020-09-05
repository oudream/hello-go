package main

import "fmt"
import "os"

func main() {
	if _, err := os.Stat("file-exists.go"); err == nil {
		fmt.Printf("File exists\n");
	} else {
		fmt.Printf("File does not exist\n");
	}
}