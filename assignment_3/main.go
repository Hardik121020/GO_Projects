package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile(os.Args[1])

	if err != nil {
		fmt.Errorf("Error: ", err)
	}
	os.Stdout.Write(data)
}
