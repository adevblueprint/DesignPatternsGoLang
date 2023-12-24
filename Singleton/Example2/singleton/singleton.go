package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type single struct {
}

var instance *single

func GetInstance() *single {
	if instance == nil {
		once.Do(func() {
				fmt.Println("Creating single instance.")
				instance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return instance
}
