package main

import (
    "context"
    "encoding/json"
    "flag"
    "fmt"
    "github.com/go-redis/redis"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    _ "net/http/pprof"
    "os"
    "os/signal"
    "sort"
    "strconv"
    "strings"
    "time"
)

const TIME_LAYOUT string = "2006-01-02 15:04:05"

type IatMessage struct {
    Did        string `json:"did"`
    Content    string `json:"content"`
    CreateTime string `json:"createTime"`
}

func init() {
    initRedis()
}

var client *redis.Client

func initRedis() {
    client = redis.NewClient(&redis.Options{
        Addr:     "172.16.124.78:6379",
        Password: "AuxVoice123", // no password set
        DB:       0,             // use default DB
    })
    pong, err := client.Ping().Result()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(pong, err)
}

func main() {
    var wait time.Duration
    flag.DurationVar(&wait, "connnect timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish")
    flag.Parse()

    r := mux.NewRouter()
    r.HandleFunc("/logs/{did}", didHandler).Methods("GET")

    srv := &http.Server{
        Addr:         "0.0.0.0:9999",
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler:      r,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil {
            log.Println(err)
        }
    }()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    <-c

    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer client.Close()
    defer cancel()
    srv.Shutdown(ctx)

    log.Println("关闭http")
    os.Exit(0)
}

// did请求
func didHandler(resp http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    didReq := params["did"]
    var did string
    var reqStartTimeStr string
    if strings.Contains(didReq, "-") {
        didParam := strings.Split(didReq, "-")
        did = didParam[0]
        reqStartTimeStr = didParam[1]
    } else {
        did = didReq
    }

    iatMessages := []IatMessage{}
    value, err := client.HGetAll("aiui:iat:" + did).Result()
    if err != nil {

    }
    reqStartTime, err := strconv.ParseInt(reqStartTimeStr, 10, 64)
    if err != nil {
        log.Printf("string to int64 error:%v\n", err)
    }
    loc, _ := time.LoadLocation("Local")
    for s, s2 := range value {
        if s != "" && s2 != "" {
            date_time := time.Unix(reqStartTime, 0)
            createTime, err := time.ParseInLocation(TIME_LAYOUT, s, loc)
            if err != nil {
                log.Printf("日期转换失败：%v\n", err)
            }
            iatMessage := IatMessage{
                did,
                s2,
                s,
            }
            if createTime.After(date_time) || createTime.Equal(date_time) {
                iatMessages = append(iatMessages, iatMessage)
            }
        }
    }

    // 按日期倒叙排列
    sort.Slice(iatMessages, func(i, j int) bool {
        // 日期格式转换
        iTime, iErr := time.ParseInLocation(TIME_LAYOUT, iatMessages[i].CreateTime, loc)
        jTime, jErr := time.ParseInLocation(TIME_LAYOUT, iatMessages[j].CreateTime, loc)
        if iErr != nil || jErr != nil {
            log.Printf("日期转换失败：%v\n", iErr)
        }

        return iTime.After(jTime) || iTime.Equal(jTime)
    })
    resp.Header().Add("Content-Type", "application/json;charset=utf-8")
    resp.Header().Add("Transfer-Encoding", "gzip")
    json.NewEncoder(resp).Encode(iatMessages)
}
