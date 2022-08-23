package main

import "fmt"

// 有序数组中查找元素
// 二分查找
func linear_search(data []int, value int) (int, bool) {
	lower_bound := 0             // 二分查找的下界
	upper_bound := len(data) - 1 // 二分查找的上界

	for lower_bound <= upper_bound {
		midpoint := (upper_bound + lower_bound) / 2
		midValue := data[midpoint]
		if value < midValue {
			upper_bound = midpoint - 1
		} else if value > midValue {
			lower_bound = midpoint + 1
		} else {
			return midpoint, true
		}
	}
	return 0, false
}

func main() {
	var data = []int{1, 2, 5, 7, 12, 23, 34, 123, 778}
	fmt.Println(linear_search(data, 12)) // 4 true
}
