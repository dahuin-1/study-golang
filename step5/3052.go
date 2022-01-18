package main

import (
	"fmt"
	"os"
	"bufio"
)

type ns struct{}

func add(s map[int]ns, v int) map[int]ns { //홈메이드 set
	s[v] = ns{}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := make(map[int]ns)
	var a int
	for i := 0; i < 10; i++ {
		fmt.Fscanf(reader, "%d\n", &a)
		var number = a % 42
		s = add(s, number)
	}
    //fmt.Println(s)
	fmt.Println(len(s))
}