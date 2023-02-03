package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	EMERGENCY = "EMERGENCY"
	ALERT     = "ALERT"
	CRITICAL  = "CRITICAL"
	ERROR     = "ERROR"
	WARNING   = "WARNING"
	NOTICE    = "NOTICE"
	INFO      = "INFO"
	DEBUG     = "DEBUG"
)

func EmergencyLog(header string, message string) {
	log(header, message, EMERGENCY)
}

func AlertLog(header string, message string) {
	log(header, message, ALERT)
}

func CriticalLog(header string, message string) {
	log(header, message, CRITICAL)
}

func ErrorLog(header string, message string) {
	log(header, message, ERROR)
}

func WarningLog(header string, message string) {
	log(header, message, WARNING)
}

func NoticeLog(header string, message string) {
	log(header, message, NOTICE)
}

func InfoLog(header string, message string) {
	log(header, message, INFO)
}

func DebugLog(header string, message string) {
	log(header, message, DEBUG)
}

func log(header string, message string, level string) {
	logFile := findOrCreateLogFile()
	defer logFile.Close()

	timeNow := time.Now()
	_, err := logFile.WriteString(
		"[" + timeNow.Format(`02-01-2006 15:04:05`) + "] " +
			"[level: " + level + "] " +
			"[header: " + header + "]\n" +
			message + "\n",
	)
	if err != nil {
		fmt.Printf("Ошибка записи в файл логирования: %s\n", err.Error())
	}
}

func findOrCreateLogFile() *os.File {
	var logFile *os.File
	var err error

	year, month, day := time.Now().Date()
	logFilePath := fmt.Sprintf("storage/logs/date-%v-%s-%v.log", year, month, day)
	files, _ := filepath.Glob(logFilePath)
	isFind := len(files) > 0

	if isFind {
		logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY, 0666)
	} else {
		logFile, err = os.Create(logFilePath)
	}

	if err != nil {
		fmt.Printf("Ошибка создания/загрузки файла логирования: %s\n", err.Error())
	}

	return logFile
}
