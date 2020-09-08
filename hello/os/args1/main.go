package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dirExe, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Args[0])
	fmt.Println(dirExe)
	dirPwd, err := os.Getwd()
	fmt.Println(dirPwd)
}
