package main

import (
    "fmt"
    "net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
    hed := r.Header
    // 获取head的值
    hed.Get("Accept")
    fmt.Fprintln(w, hed.Get("Accept"))
    w.Header().Set("ALLOWED", "GET,POST")
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("set allowed headers\n"))
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", headers)
    http.ListenAndServe(":3000", mux)
}
