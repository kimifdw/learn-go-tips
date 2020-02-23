package main

import(
    "io/ioutil"
    "log"
    "net/http"
    "fmt"
    "html/template"
    "regexp"
)

type Page struct {
    title string
    body []byte
}

// 保存文件
func (p *Page) save() error  {
    filename := p.title + ".txt"
    return ioutil.WriteFile(filename, p.body, 0600)
}

// 加载page
func loadPage(title string) (*Page,error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{title: title, body: body}, nil
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// 获取title
func getTitle(w http.ResponseWriter, r *http.Request) (string,error)  {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", nil
    }
    return m[2], nil
}

// 定义handler
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// 查看
func viewHandler(w http.ResponseWriter, r *http.Request)  {
    title, error := getTitle(w,r)
    if error != nil {
        return
    }
    p, err := loadPage(title)
    if err!=nil {
        http.Redirect(w, r, "/edit/" + title, http.StatusFound)
        return
    }
    renderTemplate(w, "view", p)
}

// 编辑
func editHandler(w http.ResponseWriter, r *http.Request)  {
    title, _ := getTitle(w, r)

    p, err := loadPage(title)
    if err != nil {
        p = &Page{title: title}
    }
    renderTemplate(w, "edit",p)
}

// 保存
func saveHandler(w http.ResponseWriter, r *http.Request)  {
    title, _:= getTitle(w, r)
    body := r.FormValue("body")
    p := &Page{title: title, body: []byte(body)}
    err := p.save()
    if err!=nil  {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/" +title, http.StatusFound)
}

// 页面模板
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page)  {
    t, err := template.ParseFiles(tmpl+".html")
    if err !=nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main()  {
    // p1 := &Page{title: "testPage", body: []byte("this is a sample page")}
    // p1.save()

    // p2, err := loadPage("testPage")
    // if err != nil {
    //     fmt.Println(err)
    // }

    // fmt.Println(string(p2.body))
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    http.HandleFunc("/save/", saveHandler)
    http.HandleFunc("/", handler)

    log.Fatal(http.ListenAndServe(":8081",nil))
}
