package main

import (
  "fmt"
  "net/http"
  "time"
)

//Logger is a middleware handler that does request logging
type Logger struct {
    handler http.Handler
}

//ServeHTTP handles the request by passing it to the real
//handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    l.handler.ServeHTTP(w, r)
    fmt.Println("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}