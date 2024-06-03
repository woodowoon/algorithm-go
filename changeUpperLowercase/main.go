package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	fmt.Scan(&s1)

	s2 := strings.Split(s1, "")

	for i := 0; i < len(s2); i++ {
		if IsUpper(s2[i]) {
			fmt.Print(strings.ToLower(s2[i]))
		} else {
			fmt.Print(strings.ToUpper(s2[i]))
		}
	}

}

func IsUpper(s string) bool {
	return strings.ToUpper(s) == s
}
