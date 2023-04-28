package goer

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Logger struct {
	filename string
	file     *os.File
}

func NewLogger(filename string) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %s", err)
	}
	return &Logger{filename: filename, file: file}, nil
}

func (l *Logger) log(level string, message string) {
	timestamp := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf("[%s][%s] %s", level, timestamp, message)
	fmt.Println(msg)
	log.Println(msg)
	l.file.WriteString(msg + "\n")
}
func (l *Logger) Info(message string) {
	l.log("INFO", message)
}
func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}
func (l *Logger) Warn(message string) {
	l.log("WARN ", message)
}
func (l *Logger) Close() error {
	err := l.file.Close()
	if err != nil {
		return fmt.Errorf("failed to close log file: %s", err)
	}
	return nil
}
