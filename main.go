package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2) == 2)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	mergedLen := len(nums1) + len(nums2)
	mergedNums := merge(nums1, nums2)
	if mergedLen > 1 {
		if mergedLen%2 == 1 {
			result = float64(mergedNums[mergedLen/2])
		} else {
			result = (float64(mergedNums[mergedLen/2]) + float64(mergedNums[mergedLen/2-1])) / 2
		}
	} else {
		result = float64(mergedNums[0])
	}
	return result
}
