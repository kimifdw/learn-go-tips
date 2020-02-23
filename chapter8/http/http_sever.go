package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type MyHandler struct {
}

// ServerHTTP 自定义路由映射
func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloWorld(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

type HelloWorldMuxHandler struct{}

func sayHelloWorldMux(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", params["name"])
}

func (muxHandler *HelloWorldMuxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "你好, %s!", params["name"])
}

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("URL:", r.URL.Path)
	fmt.Println("Scheme:", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println(k, ":", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world!")
}

// checkToken 校验token
func checkToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("token")
		if token == "www.baidu.com" {
			log.Printf("Token check access: %s\n", r.RequestURI)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func main() {
	// 注册/路由，并由sayHelloWorld函数来处理
	// http.HandleFunc("/", sayHelloWorld)

	http.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request) {
		params := req.URL.Query()
		fmt.Fprintf(resp, "你好，%s", params.Get("name"))
	})
	// nil默认使用DefaultServeMux来实现
	// handler := MyHandler{}
	// err := http.ListenAndServe(":8080", &handler)
	// if err != nil {
	// 	log.Fatalf("start http failed:%v", err)
	// }

	r := mux.NewRouter()
	// r.HandleFunc("/hello/{name}", sayHelloWorldMux)
	r.Handle("/hello/{name}", &HelloWorldMuxHandler{})
	// 限制方法、路由正则匹配
	r.HandleFunc("/hello/{name:[a-z]+}", sayHelloWorldMux).Methods("GET", "POST")

	// 路由匹配
	r.PathPrefix("/hello").HandlerFunc(sayHelloWorldMux)

	// 域名限制
	r.HandleFunc("/hello", sayHelloWorldMux).Host("goweb.test")

	// scheme限制
	r.HandleFunc("/hello", sayHelloWorldMux).Schemes("https")

	// 请求头限制
	r.HandleFunc("/request/header", func(w http.ResponseWriter, r *http.Request) {
		header := "X-Requested-With"
		fmt.Fprintf(w, "包含指定请求头【%s=%s】", header, r.Header[header])
	}).Headers("X-Requested-With", "XMLHttpRequest")

	// 请求参数限制，无效返回404
	r.HandleFunc("/request/header", func(w http.ResponseWriter, r *http.Request) {
		query := "token"
		fmt.Fprintf(w, "包含指定查询字符串【%s=%s】", query, r.FormValue(query))
	}).Queries("token", "test")

	// 自定义匹配规则
	r.HandleFunc("/custom/matcher", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "请求来自指定域名: %s", r.Referer())
	}).MatcherFunc(func(req *http.Request, match *mux.RouteMatch) boot {
		return req.Referer() == "https://www.baidu.com"
	})

	// 路由分组及路由命名
	postRouter := r.PathPrefix("/posts").Subrouter()
	postRouter.HandleFunc("/", listPosts).Methods("GET").Name("posts.index")
	postRouter.HandleFunc("/create", createPost).Methods("POST")

	// 插入中间件
	postRouter.Use(checkToken)

	// 处理静态资源
	var dir string
	// 运行传入参数：go run http_server.go -dir=static
	flag.StringVar(&dir, "dir", ".", "静态资源所在目录，默认为当前目录")
	flag.Parse()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	log.Fatal(http.ListenAndServe(":8080", r))
}
