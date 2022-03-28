package main

func getSkyline(buildings [][]int) [][]int {
	// method1: 扫描新的建筑物的高度h, 然后再根据x区分
	var skyLine [][]int
	var lastBuild []int
	for i, build := range buildings {
		if i == 0 {
			skyLine = append(skyLine, []int{build[0], build[2]})
			lastBuild = build
			continue
		}
		if build[2] > lastBuild[2] {
			// 新的建筑物比上次的高

			// 1. 若相交, 但不包含, 一定是一个关键点
			if build[0] <= lastBuild[1] && build[1] >= lastBuild[1] {
				lastBuild = build
				skyLine = append(skyLine, []int{build[0], build[2]})
			}
			// 2. 包含, 存在两个关键点
			if build[1] < lastBuild[1] {
				skyLine = append(skyLine, []int{build[0], build[2]})
				skyLine = append(skyLine, []int{build[1], lastBuild[2]})
				lastBuild = build
			}
			// 3. 若不相交, 则存在关键点
			if build[0] > lastBuild[1] {
				skyLine = append(skyLine, []int{lastBuild[1], 0})
				skyLine = append(skyLine, []int{build[0], build[2]})
				lastBuild = build
			}
		} else if build[2] == lastBuild[2] {
			// 高度相同, 两个相邻建筑物之间无高度差
			// 1. 若相交, 但不包含, 则不存在关键点
			if build[0] <= lastBuild[1] && build[1] >= lastBuild[1] {
				lastBuild[1] = build[1]
			}
			// 2. 包含, 不存在关键点
			if build[1] < lastBuild[1] {
				// nothing
			}
			// 3. 若不相交, 则存在关键点
			if build[0] > lastBuild[1] {
				skyLine = append(skyLine, []int{lastBuild[1], 0})
				skyLine = append(skyLine, []int{build[0], build[2]})
				lastBuild = build
			}
		} else {
			// 新的建筑物比上次的低
			// 1. 若相交, 但不包含, 有两个关键点
			if build[0] <= lastBuild[1] && build[1] >= lastBuild[1] {
				skyLine = append(skyLine, []int{lastBuild[1], build[2]})
				skyLine = append(skyLine, []int{build[1], 0})
				lastBuild = build
			}
			// 2. 包含, 不存在关键点
			if build[1] < lastBuild[1] {
				lastBuild[0] = build[0]
				lastBuild[1] = build[1]
			}
			// 3. 若不相交, 则存在关键点
			if build[0] > lastBuild[1] {
				skyLine = append(skyLine, []int{lastBuild[1], 0})
				skyLine = append(skyLine, []int{build[0], build[2]})
				lastBuild = build
			}

		}
	}
	return skyLine
}
