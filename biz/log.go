package biz

import (
	"io"
	"log"
	"os"
)

var (
	LogHandle *log.Logger
	Debug     bool = true
)

func init() {
	file := "logs/testmgmt.log"
	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	if Debug {
		writers := []io.Writer{
			logFile,
			os.Stdout}
		fileAndStdoutWriter := io.MultiWriter(writers...)
		LogHandle = log.New(fileAndStdoutWriter, "[TestMgmt]", log.LstdFlags|log.Lshortfile)
	} else {
		LogHandle = log.New(logFile, "[TestMgmt]", log.LstdFlags|log.Lshortfile)
	}
	return
}
