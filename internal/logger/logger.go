package logger

import (
	"io"
	"log"
	"os"
)

// LogLevel определяет уровни логирования.
type LogLevel int

// Перечисление уровней логирования.
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// Описание структуры логгера
type Logger struct {
	level   LogLevel
	logger  *log.Logger
	logFile *os.File
}

// NewLogger создает новый логер с заданным уровнем логирования и файлом логов.
func NewLogger(level LogLevel, logFilePath string) (*Logger, error) {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err

	}

	// Логи будут записываться как в файл, так и выводится на стандартный вывод.
	multiWriter := io.MultiWriter(file, os.Stdout)
	logger := log.New(multiWriter, "", log.Ldate|log.Ltime|log.Lshortfile)
	// Возвращаем указатель на новый Logger.
	return &Logger{
		level:   level,
		logger:  logger,
		logFile: file,
	}, nil
}

// SetLevel позволяет изменить текущий уровень логировния.
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Debug записывает сообщения уровня Debug.
func (l *Logger) DEBUG(v ...interface{}) {
	if l.level <= INFO {
		l.logger.SetPrefix("DEBUG: ")
		l.logger.Println(v...)
	}

}

// INFO записывает сообщения уровня Debug.
func (l *Logger) Info(v ...interface{}) {
	if l.level <= INFO {
		l.logger.SetPrefix("INFO: ")
		l.logger.Println(v...)
	}
}

// WARN записывает сообщения уровня Debug.
func (l *Logger) Warn(v ...interface{}) {
	if l.level <= WARN {
		l.logger.SetPrefix("WARN: ")
		l.logger.Println(v...)
	}
}

// WARN записывает сообщения уровня Debug.
func (l *Logger) Error(v ...interface{}) {
	if l.level <= ERROR {
		l.logger.SetPrefix("Error: ")
		l.logger.Println(v...)
	}
}

func (l *Logger) Close() {
	l.logFile.Close()
}
