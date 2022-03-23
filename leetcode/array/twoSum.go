package main

import "fmt"

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]

*/

// 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

// 事件复杂度: o(n)
func twoSum1(nums []int, target int) []int {
	var num = make(map[int]int)
	for i, val := range nums {
		another := target - val
		if index, ok := num[another]; ok {
			return []int{index, i}
		}
		num[val] = i
	}

	return nil
}

// 这里考虑会有多个结果的情况
func twoSum2(nums []int, target int) []int {
	result := []int{}
	for index1, val1 := range nums {
		for index2, val2 := range nums[index1:] {
			if val1+val2 == target {
				result = append(result, []int{index1, index2}...)
			}
		}
	}
	return result
}
