package main

/*
	TODO: 滑动窗口, 二分排序
	给你一个整数数组 nums 和两个整数 k 和 t, 请你判断是否存在 两个不同下标 i 和 j,
	使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k;
	如果存在则返回 true,不存在返回 false;

	https://leetcode-cn.com/problems/contains-duplicate-iii

*/

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i, _ := range nums {
		if i == 0 {
			continue
		}
		start := 0
		if i >= k {
			start = i - k
		}
		for j := start; j < i; j++ {
			if nums[i]-nums[j] >= (-1*t) && nums[i]-nums[j] <= t {
				return true
			}
		}
	}
	return false
}
