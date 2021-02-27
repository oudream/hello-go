package main

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

var (
	logger   *log.Logger
)

// 初始化日誌，創建新的LOGGER對象
func initLogger() bool {
	//logPath, err1 := filepath.Abs(path.Join(*logPath, "log"))
	p, _ := os.Getwd()
	logPath, err1 := filepath.Abs(path.Join(p, "log"))
	if err1 != nil {
		fmt.Printf("initLogger  ERROR! 1 : %s \n", err1)
		return false
	}

	logf, err2 := rotatelogs.New(
		logPath+".%Y%m%d%H%M",
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithRotationCount(7),
	)
	if err2 != nil {
		fmt.Printf("initLogger ERROR! 2 : %s \n", err1)
		return false
	}
	writers := []io.Writer{
		//logf}
		logf, os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logger = log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Printf("initLogger complete!")
	return true
}
