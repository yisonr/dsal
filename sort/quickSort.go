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
	key := nums[r] // 确定最后一个元素为哨兵(pivot)
	//all in [l,i) < key
	//all in [i,j] > key
	i := l
	j := l
	for j < r { // 遍历每一个元素
		if nums[j] < key { // 确保所有小于 pivot 的元素都在大于 pivot 的元素的前面
			nums[i], nums[j] = nums[j], nums[i]
			i++ // i 标记顺序第一个大于 pivot 的元素的位置
		}
		j++
	}
	nums[i], nums[r] = nums[r], nums[i] // 第一个大于 pivot 的元素和 pivot 互换位置
	return i
}

func main() {
	var Test = []int{1, 3, 2, 4, 9, 2, 19, 11, 8, 7}
	fmt.Println(Test)
	quickSort(Test, 0, len(Test)-1)
	fmt.Println(Test)
}

/*
	TODO: 如何计算时间复杂度和空间复杂度
	TODO: 快排的应用及其变种
	TODO: 分治法
	很多有用的算法在结构是递归的, 为了解决给定的问题, 多次递归的调用自己去
	解决一个相关的子问题; 这些算法通常遵循分治法, 即他们将原问题划分为多个
	规模更小且与原问题相似的子问题, 然后递归的解决子问题, 最后合并子问题的
	答案得到原问题的答案;

	分治法在每一次递归阶段都分为三个步骤:(先微分再积分?)
	- 划分: 将问题划分为一系列规模更小的相似子问题;(微分的思想?)
	- 处理: 递归的解决子问题, 如果子问题的规模足够的小, 则直接解决它;
	- 合并: 将各个子问题的答案合并为原问题的答案;
	其中在处理过程中, 如果子问题规模大到需要递归处理, 则我们称它为递归
	实例(recursive case), 如果子问题规模足够的小, 递归“到达了最低点”，
	则我们称它为基础实例(base case);
	除了解决相似问题的较小规模的子问题外, 还必须解决与原问题不太相同的子问题,
	一般将解决此类子问题作为合并步骤的一部分;
	快速排序算法是分治法的一个实例, 对于待排序的数组nums[p..r]:

	划分: 将数组nums[l..r]划分为两个子数组nums[l..m-1]和nums[m+1..r],
		使得 nums[l..m-1]的每个元素小于或等于nums[m]，nums[m+1..r]的每个
		元素大于或等于nums[m], 计算索引m的值是此划分过程的一部分;
	处理: 递归调用快速排序算法, 排序两个子数组nums[l..m-1]和nums[m+1..r];
	合并: 因为子数组已经排序, 所以不需要将它们合并起来, 整个数组 nums 现在
		已排好序;

	算法最核心的部分是划分阶段:
	初始: i,j都为l, 则数组范围[l,i)和[l,i)均无元素
	保持: 在迭代过程中, 按照nums[j]的值分为两种处理情况:
		- 若nums[j]小于key, 交换nums[i]和nums[j], 同时i后移, j后移
		- 若nums[j]大于等于key, j后移;
	终止: 当j == r时, 循环终止; 之后交换nums[i](当前数组位置中第一个大
		于等于 key的值) 和 nums[r], 使得原nums[r]的值key放入正确位置;
		同时i即是正确位置的值的数组索引号;

	(
		i 标记了顺序第一个大于 nums[r] 的元素的位置
		j 比较了循环元素的索引
	)

	快速排序是高效的排序方法. 平均时间复杂度为O(nlogn), 最坏的情况下其时间
	复杂度为 O(n^2), 且处理阶段不需要额外的存储空间(原址排序);

*/
