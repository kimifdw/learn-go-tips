package singleton

import (
	"log"
	"sync"
	"testing"
)

var w sync.WaitGroup

func TestSingletonSample(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()

	if s1 == s2 {
		log.Printf("s1的地址是 %p, s2的地址是 %p", s1, s2)
	}
}

func TestConcurrentSingleton(t *testing.T) {
	singletonList := make([]*singleton, 0)
	for i := 0; i < 100; i++ {
		w.Add(1)
		go func(group sync.WaitGroup) {
			defer w.Done()
			singletonList = append(singletonList, GetInstance())
		}(w)
	}
	w.Wait()

	for i := 0; i < len(singletonList)-1; i++ {
		if singletonList[i] != singletonList[i+1] {
			t.Fatal("instance is not equal")
		}
	}
}
