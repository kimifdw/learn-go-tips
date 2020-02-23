package main

import (
	"container/list"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

import myRpc "learn-go-tips/chapter8/rpcutil"

type Person struct {
	Name string
	Age  int
}

type ByName []Person

// 返回要排序的数组长度
func (ps ByName) Len() int {
	return len(ps)
}

// 根据条件比较元素间的大小
func (ps ByName) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

// 根据Less的结果交换数组元素的位置
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// http实现
func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>Hello,World</title>
			</head>
			<body>
				Hello,World!
			</body>
		</html>
		`,
	)
}

func main() {
	// strings包。具体使用可以查strings的api
	// Contains: 包含
	fmt.Println(strings.Contains("test", "es"))
	// Count: t在test中的个数
	fmt.Println(strings.Count("test", "t"))
	// HasPrefix：test的前缀是否为te
	fmt.Println(strings.HasPrefix("test", "te"))
	// HasSuffix：test是否以st结尾
	fmt.Println(strings.HasSuffix("test", "st"))

	// io包。重点关注Reader和Writer
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("不存在test.txt文件")
		file, err = os.Create("test.txt")
		if err != nil {
			fmt.Println("文件创建失败！")
			return
		}
		file.WriteString("测试")
	}
	// 始终会执行Close()关闭方法
	defer file.Close()

	stat, err := file.Stat() // 返回文件信息
	if err != nil {
		fmt.Println("读取文件信息异常：", err)
		return
	}

	bs := make([]byte, stat.Size()) // 创建字节数组，用于存放读取的内容
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	// 使用ioutil包操作文件
	bsNew, err := ioutil.ReadFile("test.txt") // 读取文件，读取完后会自动将file关闭掉
	if err != nil {
		// 错误处理
		return
	}
	str = string(bsNew)
	fmt.Println(str)

	// 操作目录
	dir, err := os.Open(".") // 当前目录
	if err != nil {
		fmt.Println("读取目录失败!", err)
		return
	}

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}

	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}

	// 递归的读取目录
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path + "-" + string(info.Mode()))
		return nil
	})

	// 定义错误
	err = errors.New("错误信息")
	fmt.Println(err)

	// go包含三个容器，分别是heap【堆】，List【双向链表】，Ring【环】
	var x list.List
	// 从队列尾部加入
	e1 := x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	// 从队列头部加入
	x.PushFront(1)
	x.PushFront(2)
	x.PushFront(3)
	// 在index=2的元素后面插入e1
	x.InsertAfter(2, e1)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// SORT排序
	kids := []Person{
		{"Jill", 9}, {"Jack", 10},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)

	// 调用TCP
	// go ts.Server()
	// go ts.Client()
	// 调用RPC
	go myRpc.ServerRpc()
	go myRpc.ClientRPC()

	var input string
	fmt.Scanln(&input)

	// 注册调用url时触发的函数
	http.HandleFunc("/hello", hello)
	// 监控9000端口并调用处理函数来处理请求，默认处理方式为nil，
	http.ListenAndServe(":9000", nil)

}
