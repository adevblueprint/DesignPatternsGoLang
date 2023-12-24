package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var instance *single

func GetInstance() *single {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance.")
			instance = &single{}
		} else {
			fmt.Println("Single instance already exists.")
		}
	} else {
		fmt.Println("Single instance already avaliable.")
	}

	return instance
}
