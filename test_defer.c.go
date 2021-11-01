package main

import "fmt"

func main1() {
	var arr = [4]int{}
	for i, _ := range arr {
		ii := i
		defer func() {fmt.Println(ii)}()
	}
}
