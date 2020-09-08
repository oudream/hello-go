package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)
import "os"

func fsIsExist(fp *string) bool {
	if _, err := os.Stat(*fp); err == nil {
		return false
	} else {
		return true
	}
}

func helloIsExist1() {
	if _, err := os.Stat("file-exists.go"); err == nil {
		fmt.Printf("File exists\n");
	} else {
		fmt.Printf("File does not exist\n");
	}
}

func helloJion() string {
	const f = "application_default_credentials.json"
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), "gcloud", f)
	}
	wd, _ := os.Getwd()
	return filepath.Join(wd, ".config", "gcloud", f)
}

func main() {
	helloJion()
}