package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fscanln(reader, &t)
	var a, b int
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintf(writer, "Case #%d: %d\n", i+1, a+b)
	}


}