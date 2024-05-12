package logger

type Logger interface {
	Error(...interface{})
	Warn(...interface{})
	Info(...interface{})
}
