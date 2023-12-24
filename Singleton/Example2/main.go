package main

import (
	"fmt"
	"SingletonExample2/singleton"
)

func main() {

	for i := 0; i < 30; i++ {
		_ = singleton.GetInstance()
		
	}

	// Stops scanning at a newline or EOF.
	fmt.Scanln()
}
