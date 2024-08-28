package stringproblems

import "sort"

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

func LongestCommonPrefix(strs []string) string {
	str1 := strs[0]
	str2 := strs[len(strs)-1]
	sort.Strings(strs)
	res := common(str1, str2)
	return res
}
func common(str1 string, str2 string) string {
	res := ""
	n := 0
	if len(str1) > len(str2) {
		n = len(str2)
	} else {
		n = len(str1)
	}
	for i := 0; i < n; i++ {
		if str1[i] == str2[i] {
			res = res + string(str1[i])
		} else {
			break
		}
	}
	return res
}
