// 1. 切片为数组提供动态大小
// 2. []T: 一个元素类型为T的切片，通过上界和下届限定以冒号分隔
// 3. 半开区间：包括第一个元素但排除最后一个元素
// 4. 切片不存储任何数据，更改切片元素会修改底层数组中的元素
// 5. len(q)：长度包含元素个数；cap(q)：容量是从第一个元素开始到其底层数组元素末尾的个数
// 6. 零值为nil,长度和容量为0且没有底层数组
// 7. 用内建函数make来创建
// 8. 切片可包含切片
// 9. append: 追加元素
package main

import (
    "fmt"
    "sort"
)

func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}

    var s []int = primes[1:4]
    fmt.Println(s)

    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names)
    a := names[0:2]
    b := names[1:3]

    fmt.Println(a, b)
    fmt.Println(names)

    q := []int{2, 3, 5, 7, 11, 13}
    fmt.Println(q)

    r := []bool{true, false, true, false, true}
    fmt.Println(r)

    // s := []struct {
    // 	i int
    // 	b bool
    // }{
    // 	{2, true},
    // 	{3, false},
    // 	{5, true},
    // 	{7, false},
    // 	{11, true},
    // 	{13, false},
    // }
    // fmt.Println(s)

    // 截取其长度为0
    q = q[:0]
    printSlice(q)

    // 拓展其长度
    q = q[:4]
    printSlice(q)

    q = q[2:]
    printSlice(q)

    e := make([]int, 5)
    printSlice(e)
    f := make([]int, 0, 5)
    printSlice(f)
    c := e[:2]
    printSlice(c)
    d := f[2:5]
    printSlice(d)

    var t []int
    t = append(t, 0)
    printSlice(t)

    t = append(t, 1)
    printSlice(t)

    // 可以一次性添加多个元素
    t = append(t, 2, 3, 4)
    printSlice(t)

    // 从小到大排列
    sl := IntSlice([]int{89, 14, 8, 9, 56, 95, 3})
    fmt.Println(sl)
    //sort.Sort(sl)
    sort.Ints(sl)
    fmt.Println(sl)

    langs := []Lang{
        {"rust", 2},
        {"go", 1},
        {"swift", 3},
    }
    // 传入比较值，支持自定义排序
    sort.Slice(langs, func(i, j int) bool {
        return langs[i].Rank < langs[j].Rank
    })
    fmt.Printf("%v\n", langs)
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// 排序
type IntSlice []int

func (p IntSlice) Len() int {
    return len(p)
}
func (p IntSlice) Less(i, j int) bool {
    return p[i] < p[j]
}
func (p IntSlice) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

type Lang struct {
    Name string
    Rank int
}
