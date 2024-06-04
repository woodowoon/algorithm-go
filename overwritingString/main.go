package main

import "fmt"

func main() {
	var my_string, overwrite_string string
	var s int
	fmt.Scan(&my_string, &overwrite_string, &s)

	fmt.Println(solution(my_string, overwrite_string, s))
}

func solution(my_string, overwrite_string string, s int) string {
	// 문자열 나누기 첫번째 - s 번째까지
	pre := my_string[:s]
	// 문자열 나누기
	post := my_string[s+len(overwrite_string):]
	return pre + overwrite_string + post
}
