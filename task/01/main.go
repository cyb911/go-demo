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
	//strs1 := []string{"flower", "flow", "flight"}
	//strs2 := []string{"com-dog", "com-racecar", "com-car"}
	//fmt.Println(longestCommonPrefix(strs1))
	//fmt.Println(longestCommonPrefix(strs2))

	//fmt.Println("两数之和:")
	//var nums = []int{3, 2, 4}
	//var result = twoSum(nums, 6)
	//fmt.Println(result)

	//nums := []int{1, 1, 2, 2, 3}
	//k := removeDuplicates(nums)
	//fmt.Println("唯一元素个数:", k)
	//fmt.Println("去重后数组:", nums[:k])

	//fmt.Println(plusOne([]int{1, 2, 3}))    // [1 2 4]
	//fmt.Println(plusOne([]int{4, 3, 2, 1})) // [4 3 2 2]
	//fmt.Println(plusOne([]int{9, 9, 9}))    // [1 0 0 0]

	fmt.Println(isValid("([]){}"))

}

/*
有效的括号
*/
func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 { // 奇数一定是无效的
		return false
	}

	// 定义一个切片，用作栈的存储结构
	stack := make([]byte, 0, length/2)

	for i := 0; i < length; i++ {
		t := s[i]
		switch t {
		case '(':
			stack = append(stack, ')') // 左边符号入栈，入栈值为右符号，作为预期值
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case ')', ']', '}': // 右边符号，取出栈数据对比是否一致。
			sn := len(stack)
			index := sn - 1
			b := stack[index]
			if sn == 0 || b != t {
				return false
			}
			stack = stack[:index] // 出栈操作
		default:
			return false
		}
	}
	return len(stack) == 0
}

/*
加一
*/
func plusOne(digits []int) []int {
	n := len(digits)

	if n == 0 {
		return digits
	}

	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	// 处理9，99，999 情况。

	result := make([]int, n+1)
	result[0] = 1
	return result
}

/*
删除有序数组中的重复项
*/
func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	i := 0
	for j := 1; j < length; j++ { //比较相邻的数据是否相等，不相等说明是新的元素，复制到数组下标位置i + 1 位置
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
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

/*
两数之和
*/
func twoSum(nums []int, target int) []int {
	var length = len(nums)
	if length == 0 {
		return nil
	}

	var tagMap = make(map[int]int)

	for i := 0; i < length; i++ {
		tagMap[nums[i]] = i // 数组元素作为key,下标作为value
	}

	for i := 0; i < length; i++ {
		// 目标值减去数组中当前元素值，在tagMap hash表中找出对应的数组索引下标值
		t1 := nums[i]
		t2 := target - t1
		index, ok := tagMap[t2]
		if ok {
			if index == i { // 当前数组下标不能与tagMap 中记录的下标值一致
				continue
			}
			return []int{i, index}
		}
	}
	return nil

}
