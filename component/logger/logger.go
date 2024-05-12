package logger

type Logger interface {
	Error(...interface{})
	Warn(...interface{})
	Println(...interface{})
	Fatal(...interface{})
	Debug(...interface{})
}
