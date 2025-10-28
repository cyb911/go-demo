package main

import "fmt"

/*
The number that appears only once.
*/
func main() {
	nums := []int{4, 2, 4, 3, 0, 0, 3, 9, 9}

	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	for num, count := range countMap {
		if count == 1 {
			fmt.Println("num:", num, "count:", count)
		}
	}

}
