package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if isOdd(a) { // 홀수
		fmt.Println(a, "is even")
	} else {
		fmt.Println(a, "is odd")
	}
}

func isOdd(number int) bool {
	return number%2 == 0
}
