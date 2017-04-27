package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type logger struct {
	handler http.Handler
	output  log.Logger
}

//NewLogger ...
func NewLogger() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		lgr := &logger{
			handler: h,
		}

		return lgr
	}
}

func (lg *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ww := wrapWriter(w)

	start := time.Now()
	lg.handler.ServeHTTP(ww, r)
	end := time.Now()
	duration := end.Sub(start)

	log.Println(
		fmt.Sprintf("%s[%d]", r.Method, ww.Status()),
		duration,
		"-", r.URL.String(),
	)
}

type writerProxy interface {
	http.ResponseWriter
	Status() int
	BytesWritten() int
	Tee(io.Writer)
	Unwrap() http.ResponseWriter
}

func wrapWriter(w http.ResponseWriter) writerProxy {
	return &basicWriter{ResponseWriter: w}
}

type basicWriter struct {
	http.ResponseWriter
	wroteHeader bool
	code        int
	bytes       int
	tee         io.Writer
}

func (b *basicWriter) WriteHeader(code int) {
	if !b.wroteHeader {
		b.code = code
		b.wroteHeader = true
		b.ResponseWriter.WriteHeader(code)
	}
}
func (b *basicWriter) Write(buf []byte) (int, error) {
	b.WriteHeader(http.StatusOK)
	n, err := b.ResponseWriter.Write(buf)
	if b.tee != nil {
		_, err2 := b.tee.Write(buf[:n])
		if err == nil {
			err = err2
		}
	}
	b.bytes += n
	return n, err
}
func (b *basicWriter) maybeWriteHeader() {
	if !b.wroteHeader {
		b.WriteHeader(http.StatusOK)
	}
}
func (b *basicWriter) Status() int {
	return b.code
}
func (b *basicWriter) BytesWritten() int {
	return b.bytes
}
func (b *basicWriter) Tee(w io.Writer) {
	b.tee = w
}
func (b *basicWriter) Unwrap() http.ResponseWriter {
	return b.ResponseWriter
}
