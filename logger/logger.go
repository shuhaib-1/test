package logger

import (
	"log"
	"net/http"
	"os"
)

var (
	InfoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Info(v ...interface{}) {
	InfoLogger.Println(v...)
}

func Error(v ...interface{}) {
	ErrorLogger.Println(v...)
}

func HTTPLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		InfoLogger.Printf("=== Incoming Request ===")
		InfoLogger.Printf("Method: %s", r.Method)
		InfoLogger.Printf("Path: %s", r.RequestURI)
		InfoLogger.Printf("Remote Address: %s", r.RemoteAddr)
		InfoLogger.Printf("Headers: %v", r.Header)
		next.ServeHTTP(w, r)
	})
}
