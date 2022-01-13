package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b) //앰퍼센드 입력 받아서 저장하겠다
	fmt.Println(a+b)
}

//printf : 뒤에 변수를 같이 프린트
//%d : 십진수 정수
//Scanf : 주소값 스캔