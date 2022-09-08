package main

// 此查重算法的时间复杂度为 O(N)
func hasDuplicateValue(data []int) bool {
	existValue := make(map[int]struct{})
	for _, value := range data {
		if _, ok := existValue[value]; ok {
			return true
		}
		existValue[value] = struct{}{}
	}
	return false
}
