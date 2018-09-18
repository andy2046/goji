// Package intutil provides functional utils for int.
package intutil

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"

	g "github.com/andy2046/goji"
)

// Int for int type slice.
type Int struct{}

// IInt is the Int interface.
type IInt interface {
	Filter(ints []int, predict func(int) bool) []int
	Find(ints []int, predict func(int) bool) (int, bool)
	FindIndex(ints []int, predict func(int) bool, fromIndex ...int) int
	FindLastIndex(ints []int, predict func(int) bool, fromIndex ...int) int
	ForEach(ints []int, provided func(int))
	Includes(ints []int, e int) bool
	IndexOf(ints []int, e int, fromIndex ...int) int
	LastIndexOf(ints []int, e int, fromIndex ...int) int
	From(ints []int, mapFn ...func(int) int) []int
	Of(ints ...int) []int
	Concat(ints ...[]int) []int
	Every(ints []int, predict func(int) bool) bool
	Some(ints []int, predict func(int) bool) bool
	Fill(ints []int, value int, startEndIndex ...int)
	Join(ints []int, sep ...string) string
	Map(ints []int, m func(int) int) []int
	Pop(ints []int) (int, bool, []int)
	Push(ints []int, es ...int) (int, []int)
	Reduce(ints []int, accum func(int, int) int, init ...int) int
	ReduceRight(ints []int, accum func(int, int) int, init ...int) int
	Reverse(ints []int) []int
	Shift(ints []int) (int, bool, []int)
	Unshift(ints []int, es ...int) (int, []int)
	Slice(ints []int, startEndIndex ...int) []int
	Sort(ints []int, less ...func(int, int) bool)
	Shuffle(ints []int)
	String(ints []int) string
	Drop(ints []int, n ...int) []int
	Head(ints []int) (int, bool)
	Initial(ints []int) []int
	Last(ints []int) (int, bool)
	Nth(ints []int, n ...int) (int, bool)
	Pull(ints []int, vs ...int) []int
	Difference(ints []int, vs ...[]int) []int
	Intersection(vs ...[]int) []int
	Without(ints []int, vs ...int) []int
	Remove(ints []int, predict func(int) bool) []int
	Tail(ints []int) []int
	Take(ints []int, n ...int) []int
	Union(ints ...[]int) []int
	Uniq(ints []int) []int
	Equal(s1, s2 []int) bool
}

// NewInt creates a new Int instance.
func NewInt() *Int {
	var p *Int
	return p
}

var _ IInt = NewInt()

// Filter creates a new slice with all elements that pass the predict implemented.
func (*Int) Filter(ints []int, predict func(int) bool) []int {
	var r []int

	for _, v := range ints {
		if ok := predict(v); ok {
			r = append(r, v)
		}
	}

	return r
}

// Find returns the value of the first element in the slice that satisfies the provided predict function.
// otherwise false is returned.
func (*Int) Find(ints []int, predict func(int) bool) (int, bool) {
	for _, v := range ints {
		if ok := predict(v); ok {
			return v, true
		}
	}

	return 0, false
}

// FindIndex returns the index of the first element in the slice that satisfies the provided predict function.
// otherwise -1 is returned.
func (*Int) FindIndex(ints []int, predict func(int) bool, fromIndex ...int) int {
	var start int
	n := len(ints)

	for _, id := range fromIndex {
		start = id
	}

	start = g.Min(start, n)
	if start < 0 {
		start = g.Max(n+start, 0)
	}

	for i := start; i < n; i++ {
		if ok := predict(ints[i]); ok {
			return i
		}
	}

	return -1
}

// FindLastIndex  is like FindIndex except that it iterates over elements from right to left.
func (*Int) FindLastIndex(ints []int, predict func(int) bool, fromIndex ...int) int {
	n := len(ints)
	start := n - 1

	for _, id := range fromIndex {
		start = id
	}

	start = g.Min(start, n)
	if start < 0 {
		start = g.Max(n+start, 0)
	}

	for i := start; i >= 0; i-- {
		if ok := predict(ints[i]); ok {
			return i
		}
	}

	return -1
}

// ForEach executes a provided function once for each slice element.
func (*Int) ForEach(ints []int, provided func(int)) {
	for _, v := range ints {
		provided(v)
	}
}

