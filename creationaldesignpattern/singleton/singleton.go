package main

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}
var once = &sync.Once{}

type Config struct {
	// struct data
	Host string
}

var configInstance *Config

func getConfigInstance() *Config {

	once.Do(func() {
		if configInstance == nil {
			fmt.Println("----Creating new config instance-----")
			configInstance = &Config{
				Host: "localhost",
			}
		} else {
			fmt.Println("----Config instance already exist-----")
		}
	})
	return configInstance
}

func getConfigInstanceWithMutex() *Config {
	mutex.Lock()
	defer mutex.Unlock()
	if configInstance == nil {
		fmt.Println("----Creating new config instance-----")
		configInstance = &Config{
			Host: "localhost",
		}
	} else {
		fmt.Println("----Config instance already exist-----")
	}
	return configInstance
}

func main() {
	for i := 0; i < 10; i++ {
		go getConfigInstance()
		// go getConfigInstanceWithMutex()
	}
	// time.Sleep(15 * time.Second)
	fmt.Scanln()
}
