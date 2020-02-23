package main

// Item 项
type Item struct {
    ID int  `json:"id"`
    Name string `json:"name"`
}
// Response 响应结果
type Response struct {
    Ok bool     `json:"ok"`
    ID int  `json:"id"`
    Message string `json:"msg"`
}
