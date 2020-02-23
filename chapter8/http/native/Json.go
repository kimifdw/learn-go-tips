package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Address: 地址对象
type Address struct {
    Street   string `json:"street,omitempty"`
    Landmark string `json:"landmark,omitempty"`
    Pincode  int    `json:"pincode,omitempty"`
}

// sendJSON: 发送json
func sendJSON(w http.ResponseWriter, r *http.Request) {
    a := Address{
        Street:   "Viman Nagar",
        Landmark: "Nexa",
        Pincode:  411014,
    }
    w.Header().Set("Content-Type", "application/json")
    encoder := json.NewEncoder(w)
    encoder.Encode(a)
}

// acceptJSON:  接收json
func acceptJSON(w http.ResponseWriter, r *http.Request) {
    a := Address{}
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&a)
    if err != nil {
        fmt.Fprintf(w, "Error parsing json:%v", err)
        return
    }
    fmt.Fprintln(w, "received:", &a)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", sendJSON)
    http.ListenAndServe(":3000", mux)
}
