package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", request.Method, request.URL.Path)

		next.ServeHTTP(writer, request)

		log.Printf("Completed %s %s in %v", request.Method, request.URL.Path, time.Since(start))
	})
}
