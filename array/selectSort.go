package main

/*
	比较排序
	比较: N个元素需要 (N-1) + (N-2) + (N-3) + ... + 1 次比较
	交换: 每轮的交换最多只有一次, 如果该轮的最小值已在正确位置, 就无须交换
	否则要做1次交换; 相比之下, 冒泡排序在最坏情况(完全逆序)时, 每次比较过后都
	要进行1次交换;

	根据实际的运行步数可以看到选择排序的步数大概只有冒泡排序的一半, 即选择
	排序比冒泡排序快一倍; 但选择排序的大O计法和冒泡排序是一样的, 即 O(N^2)
	而非 O(N^2/2); 因为大O计法的一条重要规则为: 大O记法忽略常数; 即大O计
	法不包含一般数字, 除非是指数;

	尽管不能比较冒泡排序和选择排序, 大O还是很重要的, 因为它能够区分不同算法
	的长期增长率; 当数据量达到一定程度时, O(N)的算法就会永远快过 O(N2), 无论
	这个 O(N)实际上是O(2N)还是O(100N); 即使是O(100N), 这个临界点也是存在的;
	可以从极限的角度理解, 大O计法表示的是随着数据量的增大所需步数的增长的趋势;

	大O记法只表明对于不同分类, 存在一临界点, 在这一点之后, 一类算法会快于另
	一类, 并永远保持下去l 至于这个点在哪里, 大O并不关心;


	所以大O是极为有用的工具, 当两种算法落在不同的大O类别时, 你就知道应该选
	择哪种; 因为在大数据的情况下必然存在一临界点使这两种算法的速度永远区分开来;
	同时也要根据不同的数据量选择适合的大O算法;
	对于大O同类的算法, 还需要进一步的解析才能分辨出具体差异;

*/

func SelectSort(data []int) []int {
	for i := range data {
		lowestNumberIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[lowestNumberIndex] {
				lowestNumberIndex = j
			}
		}

		if lowestNumberIndex != i {
			data[i], data[lowestNumberIndex] = data[lowestNumberIndex], data[i]
		}
	}

	return data
}
