package main

import (
	"fmt"
	"sync"
)

type Config struct {
	host string
}

var config *Config
var once sync.Once

func main() {
	i := 0
	for ; i < 15; i++ {
		fmt.Println("Iteration: ", i)
		go initConfig()
	}
	fmt.Scanln()
	fmt.Println("Completed: config: ", config)
}

func initConfig() {
	once.Do(func() {
		config = &Config{
			host: "localhost",
		}
		fmt.Println("Initialised config...")
	})
}
