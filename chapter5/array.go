// 1. [n]T：拥有n个T类型的值的数组
// 2. 数组长度不可改变
package main

import "fmt"

type T struct {
    n int
}

func main() {
    var a [2]string
    a[0] = "Hello"
    a[1] = "World"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes)

    ts:=[2]T{}
    for i, t := range ts[:] {
        switch i {
        case 0:
            t.n = 3
            ts[1].n = 9
        case 1:
            fmt.Print(t.n," ")
        }
    }
    fmt.Print(ts)
}
