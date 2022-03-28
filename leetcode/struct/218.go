package main

import (
	"container/heap"
	"sort"
)

// 未通过
// func getSkyline(buildings [][]int) [][]int {
// 	// method1: 扫描新的建筑物的高度h, 然后再根据x区分
// 	var skyLine [][]int
// 	var lastBuild []int
// 	for i, build := range buildings {
// 		if i == 0 {
// 			skyLine = append(skyLine, []int{build[0], build[2]})
// 			lastBuild = build
// 			continue
// 		}
// 		if build[2] > lastBuild[2] {
// 			// 新的建筑物比上次的高

// 			// 1. 若相交, 但不包含, 一定是一个关键点
// 			if build[0] <= lastBuild[1] && build[1] >= lastBuild[1] {
// 				lastBuild = build
// 				skyLine = append(skyLine, []int{build[0], build[2]})
// 			}
// 			// 2. 包含, 存在两个关键点
// 			if build[1] < lastBuild[1] {
// 				skyLine = append(skyLine, []int{build[0], build[2]})
// 				skyLine = append(skyLine, []int{build[1], lastBuild[2]})
// 				lastBuild = build
// 			}
// 			// 3. 若不相交, 则存在关键点
// 			if build[0] > lastBuild[1] {
// 				skyLine = append(skyLine, []int{lastBuild[1], 0})
// 				skyLine = append(skyLine, []int{build[0], build[2]})
// 				lastBuild = build
// 			}
// 		} else if build[2] == lastBuild[2] {
// 			// 高度相同, 两个相邻建筑物之间无高度差
// 			// 1. 若相交, 但不包含, 则不存在关键点
// 			if build[0] <= lastBuild[1] && build[1] >= lastBuild[1] {
// 				lastBuild[1] = build[1]
// 			}
// 			// 2. 包含, 不存在关键点
// 			if build[1] < lastBuild[1] {
// 				// nothing
// 			}
// 			// 3. 若不相交, 则存在关键点
// 			if build[0] > lastBuild[1] {
// 				skyLine = append(skyLine, []int{lastBuild[1], 0})
// 				skyLine = append(skyLine, []int{build[0], build[2]})
// 				lastBuild = build
// 			}
// 		} else {
// 			// 新的建筑物比上次的低
// 			// 1. 若相交, 但不包含, 有两个关键点
// 			if build[0] <= lastBuild[1] && build[1] >= lastBuild[1] {
// 				skyLine = append(skyLine, []int{lastBuild[1], build[2]})
// 				skyLine = append(skyLine, []int{build[1], 0})
// 				lastBuild = build
// 			}
// 			// 2. 包含, 不存在关键点
// 			if build[1] < lastBuild[1] {
// 				lastBuild[0] = build[0]
// 				lastBuild[1] = build[1]
// 			}
// 			// 3. 若不相交, 则存在关键点
// 			if build[0] > lastBuild[1] {
// 				skyLine = append(skyLine, []int{lastBuild[1], 0})
// 				skyLine = append(skyLine, []int{build[0], build[2]})
// 				lastBuild = build
// 			}

// 		}
// 	}
// 	return skyLine
// }

// 扫描线+优先队列
// TODO: 优先队列, container/heap, sort
// TODO: 理清本题的思路
type pair struct{ right, height int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getSkyline(buildings [][]int) (ans [][]int) {
	n := len(buildings)
	boundaries := make([]int, 0, n*2)
	for _, building := range buildings {
		boundaries = append(boundaries, building[0], building[1])
	}
	sort.Ints(boundaries)

	idx := 0
	h := hp{}
	for _, boundary := range boundaries {
		for idx < n && buildings[idx][0] <= boundary {
			heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
			idx++
		}
		for len(h) > 0 && h[0].right <= boundary {
			heap.Pop(&h)
		}

		maxn := 0
		if len(h) > 0 {
			maxn = h[0].height
		}
		if len(ans) == 0 || maxn != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, maxn})
		}
	}
	return
}

// https://leetcode-cn.com/problems/the-skyline-problem/solution/tian-ji-xian-wen-ti-by-leetcode-solution-ozse/
