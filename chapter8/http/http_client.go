package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8080/hello?name=嘟嘟", nil)
	if err != nil {
		log.Println("req failed: ", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("resp failed: ", err)
	}

	body := resp.Body

	defer body.Close()
	io.Copy(os.Stdout, body)
}
