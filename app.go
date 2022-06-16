package main

import (
  "fmt"
  "os"
  "net/http"
  "time"
  "log"
)

func main() {
  // byteCounterExample()
  listenAndServe()
}

func byteCounterExample() {
  var c ByteCounter
  c.Write([]byte("hello"))

  c = 0 // reset the counter
  var name = "Dolly"
  fmt.Fprintf(&c, "hello, %s", name)
  fmt.Println(c) // "12", = len("hello, Dolly")
}

//NewLogger constructs a new Logger middleware handler
// func NewLogger(handlerToWrap http.Handler) *Logger {
//   return &Logger{handlerToWrap}
// }

func HelloHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, World!"))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
  curTime := time.Now().Format(time.Kitchen)
  w.Write([]byte(fmt.Sprintf("the current time is %v", curTime)))
}

func listenAndServe() {
  addr := os.Getenv("ADDR")

  mux := http.NewServeMux()
  mux.HandleFunc("/v1/hello", HelloHandler)
  mux.HandleFunc("/v1/time", CurrentTimeHandler)

  //wrap entire mux with logger middleware
  // wrappedMux := NewLogger(mux) // Sostituito da &Logger{mux} in ListenAndServe
  
  log.Printf("server is listening at %s", addr)
  //use wrappedMux instead of mux as root handler
  log.Fatal(http.ListenAndServe(addr, &Logger{mux}  ))
}
