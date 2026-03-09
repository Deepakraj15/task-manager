package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/Deepakraj15/task-manager/internal/constants"
)

type Logger struct {
	file *os.File
	lock sync.Mutex
}

func GetLogger(filename string) (*Logger, error) {

	dir := constants.LOG_FILE_PATH
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create log directory: %v", err)
	}

	filePath := filepath.Join(dir, filename)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %v", err)
	}

	return &Logger{
		file: file,
	}, nil
}

func (l *Logger) write(level LogLevel, message string) {
	if l == nil {
		fmt.Println("Logger is nil. Message:", message)
		return
	}

	l.lock.Lock()
	defer l.lock.Unlock()

	logLine := fmt.Sprintf(
		"%s [%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		level,
		message,
	)
	fmt.Println(logLine)
	l.file.WriteString(logLine)
}

// Debug/Info/Error
func (l *Logger) Debug(msg string)          { l.write(LogLevel(Debug), msg) }
func (l *Logger) Info(msg string)           { l.write(LogLevel(Info), msg) }
func (l *Logger) Error(msg string, e error) { l.write(LogLevel(Error), msg+". ERROR -> "+e.Error()) }
