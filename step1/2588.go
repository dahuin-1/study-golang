package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
    // b % 10 = 1의 자리수 이하의 수
	// b % 100 = 10의 자리수 이하의 수
	fmt.Println(a * (b % 10))
	fmt.Println(a * (b % 100 - b % 10)/10)
	fmt.Println(a * (b - b % 100)/100)
	fmt.Println(a * b)
}


/*
Scanln	공백으로 구분하여 입력
Scan	공백과 개행으로 구분하여 입력
Scanf	포멧 지정자를 이용하여 개발자가 원하는 형태로 입력
*/