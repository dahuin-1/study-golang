package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d",&t)
	var a, b int
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d",&a, &b)
		fmt.Println(a+b)
	}
}