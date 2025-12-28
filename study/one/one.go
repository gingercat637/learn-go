package one

import (
	"fmt"
	"os"
)

func Run() {
	fmt.Println("Hello World!")

	fmt.Println("cmd args:")
	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	fmt.Println("use air")
}