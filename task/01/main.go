package main

import "fmt"

func main() {
	// 声明一个数组
	nums := []int{4, 2, 4, 3, 0, 0, 3, 9, 9}

	// 使用map进行技术
	countMap := make(map[int]int)

	for _, num := range nums {
		countMap[num]++
	}

	//找出map，次数=1的key
	for num, count := range countMap {
		if count == 1 {
			fmt.Println("num:", num, "count:", count)
		}
	}

}
