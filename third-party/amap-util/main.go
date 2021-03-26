package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"io/ioutil"
	"net/http"
)

type GeoResp struct {
	Status  string     `json:"status"`
	Info    string     `json:"info"`
	GeoCode []GeoCodes `json:"geocodes"`
}

type GeoCodes struct {
	Location string `json:"location"`
}

type DirectionResp struct {
	Data DirectionData `json:"data"`
}

type DirectionData struct {
	Paths []PathData `json:"paths"`
}

type PathData struct {
	Distance int32 `json:"distance"`
}

func main() {
	readFile("1.xlsx")
}

func getGeo(address string) int32 {
	key := "2a5eacf6bd972539822c72a480285b7d"
	addressUrl := "https://restapi.amap.com/v3/geocode/geo?key=" + key + "&address=" + address
	resp, err := http.Get(addressUrl)
	if err != nil {
		fmt.Print(err)
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("resp body error:%v\n", err)
	}

	//fmt.Println(string(result))

	var geoResp = GeoResp{}
	err = json.Unmarshal(result, &geoResp)
	if err != nil {
		fmt.Printf("geoResp error:%v\n", err)
	}
	if len(geoResp.GeoCode) == 0 {
		return 0
	}
	return getDistance(key, geoResp.GeoCode[0].Location)
}

func getDistance(key string, destination string) int32 {
	origin := "121.382423,31.245406"
	directionUrl := "https://restapi.amap.com/v4/direction/bicycling?key=" + key + "&origin=" + origin + "&destination=" + destination

	resp, err := http.Get(directionUrl)
	if err != nil {
		fmt.Print(err)

	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("resp body error:%v\n", err)

	}

	//fmt.Println(string(result))

	var directionResp = DirectionResp{}
	err = json.Unmarshal(result, &directionResp)
	if err != nil {
		fmt.Printf("direction resp body error:%v\n", err)

	}

	paths := directionResp.Data.Paths
	if len(paths) == 0 {
		return 0
	}
	return paths[0].Distance
}

func readFile(fileName string) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	area := []string{"普陀区", "虹口区", "嘉定区", "长宁区"}
	for _, area := range area {
		// Get all the rows in the Sheet1.
		rows, _ := f.GetRows(area)

		for index, _ := range rows {
			address, _ := f.GetCellValue(area, fmt.Sprintf("F%d", index))
			if address == "地址1" || address == "" {
				continue
			}

			fmt.Printf("地址:%s,index:%d\n", address, index)

			if index+1 == len(rows) {
				address, _ := f.GetCellValue(area, fmt.Sprintf("F%d", index+1))
				fmt.Printf("地址:%s,index:%d\n", address, index+1)
			}
			distince := getGeo(address)
			fmt.Println(address + ":" + string(distince))
			_ = f.SetCellValue(area, fmt.Sprintf("I%d", index), distince)
		}

	}

	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}

}
