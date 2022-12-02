package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

const DefaultStackBodyString = "500 Internal Server Error"

type PanicInformation struct {
	RecoveredPanic interface{}
	Stack          []byte
	Request        *http.Request
}

func (p *PanicInformation) StackAsString() string {
	return string(p.Stack)
}

func (p *PanicInformation) RequestDescription() string {
	if p.Request == nil {
		return "Request is nil"
	}

	var queryOutput string
	if p.Request.URL.RawQuery != "" {
		queryOutput = "?" + p.Request.URL.RawQuery
	}

	return fmt.Sprintf("%s %s%s", p.Request.Method, p.Request.URL.Path, queryOutput)
}

type Recovery struct {
	Logger           ILogger
	PanicHandlerFunc func(*PanicInformation)
	PrintStack       bool
	StackAll         bool
	StackSize        int
}

func NewRecovery() *Recovery {
	return &Recovery{
		Logger:     log.New(os.Stdout, "[ak1m1tsu] ", 0),
		PrintStack: true,
		StackAll:   false,
		StackSize:  1024 * 8,
	}
}

func (rec *Recovery) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			infos := &PanicInformation{
				RecoveredPanic: err,
				Request:        r,
				Stack:          make([]byte, rec.StackSize),
			}
			infos.Stack = infos.Stack[:runtime.Stack(infos.Stack, rec.StackAll)]

			if w.Header().Get("Content-Type") == "" {
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			}
			fmt.Fprint(w, DefaultStackBodyString)
			rec.Logger.Printf("PANIC: %s\n%s", err, infos.Stack)
		}
	}()

	next(w, r)
}
