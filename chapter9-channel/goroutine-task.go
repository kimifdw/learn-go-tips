package main

import (
	"fmt"
	"sync"
)

type Task struct {
	begin  int
	end    int
	result chan<- int
}

func initTask(taskchan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskchan <- tsk
	}
	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskchan <- tsk
	}
	close(taskchan)
}

func distributeTask(taskchan <-chan task, wait *sync.WaitGroup, result chan int) {
	for v := range taskchan {
		wait.Add(1)
		go processTask(v, wait)
	}
	wait.Wait()
	close(result)
}

func processTask(t task, wait *sync.WaitGroup) {
	t.do()
	wait.Done()
}

func processResult(resultchan chan int) int {
	sum := 0
	for r := range resultchan {
		sum += r
	}
	return sum
}

func (t *task) Do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	t.result <- sum
}

func main() {
	taskchan := make(chan task, 10)

	resultchan := make(chan int, 10)

	// 用于同步等待任务的执行
	wait := &sync.WaitGroup{}

	go initTask(taskchan, resultchan, 100)

	go distributeTask(taskchan, wait, resultchan)

	sum := processResult(resultchan)

	fmt.Println("sum=", sum)
}
