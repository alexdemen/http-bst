package log

import (
	"fmt"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"io"
	"net/http"
	"time"
)

type Logger struct {
	instance zerolog.Logger
}

func (l *Logger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := l.instance.Info()
	entry.Str("url", r.URL.Path)
	return &StructuredLoggerEntry{entry}
}

type StructuredLoggerEntry struct {
	Event *zerolog.Event
}

func (s *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	s.Event.
		Int("resp_status", status).
		Int64("elapsed", elapsed.Milliseconds()).
		Send()

}

func (s StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	s.Event = s.Event.
		Str("stack", string(stack)).
		Str("panic", fmt.Sprintf("%+v", v))
}

func NewLogger(writer io.Writer) *Logger {
	return &Logger{
		instance: zerolog.New(writer).With().Timestamp().Logger(),
	}
}

func (l *Logger) Info(data map[string]interface{}) {
	if l != nil {
		collectAndSend(l.instance.Info(), data)
	}
}

func (l *Logger) Error(data map[string]interface{}) {
	if l != nil {
		collectAndSend(l.instance.Error(), data)
	}
}

func collectAndSend(event *zerolog.Event, data map[string]interface{}) {
	for key, val := range data {
		switch v := val.(type) {
		case int:
			event.Int(key, v)
		case string:
			event.Str(key, v)
		}
	}

	event.Send()
}
