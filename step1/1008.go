package main

import "fmt"

func main() {
	var a float64 //java -> double
	var b float64
	fmt.Scanf("%f %f", &a, &b)
	fmt.Println(a/b)
}

/*
when you print the float64 up to 20 decimal places, you get a lot more digits than just 0.30...04,
 you get 0.30000000000000004440892098500626161694526672363281 and the rest gets cut off. 
 I'm guessing that with a float32, a lot more gets cut off and it gets rounded to an even 0.3

 Java 의 float / double는 float32및 float64
*/