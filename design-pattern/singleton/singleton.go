package singleton

import "sync"

type singleton struct {
	name string
}

var mySingleton *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		mySingleton = &singleton{}
	})
	return mySingleton
}
