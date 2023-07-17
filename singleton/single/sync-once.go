package single

import (
	"fmt"
	"sync"
)

var once sync.Once

func GetInstanceOnce() *single {
	if singleInstance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now")
			singleInstance = &single{}
		})
	} else {
		fmt.Println("Single instance already created")
	}

	return singleInstance
}
