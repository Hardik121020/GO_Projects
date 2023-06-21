package main

import "fmt"

func mapPrint(m map[string]string) {
	for key, value := range m {
		fmt.Println(key + "    " + value)
	}
}
