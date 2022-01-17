package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	//var n int
	reader := bufio.NewReader(os.Stdin)
	//fmt.Fscanln(reader, &n)

	var sequence = make([]int, 9)
	//var mymap map[int]int
	var mymap = make(map[int] int)
	for i := 0; i < 9; i++ {
		fmt.Fscan(reader, &sequence[i])
		mymap[sequence[i]] = i
	}

	sort.Slice(sequence, func(i, j int) bool {
		return sequence[i] < sequence[j]
	})

	fmt.Printf("%d\n", sequence[8])
	fmt.Printf("%d\n", mymap[sequence[8]])
}