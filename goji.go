// Package goji provide functional utility for Slice.
package goji

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Range creates a range of numbers progressing from zero up to, but not including end.
func Range(end int) []struct{} {
	return make([]struct{}, end)
}

type sorter struct {
	len  int
	swap func(i, j int)
	less func(i, j int) bool
}

func (x sorter) Len() int           { return x.len }
func (x sorter) Swap(i, j int)      { x.swap(i, j) }
func (x sorter) Less(i, j int) bool { return x.less(i, j) }

// Sort sorts elements, n is the number of elements.
// less compares the elements with indexes i and j.
// swap swaps the elements with indexes i and j.
func Sort(n int, swap func(int, int), less func(int, int) bool) {
	sort.Sort(sorter{len: n, swap: swap, less: less})
}

// Reverse reverses an array/slice, n is the number of elements.
// swap swaps the elements with indexes i and j.
func Reverse(n int, swap func(int, int)) {
	len := n - 1
	for i := 0; i < len; i++ {
		swap(i, len)
		len--
	}
}

// Filter filters elements, n is the number of elements.
// predict check element with index i and returns true or false.
// handlers receive accumulated indexes idx from predict, if any.
func Filter(n int, predict func(int) bool, handlers ...func(...int)) []int {
	var idx []int

	for i := 0; i < n; i++ {
		if ok := predict(i); ok {
			idx = append(idx, i)
		}
	}

	for _, handler := range handlers {
		handler(idx...)
	}

	return idx
}

// Includes determines whether an array/slice includes a certain element, returning true or false.
// n is the number of elements. comp determines whether the given element with index j is found.
func Includes(n int, comp func(int) bool) bool {
	for i := 0; i < n; i++ {
		if comp(i) {
			return true
		}
	}
	return false
}

// Map calls a provided function m on every element.
// n is the number of elements. m map element with index i.
func Map(n int, m func(int)) {
	for i := 0; i < n; i++ {
		m(i)
	}
}

// Shuffle randomizes the order of elements.
// n is the number of elements. swap swaps the elements with indexes i and j.
func Shuffle(n int, swap func(int, int)) {
	rand.Shuffle(n, swap)
}

// Concat merges two slices.
func Concat() {
	fmt.Println("no Concat, use append.")
}

// Every tests whether all elements pass the test implemented by predict.
// n is the number of elements. j is the index of array/slice.
func Every(n int, predict func(int) bool) bool {
	for i := 0; i < n; i++ {
		if ok := predict(i); !ok {
			return false
		}
	}

	return true
}

// Some tests whether at least one element pass the test implemented by predict.
// n is the number of elements. j is the index of array/slice.
func Some(n int, predict func(int) bool) bool {
	for i := 0; i < n; i++ {
		if ok := predict(i); ok {
			return true
		}
	}

	return false
}

// Slice returns a portion of array/slice selected from start to end (end not included).
// n is the number of elements. start and end indexes are optional, default to zero and n.
// start and end indexes are passed to slice func.
func Slice(slice func(int, int), n int, startEndIndex ...int) {
	start, end := 0, n

	switch len(startEndIndex) {
	case 1:
		start = startEndIndex[0]
	case 2:
		start, end = startEndIndex[0], startEndIndex[1]
	}

	start = Min(start, n)
	if start < 0 {
		start = Max(n+start, 0)
	}

	end = Min(end, n)
	if end < 0 {
		end = Max(n+end, 0)
	}

	if start < end {
		slice(start, end)
	}
}

// Join joins all elements of array/slice into a string.
// The separator string sep is placed between elements in the resulting string.
// n is the number of elements. func join receive element index i and return a string.
func Join(join func(int) string, n int, sep ...string) string {
	v := []string{}
	for i := 0; i < n; i++ {
		v = append(v, join(i))
	}

	s := ","
	for _, se := range sep {
		s = se
	}
	return strings.Join(v, s)
}

// Fill fills all elements from a start index to an end index with a static value.
// the end index is not included.
// n is the number of elements. start and end index are optional, default to zero and n.
func Fill(fill func(int), n int, startEndIndex ...int) {
	start, end := 0, n

	switch len(startEndIndex) {
	case 1:
		start = startEndIndex[0]
	case 2:
		start, end = startEndIndex[0], startEndIndex[1]
	}

	start = Min(start, n)
	if start < 0 {
		start = Max(n+start, 0)
	}

	end = Min(end, n)
	if end < 0 {
		end = Max(n+end, 0)
	}

	for start < end {
		fill(start)
		start++
	}
}

// FindIndex returns the index of the first element that satisfies the provided testing function.
// Otherwise -1 is returned. n is the number of elements.
func FindIndex(n int, testing func(int) bool, fromIndex ...int) int {
	start := 0

	for _, id := range fromIndex {
		start = id
	}

	start = Min(start, n)
	if start < 0 {
		start = Max(n+start, 0)
	}

	for i := start; i < n; i++ {
		if ok := testing(i); ok {
			return i
		}
	}

	return -1
}

// FindLastIndex iterates over elements from right to left.
func FindLastIndex(n int, testing func(int) bool, fromIndex ...int) int {
	start := n - 1

	for _, id := range fromIndex {
		start = id
	}

	start = Min(start, n)
	if start < 0 {
		start = Max(n+start, 0)
	}

	for i := start; i >= 0; i-- {
		if ok := testing(i); ok {
			return i
		}
	}

	return -1
}

// Difference returns difference of array/slice I's values not included in the other given array J.
// lenI, lenJ is the length of array/slice I, J.
// comparator compares values from array/slice I, J.
func Difference(lenI, lenJ int, comparator func(int, int) bool, handlers ...func(...int)) []int {
	idx := []int{}
	cp := func(i int) func(j int) bool {
		return func(j int) bool {
			return comparator(i, j)
		}
	}

	for i := 0; i < lenI; i++ {
		if ok := Includes(lenJ, cp(i)); !ok {
			idx = append(idx, i)
		}
	}

	for _, handler := range handlers {
		handler(idx...)
	}

	return idx
}

// Intersection returns values that are included in both given array/slice.
func Intersection(lenI, lenJ int, comparator func(int, int) bool, handlers ...func(...int)) []int {
	idx := []int{}
	cp := func(i int) func(j int) bool {
		return func(j int) bool {
			return comparator(i, j)
		}
	}

	for i := 0; i < lenI; i++ {
		if ok := Includes(lenJ, cp(i)); ok {
			idx = append(idx, i)
		}
	}

	for _, handler := range handlers {
		handler(idx...)
	}

	return idx
}
