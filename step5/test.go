package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var numbers = make([]int, 9)
	var max = 0
	var maxIdx = 0
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < len(numbers); i++ {
		fmt.Fscanf(reader, "%d\n", &numbers[i])
		if max < numbers[i] {
			max = numbers[i]
			maxIdx = i
		}
	}
	fmt.Println(max)
	fmt.Println(maxIdx + 1)

}