// Includes determines whether a slice includes a certain element, returning true or false.
func (*Int) Includes(ints []int, e int) bool {
	for _, v := range ints {
		if e == v {
			return true
		}
	}

	return false
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
func (*Int) IndexOf(ints []int, e int, fromIndex ...int) int {
	var start int
	n := len(ints)

	for _, id := range fromIndex {
		start = id
	}

	start = g.Min(start, n)
	if start < 0 {
		start = g.Max(n+start, 0)
	}

	for i := start; i < n; i++ {
		if e == ints[i] {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
func (*Int) LastIndexOf(ints []int, e int, fromIndex ...int) int {
	var start int
	n := len(ints)

	for _, id := range fromIndex {
		start = id
	}

	start = g.Min(start, n)
	if start < 0 {
		start = g.Max(n+start, 0)
	}

	for i := start; i >= 0; i-- {
		if e == ints[i] {
			return i
		}
	}

	return -1
}

// From creates a new, shallow-copied slice.
func (*Int) From(ints []int, mapFn ...func(int) int) []int {
	r := make([]int, len(ints))
	f := func(i int) int {
		v := i
		for _, handler := range mapFn {
			v = handler(v)
		}
		return v
	}

	for i, v := range ints {
		r[i] = f(v)
	}

	return r
}

// Of creates a new slice instance with a variable number of arguments.
func (*Int) Of(ints ...int) []int {
	r := make([]int, len(ints))

	for i, v := range ints {
		r[i] = v
	}
	return r
}

// Concat merges two or more slices.
func (*Int) Concat(ints ...[]int) []int {
	var r []int

	for _, v := range ints {
		r = append(r, v...)
	}

	return r
}

// Every tests whether all elements in the slice pass the predict implemented.
func (*Int) Every(ints []int, predict func(int) bool) bool {
	for _, v := range ints {
		if ok := predict(v); !ok {
			return false
		}
	}

	return true
}

// Some tests whether at least one element in the slice pass the predict implemented.
func (*Int) Some(ints []int, predict func(int) bool) bool {
	for _, v := range ints {
		if ok := predict(v); ok {
			return true
		}
	}

	return false
}

// Fill fills all the elements of a slice from a start index to an end index with a static value.
// the end index is not included.
func (*Int) Fill(ints []int, value int, startEndIndex ...int) {
	n := len(ints)
	start, end := 0, n

	switch len(startEndIndex) {
	case 1:
		start = startEndIndex[0]
	case 2:
		start, end = startEndIndex[0], startEndIndex[1]
	}

	start = g.Min(start, n)
	if start < 0 {
		start = g.Max(n+start, 0)
	}

	end = g.Min(end, n)
	if end < 0 {
		end = g.Max(n+end, 0)
	}

	for start < end {
		ints[start] = value
		start++
	}
}

// Join joins all elements of a slice into a string.
func (*Int) Join(ints []int, sep ...string) string {
	v := make([]string, len(ints))

	for i, t := range ints {
		v[i] = strconv.Itoa(t)
	}

	s := ","
	for _, se := range sep {
		s = se
	}
	return strings.Join(v, s)
}

// Map creates a new slice with the results of calling a provided function on every element in the calling slice.
func (*Int) Map(ints []int, m func(int) int) []int {
	r := make([]int, len(ints))

	for i, v := range ints {
		r[i] = m(v)
	}

	return r
}

// Pop removes the last element from a slice and returns that element.
func (*Int) Pop(ints []int) (int, bool, []int) {
	n := len(ints)
	if n == 0 {
		return 0, false, []int{}
	}
	return ints[n-1], true, ints[:n-1]
}

// Push adds one or more elements to the end of a slice and returns the new length of the slice.
func (*Int) Push(ints []int, es ...int) (int, []int) {
	m := len(ints)
	n := m + len(es)
	r := make([]int, n)
	copy(r, ints)
	for i, v := range es {
		r[i+m] = v
	}
	return n, r
}

// Reduce applies a function against each element in the slice to reduce it to a single value.
func (*Int) Reduce(ints []int, accum func(int, int) int, init ...int) int {
	var iv int
	for _, v := range init {
		iv = v
	}
	for _, v := range ints {
		iv = accum(iv, v)
	}
	return iv
}

// ReduceRight applies a function against each value of the slice (from right-to-left) to reduce it to a single value.
func (*Int) ReduceRight(ints []int, accum func(int, int) int, init ...int) int {
	n := len(ints) - 1
	var iv int
	for _, v := range init {
		iv = v
	}
	for i := range ints {
		iv = accum(iv, ints[n-i])
	}
	return iv
}

// Reverse reverses a slice.
func (*Int) Reverse(ints []int) []int {
	len := len(ints) - 1
	for i := 0; i < len; i++ {
		ints[i], ints[len] = ints[len], ints[i]
		len--
	}
	return ints
}

// Shift removes the first element from a slice and returns that element.
func (*Int) Shift(ints []int) (int, bool, []int) {
	n := len(ints)
	if n == 0 {
		return 0, false, []int{}
	}
	return ints[0], true, ints[1:]
}

// Unshift adds one or more elements to the beginning of a slice and returns the new length of the slice.
func (*Int) Unshift(ints []int, es ...int) (int, []int) {
	m := len(ints)
	n := len(es)
	r := make([]int, n+m)

	for i, v := range ints {
		r[i+n] = v
	}

	for i, v := range es {
		r[i] = v
	}

	return n + m, r
}

// Slice returns a portion of a slice selected from start to end (end not included).
// start and end indexes are optional, default to zero and length of the slice.
func (*Int) Slice(ints []int, startEndIndex ...int) []int {
	n := len(ints)
	start, end := 0, n

	switch len(startEndIndex) {
	case 1:
		start = startEndIndex[0]
	case 2:
		start, end = startEndIndex[0], startEndIndex[1]
	}

	start = g.Min(start, n)
	if start < 0 {
		start = g.Max(n+start, 0)
	}

	end = g.Min(end, n)
	if end < 0 {
		end = g.Max(n+end, 0)
	}

	r := []int{}

	for start < end {
		r = append(r, ints[start])
		start++
	}

	return r
}

// Sort sorts elements.
func (*Int) Sort(ints []int, less ...func(int, int) bool) {
	swap := func(i, j int) { ints[i], ints[j] = ints[j], ints[i] }
	comp := func(i, j int) bool { return ints[i] < ints[j] }

	if len(less) > 0 {
		comp = less[0]
	}

	sort.Sort(g.Sorter{N: len(ints), SwapFunc: swap, LessFunc: comp})
}

// Shuffle randomizes the order of elements.
func (*Int) Shuffle(ints []int) {
	swap := func(i, j int) { ints[i], ints[j] = ints[j], ints[i] }
	rand.Shuffle(len(ints), swap)
}

// String returns a string representing the specified slice.
func (i *Int) String(ints []int) string {
	return i.Join(ints)
}

// Drop creates a slice with n elements dropped from the beginning.
func (*Int) Drop(ints []int, n ...int) []int {
	start := 1
	lenin := len(ints)

	if len(n) > 0 && n[0] >= 0 {
		start = n[0]
	}

	if start >= lenin {
		return []int{}
	}
	r := make([]int, lenin-start)
	copy(r, ints[start:])
	return r
}

// Head gets the first element of slice.
func (*Int) Head(ints []int) (int, bool) {
	if len(ints) > 0 {
		return ints[0], true
	}

	return 0, false
}

// Initial gets all but the last element of slice.
func (*Int) Initial(ints []int) []int {
	n := len(ints)
	if n > 1 {
		return ints[:n-1]
	}

	return []int{}
}

// Last gets the last element of slice.
func (*Int) Last(ints []int) (int, bool) {
	n := len(ints)
	if n > 0 {
		return ints[n-1], true
	}

	return 0, false
}

// Nth gets the element at index n of slice.
// if n is negative, the nth element from the end is returned.
func (*Int) Nth(ints []int, n ...int) (int, bool) {
	var start int
	lenin := len(ints)

	for _, i := range n {
		start = i
	}

	if start < 0 {
		start = lenin + start
	}

	if start >= 0 && start < lenin {
		return ints[start], true
	}

	return 0, false
}

// Pull removes all given values from slice.
func (i *Int) Pull(ints []int, vs ...int) []int {
	var r []int
	for _, v := range ints {
		if ok := i.Includes(vs, v); !ok {
			r = append(r, v)
		}
	}
	return r
}

// Difference creates a slice of values not included in the other given slices.
func (i *Int) Difference(ints []int, vs ...[]int) []int {
	var r []int
	var counter int
	n := len(vs)

	for _, t := range ints {
		counter = 0
		for _, v := range vs {
			if ok := i.Includes(v, t); !ok {
				counter++
			}
		}
		if counter == n {
			r = append(r, t)
		}
	}
	return r
}

// Intersection creates a slice of unique values that are included in all given slices.
func (i *Int) Intersection(vs ...[]int) []int {
	var r []int
	n := len(vs)
	var counter int

	if n == 0 {
		return r
	}
	if n == 1 {
		return append(r, vs[0]...)
	}

	for _, t := range vs[0] {
		counter = 0
		for j := 1; j < n; j++ {
			if ok := i.Includes(vs[j], t); ok {
				counter++
			}
		}
		if counter == n-1 {
			r = append(r, t)
		}
	}

	return r
}

// Without creates a slice excluding all given values.
func (i *Int) Without(ints []int, vs ...int) []int {
	return i.Pull(ints, vs...)
}

// Remove removes all elements from slice that predicate returns truthy for.
func (*Int) Remove(ints []int, predict func(int) bool) []int {
	var r []int

	for _, v := range ints {
		if ok := predict(v); !ok {
			r = append(r, v)
		}
	}

	return r
}

// Tail gets all but the first element of slice.
func (*Int) Tail(ints []int) []int {
	n := len(ints)
	if n > 1 {
		return ints[1:]
	}

	return []int{}
}

// Take creates a slice with n elements taken from the beginning.
func (*Int) Take(ints []int, n ...int) []int {
	start := 1
	lenin := len(ints)

	if len(n) > 0 && n[0] >= 0 {
		start = n[0]
	}

	if start >= lenin {
		return ints
	}

	return ints[:start]
}

// Union creates a slice of unique values, in order, from all given slices.
func (i *Int) Union(ints ...[]int) []int {
	var r []int
	for _, vs := range ints {
		for _, v := range vs {
			if ok := i.Includes(r, v); !ok {
				r = append(r, v)
			}
		}
	}
	return r
}

// Uniq creates a duplicate-free version of a slice.
func (i *Int) Uniq(ints []int) []int {
	var r []int
	for _, v := range ints {
		if ok := i.Includes(r, v); !ok {
			r = append(r, v)
		}
	}
	return r
}

// Equal performs a deep comparison between two slice values.
func (*Int) Equal(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}
