package log

import (
	"fmt"
	"os"
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

func (l *Logger) write(level, msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	line := fmt.Sprintf("%s [%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		level,
		msg,
	)
	_, _ = l.file.WriteString(line)
}

func (l *Logger) Info(msg string) {
	l.write("INFO", msg)
}

func (l *Logger) Debug(msg string) {
	l.write("DEBUG", msg)
}

func (l *Logger) Error(msg string) {
	l.write("ERROR", msg)
}

func (l *Logger) Warn(msg string) {
	l.write("WARN", msg)
}
