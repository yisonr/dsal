package main

/*
给你一个整数数组 nums, 如果任一值在数组中出现至少两次, 返回 true ,
如果数组中每个元素互不相同, 返回 false;

示例 1：
输入：nums = [1,2,3,1]
输出：true

示例 2：
输入：nums = [1,2,3,4]
输出：false

示例 3：
输入：nums = [1,1,1,3,3,4,3,2,4,2]
输出：true

提示：
1 <= nums.length <= 10^5
-10^9 <= nums[i] <= 10^9

https://leetcode-cn.com/problems/contains-duplicate
*/

// 占用内存: 8.2MB
func containsDuplicate(nums []int) bool {
	var elemMap = make(map[int]struct{})
	for _, num := range nums {
		if _, ok := elemMap[num]; ok {
			return true
		} else {
			elemMap[num] = struct{}{}
		}
	}
	return false
}

// leetcode
// 占用内存: 7720kb
// func containsDuplicate(nums []int) bool {
// 	sort.Ints(nums)
// 	fmt.Println(nums)
// 	for i:=0;i<len(nums);i++{
// 		if i == len(nums)-1{
// 			return false
// 		}
// 		if nums[i] == nums[i+1]{
// 			return true
// 		}
// 	}
// 	return false
// }

// TODO: map 原理, 对内存的占用; 分析上述两种写法对内存的占用区别?
// TODO: 分析如果提供的数组长度和元素的大小在题目条件的临界值时, 对内存的占用?

func main() {
	containsDuplicate([]int{1, 2, 3, 1})
}
