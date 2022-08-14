package mocking

import (
	"fmt"
	"io"
	"time"
)

func CountDown(writer io.Writer, sleeper *SpySleeper) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprint(writer, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(writer, "Go!")

}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type ConfigurableSleeper struct {
	duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}
