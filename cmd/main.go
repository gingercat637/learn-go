package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!")

	fmt.Println("cmd args:")
	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	fmt.Println("use air")
}
