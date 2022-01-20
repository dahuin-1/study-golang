package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	var inputs = make([]int, n)
	var max = 0
	var sum = 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &inputs[i])
		if max < inputs[i] {
			max = inputs[i]
		}
	}
	var newScore = make([]int, n)
	for i := 0; i < n; i++ {
		newScore[i] = inputs[i] / max * 100
		sum += newScore[i]
	}
	fmt.Println(sum / n)
}