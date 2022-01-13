package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	if a >= 90 && a <= 100 {
		fmt.Println("A")
	} else if (a<=89 && a>=80){
		fmt.Println("B")
	} else if (a<=79 && a>=70){
		fmt.Println("C")
	} else if (a<=69 && a>=60){
		fmt.Println("D")
	}  else {
		fmt.Println("F")
	}
}