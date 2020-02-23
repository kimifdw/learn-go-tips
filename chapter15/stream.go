package main

import(
	"os"
    "log"
    "encoding/json"
)

func main()  {
    dec := json.NewDecoder(os.Stdin)
    enc := json.NewEncoder(os.Stdout)

    for {
        var v map[string]interface{}
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }
        if err := enc.Encode(&v); err !=nil {
            log.Println(err)
        }
    }
}
