package main

import "fmt"

/*
	冒泡排序的执行步骤分为两种:
	比较: 比较两个数看哪个最大
		N 个元素需要的比较次数为: (N-1) + (N-2) + ... + 1
	交换: 交换两个数的位置以使它们按顺序排列
		如果数组不只是随机打乱, 而是完全反序, 在这种最坏的情况下,
		每次比较过后都得进行一次交换, 则交换的次数和比较的次数相同

	则一个含有10个元素的数组, 需要45次比较, 45次交换, 共90步
	  一个含有20个元素的数组, 需要190次比较, 190次交换, 共380步
	  ......

	效率太低了, 元素量呈倍数增长, 步数却呈指数增长; 观察规律可以发现
	随着N的增长, 步数大约增长为N^2;
	因此描述冒泡排序效率的大O计法是O(N^2), 也被叫做二次时间

	大O测量的是步数与数据量的关系

*/

// 冒泡排序
func BullleSort(data []int) []int {
	unsortedUntilIndex := len(data) - 1
	sorted := false
	compareStep := 0
	swapStep := 0

	for !sorted {
		sorted = true
		for i := range data[:unsortedUntilIndex] {
			compareStep++
			if data[i] > data[i+1] {
				sorted = false // 当发生任何交换时, 表明还需要下一次的冒泡
				// 没有任何交换时, 表示 unsortedUntilIndex 该索引之前的数据都
				// 已经排好序
				data[i], data[i+1] = data[i+1], data[i]
				swapStep++
			}
		}
		unsortedUntilIndex = unsortedUntilIndex - 1 // 表示该轮冒泡结束了,
		// 该轮冒泡中最右侧的值处于正确的位置
	}

	fmt.Println(compareStep, swapStep)
	return data
}
