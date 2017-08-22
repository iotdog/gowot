package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/leesper/holmes"
)

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{w, http.StatusOK}
}

func LoggerMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	rw := NewResponseWriter(w)
	next(rw, r)

	logging := fmt.Sprintf("%s -- %v %s %s %s %s - %s %v",
		r.RemoteAddr,
		start,
		r.Method,
		r.URL.Path,
		r.Proto,
		http.StatusText(rw.StatusCode),
		r.Header.Get("User-Agent"),
		time.Since(start))

	holmes.Infoln(logging)
}
