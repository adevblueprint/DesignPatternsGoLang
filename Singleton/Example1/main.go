package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 30; i++ {
		go GetInstance()
	}

	// Stops scanning at a newline or EOF.
	fmt.Scanln()
}
