package main

import (
	"os"
	"fmt"
	"path/filepath"


)

//
func generateFileName(path string) {
	// 初始化slice
	nameList := []string{}
	// 遍历目录
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		// 读取文件名
		fileName := f.Name()
		if len(fileName) == 11 {
			// 文件名截取
			rs := []rune(fileName)
			nameList = append(nameList, string(rs[0:7]))
		}
		return nil;
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	// 文件创建
	file, err := os.Create("fileName.txt");
	if err != nil {
		fmt.Println("fileName.txt文件")
		file, err = os.Create("fileName.txt")
		if err != nil {
			fmt.Println("文件创建失败！")
			return
		}
	}
	fmt.Println(len(nameList))
	for _,v := range nameList {
		file.WriteString(v+"\r\n")
	}
	defer file.Close()

}

func main() {
	generateFileName("/Users/kimiyu/Desktop/images")
}
