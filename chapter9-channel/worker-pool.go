package main

import (
	"fmt"
	"sync"
)

// Worker：工人
type Worker struct {
	ID       int
	Name     string
	StopChan chan bool
}

// Job：任务
type Job interface {
	Start(worker *Worker) error
}

func (w *Worker) Start(jobQueue chan Job) {
	w.StopChan = make(chan bool)
	successChan := make(chan bool)

	go func() {
		successChan <- true
		for true {
			job := <-jobQueue
			if job != nil {
				job.Start(w)
			} else {
				fmt.Printf("worker %s to be stopped\n", w.Name)
				w.StopChan <- true
				break
			}
		}
	}()

	<-successChan
}

func (w *Worker) Stop() {
	_ = <-w.StopChan
	fmt.Printf("worker %s stopped\n", w.Name)
}

// Pool：线程池
type Pool struct {
	Name      string
	Size      int
	Workers   []*Worker
	QueueSize int
	Queue     chan Job
}

// Initialize：初始化线程池
func (p *Pool) Initialize() {
	if p.Size < 1 {
		p.Size = 1
	}

	p.Workers = []*Worker{}
	for i := 1; i <= p.Size; i++ {
		worker := &Worker{
			ID:   i - 1,
			Name: fmt.Sprintf("%s-worker-%d", p.Name, i-1),
		}
		p.Workers = append(p.Workers, worker)
	}

	if p.QueueSize < 1 {
		p.QueueSize = 1
	}
	p.Queue = make(chan Job, p.QueueSize)
}

func (p *Pool) Start() {
	for _, worker := range p.Workers {
		worker.Start(p.Queue)
	}
	fmt.Println("all workers started")
}

func (p *Pool) Stop() {
	close(p.Queue)

	var wg sync.WaitGroup
	for _, worker := range p.Workers {
		wg.Add(1)
		go func(w *Worker) {
			defer wg.Done()
			w.Stop()
		}(worker)
	}
	wg.Wait()
	fmt.Println("all workers stopped")
}

type PrintJob struct {
	Index int
}

func (pj *PrintJob) Start(worker *Worker) error {
	fmt.Printf("job %s-%d\n", worker.Name, pj.Index)
	return nil
}

func main() {
	pool := &Pool{
		Name:      "test",
		Size:      5,
		QueueSize: 20,
	}
	pool.Initialize()
	pool.Start()
	defer pool.Stop()

	for i := 1; i <= 100; i++ {
		job := &PrintJob{
			Index: i,
		}
		pool.Queue <- job
	}
}
