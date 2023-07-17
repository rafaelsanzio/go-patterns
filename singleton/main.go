package main

import (
	"fmt"

	"github.com/rafaelsanzio/go-patterns/singleton/single"
)

func main() {
	for i := 0; i < 30; i++ {
		//go single.GetInstance()
		go single.GetInstanceOnce()
	}

	fmt.Scanln()
}
