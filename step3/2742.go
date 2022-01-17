package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanf(reader,"%d", &t)
	for i := t; i > 0; i-- {
		fmt.Fprintln(writer,i)
	}
	writer.Flush()	
}