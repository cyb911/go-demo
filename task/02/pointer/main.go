package main

import "fmt"

func main() {
	//num := 20
	//fmt.Println("修改前：", num)
	//addTen(&num)
	//fmt.Println("修改后：", num)

	values := []int{1, 2, 3, 4, 5}
	fmt.Println("修改前:", values)
	multiplyByTwo(&values)
	fmt.Println("修改后:", values)
}

/*
1.编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，
然后在主函数中调用该函数并输出修改后的值。
*/
func addTen(num *int) {
	*num = *num + 10
}

/*
2.实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2
*/
func multiplyByTwo(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= (*nums)[i] * 2
	}
}
