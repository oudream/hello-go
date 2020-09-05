package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

//var cfp = flag.String("cfp", "", "config file path")
//var lfp = flag.String("lfp", "", "log file path")
//var cp = flag.String("cp", "", "config path")
var lp = flag.String("lp", "", "log path")

func main() {
	flag.Parse()

	logFilePath, err := filepath.Abs(path.Join(*lp, "access_log"))

	logf, err := rotatelogs.New(
		logFilePath+".%Y%m%d%H%M",
		//rotatelogs.WithLinkName(logFilePath),
		//rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Minute),
		rotatelogs.WithRotationCount(3),
	)
	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}
	//log.SetOutput(logf)
	writers := []io.Writer{
		logf,
		os.Stdout}
	fileAndStdoutWriter := io.MultiWriter(writers...)
	// 创建新的log对象
	logger := log.New(fileAndStdoutWriter, "", log.Ldate|log.Ltime|log.Lshortfile)

	for i := 0; i < 10*60*5; i++ {
		f1 := 1.234*float64(i)
		logger.Printf("%d - %.2f --!", i , f1)
		time.Sleep(100 * time.Millisecond)
	}
}
