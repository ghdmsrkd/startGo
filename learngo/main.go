package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/startGo/learngo/accounts"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(str string) (length int, upper string) {
	defer fmt.Println("lendAndUpper is done!")
	length = len(str)
	upper = strings.ToUpper(str)

	return
}

func repeatMe(str ...string) {
	fmt.Println(str)
}

func superAdd(numbers ...int) (total int) {
	total = 0
	for idx, number := range numbers {
		fmt.Println("idx = ", idx)
		total += number
	}
	return
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge > 19 {
		return true
	} else if age < 18 {
		return false
	}
	return false
}

func main() {
	//fmt.Println("Hello World!!")

	//var ask string = "go는 어떠니?"
	//say := "go는 정말 신기하다."

	//something.SayAnything(ask)
	//something.SayAnything(say)

	/**/
	//fmt.Println(multiply(2, 3))
	//strLen, strUpper := lenAndUpper("강은호")
	//fmt.Println(strLen, strUpper)

	/**/
	//repeatMe("강은호님", "안녕하세요", "즐거운 하루입니다.")

	/*for test*/
	//fmt.Println(superAdd(1, 2, 3, 4, 5))

	/*if test*/
	//fmt.Println(canIDrink(17))

	/*pointer test */
	//a := 2
	//b := &a
	//a = 10
	//fmt.Println(a, *b)

	/*arrays test*/
	//name := []string{"aa", "bb", "cc"}
	//name = append(name, "dd")
	//fmt.Println(name)

	/*map Test*/
	//eunho := map[string]string{"name": "kang", "age": "13"}
	//for key, value := range eunho {
	//fmt.Println(key, value)
	//}

	account := accounts.NewAccount("kang")
	account2 := accounts.NewAccount("kim")

	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account, account2)
	fmt.Println(account.Balance())
}
