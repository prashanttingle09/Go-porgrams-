package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// inputStr, _ := reader.ReadString('\n')
	// inputStr = strings.TrimSpace(inputStr)
	// stringSlice := strings.Split(inputStr, " ")
	// nums := StringToInt(stringSlice)
	nums := []int{2, 7, 5, 5, 7, 3, 5, 7}
	GenTriplet(nums, 15)
}
func StringToInt(stringSlice []string) []int {
	var nums []int
	for _, str := range stringSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func GenTriplet(nums []int, target int) []int {
	res := []int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	n := len(nums)
	for i := 0; i < n-2; i++ {
		l := i + 1
		r := n - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				res = append(res, i, l, r)
			} else if sum > target {
				r--
			} else {
				l++
			}
		}
	}
	fmt.Println("sorted slice", res)
	return res
}
