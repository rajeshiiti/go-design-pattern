package main

import (
	"fmt"
	"sync"
)

type Config struct {
	host string
}

var config *Config
var mutex sync.Mutex

func main() {
	i := 0
	for ; i < 15; i++ {
		fmt.Println("Iteration: ", i)
		go initConfig()
	}
	fmt.Scanln()
	fmt.Println("Completed")
}

func initConfig() {
	if config == nil {
		mutex.Lock()
		defer mutex.Unlock()
		config = &Config{
			host: "localhost",
		}
		fmt.Println("initialized config")
	} else {
		fmt.Println("Config already initialized")
	}

}
