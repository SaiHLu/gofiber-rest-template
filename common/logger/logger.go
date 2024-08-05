package logger

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/fatih/color"
)

type customLogger struct {
	infoLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
}

var logger *customLogger

func init() {
	logger = &customLogger{
		infoLogger:  log.New(os.Stdout, color.BlueString("INFO: "), log.LstdFlags|log.Ldate),
		debugLogger: log.New(os.Stdout, color.YellowString("DEBUG: "), log.LstdFlags|log.Ldate),
		errorLogger: log.New(os.Stdout, color.RedString("ERROR: "), log.LstdFlags|log.Ldate),
	}
}

func Info(message string) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		file = filepath.Base(file)
		logger.infoLogger.Println(color.BlueString(file), color.BlueString("%d", line), color.BlueString(message))
	} else {
		logger.infoLogger.Println(message)
	}
}

func Debug(message string) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		file = filepath.Base(file)
		logger.debugLogger.Println(color.YellowString(file), color.YellowString("%d", line), color.YellowString(message))
	} else {
		logger.debugLogger.Println(message)
	}
}

func Error(message string) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		file = filepath.Base(file)
		logger.errorLogger.Println(color.RedString(file), color.RedString("%d", line), color.RedString(message))
	} else {
		logger.errorLogger.Println(message)
	}
}
