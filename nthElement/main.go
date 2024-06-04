package main

func main() {
	num_list := [5]int{0, 1, 2, 3, 4}
	n := 2

	solution(num_list[:], n)
}

func solution(num_list []int, n int) []int {
	var new_num_list []int

	for i := n - 1; i < len(num_list); i++ {
		new_num_list = append(new_num_list, num_list[i])
	}

	return new_num_list
}
