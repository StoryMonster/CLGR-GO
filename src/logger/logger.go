package logger

import (
	"log"
	"io"
)

const (
	Debug = 1
	Info = 2
	Warn = 3
	Error = 4
)

type Logger struct {
	logger *log.Logger
	level int
}

func New(logWriter io.Writer, level int) *Logger {
	logger := log.New(logWriter, "[CLGR]", log.Ldate | log.Ltime)
	return &Logger{logger, level}
}

func (logger *Logger) DEBUG(str string) {
	if logger.level > Debug { return }
    logger.write("[DEBUG] " + str)
}

func (logger *Logger) INFO(str string) {
	if logger.level > Info { return }
    logger.write("[INFO] " + str)
}

func (logger *Logger) WARN(str string) {
	if logger.level > Warn { return }
    logger.write("[WARN] " + str)
}

func (logger *Logger) ERROR(str string) {
	if logger.level > int(Error) { return }
    logger.write("[ERROR] " + str)
}

func (logger *Logger) write(str string) {
	logger.logger.Println(str)
}