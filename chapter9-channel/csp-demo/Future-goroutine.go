package main

import "fmt"

type Function func(string) (string, error)

type Future interface {
	SuccessCallback() error
	FailCallback() error
	Execute(Function) (bool, chan struct{})
}

type AccountCache struct {
	Name string
}

func (a *AccountCache) SuccessCallback() error {
	fmt.Println("It is success~")
	return nil
}

func (a *AccountCache) FailCallback() error {
	fmt.Println("It is failed!")
	return nil
}

func (a *AccountCache) Execute(f Function) (bool, chan struct{}) {
	done := make(chan struct{})

	go func(a *AccountCache) {
		_, err := f(a.Name)
		if err != nil {
			_ = a.FailCallback()
		} else {
			_ = a.SuccessCallback()
		}
		done <- struct{}{}
	}(a)
	return true, done
}

func NewAccountCache(name string) *AccountCache {
	return &AccountCache{
		name,
	}
}

func testFuture() {
	var future Future
	future = NewAccountCache("Emily")
	updateFunc := func(name string) (string, error) {
		fmt.Println("cache updated:", name)
		return name, nil
	}

	_, done := future.Execute(updateFunc)
	defer func() {
		<-done
	}()
}

func main() {
	testFuture()
}
