package main

import (
    "fmt"
    "net/http"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", sendResponse)
    server := http.Server{
        Addr:    ":3000",
        Handler: customMiddleware(mux),
    }
    server.ListenAndServe()
}

func sendResponse(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("default page"))
}

func customMiddleware(hand http.Handler) http.Handler {
    handlerFunc := func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("received new request...")
        w.Header().Set("X-Correlation-Id", "Guid")
        hand.ServeHTTP(w, r)
    }
    return http.HandlerFunc(handlerFunc)
}
