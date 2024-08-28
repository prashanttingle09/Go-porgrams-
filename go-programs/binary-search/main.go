package main

import "fmt"

func main() {
	sums := []int{3, 7, 9, 10}
	n := 5
	fmt.Println("binary search:", BinarySearch(n, sums))
}

func BinarySearch(target int, nums []int) int {
	l := 0
	r := len(nums) - 1
	if len(nums) < 1 {
		return -1
	}
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
