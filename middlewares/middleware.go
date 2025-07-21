package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		start := time.Now()
		latency := time.Since(start)
		fmt.Printf("%s\n", r.Method)
		fmt.Printf("%s\n", r.RequestURI)
		fmt.Printf("%v\n", latency)
	})
}
