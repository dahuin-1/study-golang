package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)
	//defer writer.Flush()

	var sequence  = make([]int, t)
	for i := range sequence {
		fmt.Fscanf(reader, "%d ", &sequence[i])
	}

	var max = 0
	var min = sequence[0]

	for i := 0; i < len(sequence); i++ {
		if max < sequence[i] {
			max = sequence[i]
		}
		if min > sequence[i] {
			min = sequence[i]
		}
	}

	fmt.Printf("%d %d\n", min, max)

}