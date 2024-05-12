package mylogger

import (
	"fmt"
	"log"
	"os"
)

const (
	errorColor = "\033[31m"
	infoColor  = "\033[32m"
	warnColor  = "\033[33m"
	fatalColor = "\033[35m"
	debugColor = "\033[37m"
	resetColor = "\033[0m"
)

type Logger struct {
	prefix string
	config *Config
}

func (l Logger) Fatal(v ...interface{}) {
	output := fmt.Sprintf("[FATAL] [%s]: %s", l.prefix, fmt.Sprintf("%s", v...))
	if l.config.isPersistent {
		l.WriteToFile(output)
	}
	log.Fatal(fatalColor, output, resetColor)
}

func (l Logger) Error(v ...interface{}) {
	output := fmt.Sprintf("[ERROR] [%s]: %s", l.prefix, fmt.Sprintf("%s", v...))
	log.Println(errorColor, output, resetColor)
	if l.config.isPersistent {
		l.WriteToFile(output)
	}
}

func (l Logger) Warn(v ...interface{}) {
	output := fmt.Sprintf("[WARN] [%s]: %s", l.prefix, fmt.Sprintf("%s", v...))
	log.Println(warnColor, output, resetColor)
	if l.config.isPersistent {
		l.WriteToFile(output)
	}
}

func (l Logger) Println(v ...interface{}) {
	if !l.config.debugMode {
		return
	}
	output := fmt.Sprintf("[INFO]  [%s]: %s", l.prefix, fmt.Sprintf("%s", v...))
	log.Println(infoColor, output, resetColor)
	if l.config.isPersistent {
		l.WriteToFile(output)
	}
}

func (l Logger) Debug(v ...interface{}) {
	if !l.config.debugMode {
		return
	}
	output := fmt.Sprintf("[DEBUG]  [%s]:[%s]", l.prefix, fmt.Sprintf("%s", v...))
	log.Println(debugColor, output, debugColor)
	if l.config.isPersistent {
		l.WriteToFile(output)
	}
}

func NewLogger(prefix string, config *Config) *Logger {
	if config == nil {
		config = DefaultConfig()
	}
	return &Logger{
		prefix: prefix,
		config: config,
	}
}

func (l Logger) WriteToFile(output string) {
	// Write to file
	file, err := os.OpenFile(l.config.persistentPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	_, err = fmt.Fprintf(file, output)
	if err != nil {
		log.Println(err)
	}
}
