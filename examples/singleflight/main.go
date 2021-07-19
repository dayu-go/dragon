package main

import (
	"errors"
	"log"
	"sync"

	"golang.org/x/sync/singleflight"
)

var errorNotExist = errors.New("not exist")
var g singleflight.Group

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			data, err := getData("key")
			// data, err := getDataInSingleflight("key")
			if err != nil {
				log.Print(err)
				return
			}
			log.Println("receive data: ", data)
		}()
	}
	wg.Wait()
}

// getData 获取数据
func getData(key string) (string, error) {
	data, err := mockGetDataFromCache(key)
	if err != nil {
		if err == errorNotExist {
			// 模拟从db中获取数据
			data, err := mockGetDataFromDB(key)
			if err != nil {
				return "", err
			}

			// TODO: set cache
			return data, nil
		}
		return "", err
	}
	return data, nil
}

// getDataInSingleflight 使用单飞模式获取数据
func getDataInSingleflight(key string) (string, error) {
	data, err := mockGetDataFromCache(key)
	if err != nil {
		if err == errorNotExist {
			// 模拟从db中获取数据
			v, err, _ := g.Do(key, func() (interface{}, error) {
				return mockGetDataFromDB(key)
			})
			if err != nil {
				return "", err
			}
			d, _ := v.(string)
			// TODO: set cache
			return d, nil
		}
		return "", err
	}
	return data, nil
}

// mockGetData 模拟从DB获取数据
func mockGetDataFromDB(key string) (string, error) {
	log.Printf("get %s from db", key)
	return "hello", nil
}

// mockGetDataFromCache 模拟从cache中获取数据
func mockGetDataFromCache(key string) (string, error) {
	if key == "hello" {
		return "hello", nil
	}
	return "", errorNotExist
}
