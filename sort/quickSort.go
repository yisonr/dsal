package main

import "fmt"

func quickSort(nums []int, l, r int) { //[l,r]
	if l < r {
		m := partition(nums, l, r)
		quickSort(nums, l, m-1)
		quickSort(nums, m+1, r)
	}
}

func partition(nums []int, l int, r int) int {
	key := nums[r]
	//all in [l,i) < key
	//all in [i,j] > key
	i := l
	j := l
	for j < r {
		if nums[j] < key {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}

func main() {
	var Test = []int{1, 3, 2, 4, 9, 2, 4, 5}
	fmt.Println(Test)
	quickSort(Test, 0, len(Test)-1)
	fmt.Println(Test)
}
