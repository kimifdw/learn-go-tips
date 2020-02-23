package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, Gophers!"))
}

func main() {
    registerRoutes()
    server := http.Server{
        Addr:    ":3000",
        Handler: mux,
    }
    server.ListenAndServe()
}
