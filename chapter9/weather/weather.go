package weather

import (
    "fmt"
    "golang.org/x/sync/singleflight"
)

type Info struct {
    TempC, TempF int
    Conditions   string
}

var group singleflight.Group

func City(city string) (*Info, error) {
    results, err, _ := group.Do(city, func() (interface{}, error) {
        info, err := fetchWeatherFromDB(city)
        return info, err
    })
    if err != nil {
        return nil, fmt.Errorf("weather.City %s: %w", city, err)
    }
    return results.(*Info), nil
}
