package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	var a, b, c int
	var counts = make([]int, 10)
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n%d\n%d\n", &a, &b, &c)
	var number = a * b * c
	for true {
		counts[number % 10]++ 
		number = number / 10
		if number == 0 {
			break
		}
	}
	for i := range counts {
		fmt.Println(counts[i])
	}
}