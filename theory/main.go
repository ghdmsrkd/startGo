package main

import (
	"fmt"
	"strings"
)

func add(a int, b int) int {
	defer doneBot()
	return a + b
}

func multiple(a, b int) (int) {
	defer doneBot()
	return a * b
}

func upperAndLen(a string) (upper string, length int){
	// lend() 함수는 한들을 3으로 계산 한다..?!
	upper = strings.ToUpper(a)
	length = len(a) 
	return 
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func doneBot() {
	fmt.Println("i'm done")
}

func main() {
	a := 1
	var b = 2
	var c int = 3
	fmt.Println(a, b, c)

	aa := 1
	const bb = 2
	const cc = 3
	fmt.Println(aa,bb)

	fmt.Println(add(2,2))
	fmt.Println(multiple(2,3))

	upper, len := upperAndLen("안녕 Go!!")
	fmt.Println(upper, len)

	repeatMe("a", "b", "c", "d")
}