package log

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	mu   sync.Mutex
	file *os.File
}

// New - создание или открытие файла для записи
func New(path string) (*Logger, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return &Logger{file: f}, nil
}

func (l *Logger) Close() {
	l.file.Close()
}

func (l *Logger) write(level, msg string, fields map[string]any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var logLine strings.Builder
	logLine.WriteString(fmt.Sprintf("%s [%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		level,
		msg,
	))

	if len(fields) > 0 {
		for k, v := range fields {
			logLine.WriteString(fmt.Sprintf(" | %s=%v", k, v))
		}
	}

	logLine.WriteString("\n")

	_, _ = l.file.WriteString(logLine.String())
}

func (l *Logger) Info(msg string) {
	l.write("INFO", msg, nil)
}

func (l *Logger) Debug(msg string) {
	l.write("DEBUG", msg, nil)
}

func (l *Logger) Error(msg string) {
	l.write("ERROR", msg, nil)
}

func (l *Logger) Warn(msg string) {
	l.write("WARN", msg, nil)
}

func (l *Logger) InfoWithFields(msg string, fields map[string]any) {
	l.write("INFO", msg, fields)
}

func (l *Logger) ErrorWithFields(msg string, fields map[string]any) {
	l.write("ERROR", msg, fields)
}
