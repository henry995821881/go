package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args
	arg1 := args[1]

	err := os.Remove(arg1)

	if err != nil {
		fmt.Println(err)
	}

}
