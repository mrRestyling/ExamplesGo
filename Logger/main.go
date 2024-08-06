package main

import (
	"log"
	"os"
)

var (
	l Loggers
)

func Init() {

	flags := log.LstdFlags | log.Lshortfile

	fileInfo, _ := os.OpenFile("log_info.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	fileWarn, _ := os.OpenFile("log_warn.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	fileErr, _ := os.OpenFile("log_err.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	logInfo := log.New(fileInfo, "INFO:\t", log.LstdFlags)
	logWarn := log.New(fileWarn, "WARN:\t", flags)
	logErr := log.New(fileErr, "ERR:\t", flags)

	l = Loggers{
		logInfo: logInfo,
		logWarn: logWarn,
		logErr:  logErr,
	}

	// logInfo.Println("Information")
	// logWarn.Println("Warning!")
	// logErr.Println("ОШИБКА!")
}

type Loggers struct {
	logInfo *log.Logger
	logWarn *log.Logger
	logErr  *log.Logger
}

func main() {

	Init()
	l.info("Какая-то информация", 007)
	l.warn("Внимание!")
	l.err("Пример ошибки")
}

func (l *Loggers) info(v ...interface{}) {
	l.logInfo.Println(v...)
}

func (l *Loggers) warn(v ...interface{}) {
	l.logWarn.Println(v...)
}
func (l *Loggers) err(v ...interface{}) {
	l.logErr.Println(v...)
}
