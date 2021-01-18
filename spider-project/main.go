package main

import (
    "fmt"
    "github.com/axgle/mahonia"
    "github.com/gocolly/colly/v2"
    "strings"
    "time"
)

func main() {
    c := colly.NewCollector(
        colly.AllowedDomains("www.zgsydw.com"),
    )

    //On every a element which has href attribute call callback
    //c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    //   link := e.Attr("href")
    //   // Print link
    //   fmt.Printf("Link found: %q -> %s\n", e.Text, link)
    //   // Visit link found on page
    //   // Only those links are visited which are in AllowedDomains
    //   c.Visit(e.Request.AbsoluteURL(link))
    //})

    // 解决中文显示问题
    decoder := mahonia.NewDecoder("gbk")

    c.OnHTML("div[class='ggxx_nr'] ul", func(e *colly.HTMLElement) {

        e.ForEach("li", func(i int, item *colly.HTMLElement) {
            link := item.ChildAttr("a[href]", "href")
            title := item.ChildAttr("a[href]", "title")
            dateStr := item.ChildText("span")
            date, err := time.Parse("2006-01-02", dateStr)
            if err != nil {
                panic("日期转换失败")
            }

            day := date.Day()
            result := decoder.ConvertString(title)
            // 查看当天日期且标题存在招字的访问地址
            if day == time.Now().Day() && strings.ContainsAny(result, "宁波") {
                fmt.Printf("[%s]Link found: %s -> %s\n", dateStr, result, link)
                c.Visit(e.Request.AbsoluteURL(link))
            } else {
                remainDay := time.Now().Day() - day
                if remainDay > 0 {
                    fmt.Printf("[%d天前]Link found: %s -> %s\n", remainDay, result, link)
                }
            }
        })
    })

    // Before making a request print "Visiting ..."
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
    })

    // Start scraping on https://hackerspaces.org
    c.Visit("http://www.zgsydw.com/zhejiang/zhaopin/")
}
