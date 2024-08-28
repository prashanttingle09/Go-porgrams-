package stringproblems

import "fmt"

func main() {
	s := "abccba"
	fmt.Println("leftmost repeating character index is:", leftMost(s))
}

func leftMost(s string) int {
	n := len(s)
	res := -1
	if n < 1 {
		return res
	}
	var arr [256]bool
	for i := n - 1; i >= 0; i-- {
		if arr[s[i]] == false {
			arr[s[i]] = true
		} else {
			res = i
		}
	}
	return res
}

//abccba output ==> 1 which is first repeating charcter index
