// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 4, 4, 4, 2, 3, 1, 5, 5, 5, 5}
	// 5,4,2,3,1
	res := []int{}
	m := make(map[int]int)
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
			res = append(res, v)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] > res[j]
	})
	fmt.Println(res)
}
