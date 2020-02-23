package main

import (
    "fmt"
    "net/http"
)

func operateCookie(w http.ResponseWriter, r *http.Request) {
    cookie := http.Cookie{
        Name:  "cookie-1",
        Value: "hello world",
    }
    http.SetCookie(w, &cookie)

    cookies := r.Cookies()
    for _, cookie := range cookies {
        fmt.Fprintln(w, cookie)
    }
    // get named cookie
    cookieObj, err := r.Cookie("cookie-1")
    if err != nil {
        fmt.Fprintln(w, err.Error())
    }
    fmt.Fprintln(w, cookieObj)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", operateCookie)
    http.ListenAndServe(":3000", mux)
}
