package main

import "fmt"

func generator(max int) <-chan int {
    out := make(chan int, 100)
    go func() {
        for i := 1; i <= max; i++ {
            out <- i
            fmt.Printf("input value:%d\n", i)
        }
        close(out)
    }()
    return out
}

func power(in <-chan int) <-chan int {
    out := make(chan int, 100)
    go func() {
        for v := range in {
            midValue := v * v
            out <- midValue
            fmt.Printf("v*v=%d\n", midValue)
        }
        close(out)
    }()
    return out
}

func sum(in <-chan int) <-chan int {
    out := make(chan int, 100)
    go func() {
        var sum int
        for v := range in {
            sum += v
        }
        out <- sum
        close(out)
    }()
    return out
}

func main() {
    fmt.Println(<-sum(power(generator(3))))
}
