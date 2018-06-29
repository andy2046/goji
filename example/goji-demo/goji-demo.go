package main

import (
	"fmt"

	G "github.com/andy2046/goji"
)

func main() {
	preFilter := []int{5, 4, 3, 2, 1}
	var postFilter []int

	G.Filter(
		len(preFilter),
		func(i int) bool {
			if preFilter[i] > 2 {
				return true
			}
			return false
		},
		func(nums ...int) {
			for _, num := range nums {
				postFilter = append(postFilter, preFilter[num])
			}
		})
	fmt.Println(preFilter)
	fmt.Println(postFilter)
}
