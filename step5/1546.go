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
	var sum float32 = 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &inputs[i])
		if  max < inputs[i] {
			max = inputs[i]
		}
	}
	//fmt.Println(max)
	var newScore = make([]float32, n)
	for i := 0; i < n; i++ {
		//fmt.Println(inputs[i])
		newScore[i] = float32(inputs[i] / max) * 100
		fmt.Println(float32(inputs[i] / max))
		sum += newScore[i]
	}
	fmt.Println(sum)
	fmt.Println(sum / float32(n))
}