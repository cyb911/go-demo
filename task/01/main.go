package main

import "fmt"

func main() {
	//nums := []int{4, 2, 4, 3, 0, 0, 3, 9, 9}
	//fmt.Println(findAppearsOnce(nums))
	//
	//fmt.Println(isPalindrome(121))
	//fmt.Println(isPalindrome(-121))
	//fmt.Println(isPalindrome(10))
	//fmt.Println(isPalindrome(12321))
	strs1 := []string{"flower", "flow", "flight"}
	strs2 := []string{"com-dog", "com-racecar", "com-car"}
	fmt.Println(longestCommonPrefix(strs1))
	fmt.Println(longestCommonPrefix(strs2))
}

/*
只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构
来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func findAppearsOnce(nums []int) int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}

	return -1
}

/*
回文数
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	a := x
	b := 0

	// 获取反转后的数字b，通过去摸法获取数字的每一个数,等到一个反转后的值，最后判断是否和原值相等。
	for a != 0 {
		t := a % 10
		b = b*10 + t
		a /= 10
	}

	return x == b
}

/*
最长公共前缀
Case1: ["flower","flow","flight"]
Case2: ["dog","racecar","car"]
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = commonPrefix(prefix, strs[i])
		if prefix == "" { // 任意两个字符之间不存公共前缀，则任务整个集合中的元素不存在公共前缀。
			return ""
		}
	}
	return prefix
}

/*
两个字符串中的公共前缀（适用于 ASCII）
*/
func commonPrefix(a, b string) string {
	// 字符串进行字节数组元素比较
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			return a[:i]
		}
	}
	return a[:minLen]
}
