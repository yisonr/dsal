package main

/*
	给你一个整数数组 nums 和一个整数 k, 判断数组中是否存在两个不同的索引
	i 和 j 满足 nums[i] == nums[j] 且 abs(i - j) <= k; 如果存在, 返回 true,
	否则, 返回 false;
*/

// 内存占用: 11MB
// func containsNearbyDuplicate(nums []int, k int) bool {
// 	var existMap = make(map[int]int)
// 	for i, num := range nums {
// 		if j, ok := existMap[num]; ok {
// 			// nums[i] == nums[j] 存在
// 			if i-j <= k && i-j >= -1*k {
// 				return true
// 			}
// 			// 索引距离超出k， 更新索引
// 		}
// 		existMap[num] = i
// 	}
// 	return false
// }

// leetcode
// 内存占用: 7608KB
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 {
		return false
	}
	numMap := map[int]struct{}{}
	for i, num := range nums {
		if i > k {
			delete(numMap, nums[i-k-1])
		}
		if _, ok := numMap[num]; ok {
			return true
		}
		numMap[num] = struct{}{}
	}
	return false
}
