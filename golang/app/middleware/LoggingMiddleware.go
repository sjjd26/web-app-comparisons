package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", request.Method, request.URL.Path)

		next.ServeHTTP(writer, request)

		fmt.Printf("%v", writer.Header().Get("Status"))

		log.Printf("Completed %s %s in %v, status code %v", request.Method, request.URL.Path, time.Since(start), writer.Header().Get("Status"))
	})
}
