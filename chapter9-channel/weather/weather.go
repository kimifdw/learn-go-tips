package weather

//
//import (
//	"context"
//	"fmt"
//	"golang.org/x/sync/errgroup"
//	"golang.org/x/sync/semaphore"
//	"golang.org/x/sync/singleflight"
//	"sync"
//)
//
//type Info struct {
//	TempC, TempF int
//	Conditions   string
//}
//
//var group singleflight.Group
//
//func City(city string) (*Info, error) {
//	results, err, _ := group.Do(city, func() (interface{}, error) {
//		info, err := fetchWeatherFromDB(city)
//		return info, err
//	})
//	if err != nil {
//		return nil, fmt.Errorf("weather.City %s: %w", city, err)
//	}
//	return results.(*Info), nil
//}
//
//// CitiesWithErrgroup 查询多个城市的天气：errgroup
//func CitiesWithErrgroup(cities ...string) ([]*Info, error) {
//	var g errgroup.Group
//	var mu sync.Mutex
//	res := make([]*Info, len(cities))
//
//	for i, city := range cities {
//		i, city := i, city
//		g.Go(func() error {
//			info, err := City(city)
//			mu.Lock()
//			defer mu.Unlock()
//			// 写入保证线程安全
//			res[i] = info
//
//			return err
//		})
//	}
//	if err := g.Wait(); err != nil {
//		return nil, err
//	}
//	return res, nil
//}
//
//// CitiesWithSemaphores 用信号量方式查询多个城市天气
//func CitiesWithSemaphores(cities ...string) ([]*Info, error) {
//	var g errgroup.Group
//	var mu sync.Mutex
//	res := make([]*Info, len(cities))
//	// 定义10个任务
//	sem := make(chan struct{}, 10)
//
//	for i, city := range cities {
//		i, city := i, city
//		// 执行第一个任务
//		sem <- struct{}{}
//		g.Go(func() error {
//			info, err := City(city)
//			mu.Lock()
//			defer mu.Unlock()
//			res[i] = info
//			// 任务完成
//			<-sem
//			return err
//		})
//	}
//
//	if err := g.Wait(); err != nil {
//		return nil, err
//	}
//
//	return res, nil
//}
//
//func Cities(cities ...string) ([]*Info, error) {
//	ctx := context.TODO()
//	var g errgroup.Group
//	var mux sync.Mutex
//	res := make([]*Info, len(cities))
//	// 并发处理100个字符
//	sem := semaphore.NewWeighted(100)
//	for i, city := range cities {
//		i, city := i, city
//		cost := int64(len(city))
//		//
//		if err := sem.Acquire(ctx, cost); err != nil {
//			break
//		}
//		g.Go(func() error {
//			info, err := City(city)
//			mux.Lock()
//			defer mux.Unlock()
//			res[i] = info
//			// 释放
//			sem.Release(cost)
//			return err
//		})
//	}
//	if err := g.Wait(); err != nil {
//		return nil, err
//	} else if err := ctx.Err(); err != nil {
//		return nil, err
//	}
//
//	return res, nil
//}
