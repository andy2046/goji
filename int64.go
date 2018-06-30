package goji

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

// Int64 for int64 type slice.
type Int64 struct{}

// Int64Helper is the Int64 interface.
type Int64Helper interface {
	Min(x, y int64) int64
	Max(x, y int64) int64
	Filter(ints []int64, predict func(int64) bool) []int64
	Find(ints []int64, predict func(int64) bool) (int64, bool)
	FindIndex(ints []int64, predict func(int64) bool, fromIndex ...int64) int64
	FindLastIndex(ints []int64, predict func(int64) bool, fromIndex ...int64) int64
	ForEach(ints []int64, provided func(int64))
	Includes(ints []int64, e int64) bool
	IndexOf(ints []int64, e int64, fromIndex ...int64) int64
	LastIndexOf(ints []int64, e int64, fromIndex ...int64) int64
	From(ints []int64, mapFn ...func(int64) int64) []int64
	Of(ints ...int64) []int64
	Concat(ints ...[]int64) []int64
	Every(ints []int64, predict func(int64) bool) bool
	Some(ints []int64, predict func(int64) bool) bool
	Fill(ints []int64, value int64, startEndIndex ...int64)
	Join(ints []int64, sep ...string) string
	Map(ints []int64, m func(int64) int64) []int64
	Pop(ints []int64) (int64, bool, []int64)
	Push(ints []int64, es ...int64) (int64, []int64)
	Reduce(ints []int64, accum func(int64, int64) int64, init ...int64) int64
	ReduceRight(ints []int64, accum func(int64, int64) int64, init ...int64) int64
	Reverse(ints []int64) []int64
	Shift(ints []int64) (int64, bool, []int64)
	Unshift(ints []int64, es ...int64) (int64, []int64)
	Slice(ints []int64, startEndIndex ...int64) []int64
	Sort(ints []int64, less ...func(int, int) bool)
	Shuffle(ints []int64)
	String(ints []int64) string
	Drop(ints []int64, n ...int64) []int64
	Head(ints []int64) (int64, bool)
	Initial(ints []int64) []int64
	Last(ints []int64) (int64, bool)
	Nth(ints []int64, n ...int64) (int64, bool)
	Pull(ints []int64, vs ...int64) []int64
	Difference(ints []int64, vs ...[]int64) []int64
	Intersection(vs ...[]int64) []int64
	Without(ints []int64, vs ...int64) []int64
	Remove(ints []int64, predict func(int64) bool) []int64
	Tail(ints []int64) []int64
	Take(ints []int64, n ...int64) []int64
	Union(ints ...[]int64) []int64
	Uniq(ints []int64) []int64
	Equal(s1, s2 []int64) bool
}

// NewInt64 creates a new Int64 instance.
func NewInt64() *Int64 {
	return &Int64{}
}

var _ Int64Helper = NewInt64()

// Min returns the smaller of x or y.
func (*Int64) Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

// Max returns the larger of x or y.
func (*Int64) Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

// Filter creates a new slice with all elements that pass the predict implemented.
func (*Int64) Filter(ints []int64, predict func(int64) bool) []int64 {
	var r []int64

	for _, v := range ints {
		if ok := predict(v); ok {
			r = append(r, v)
		}
	}

	return r
}

// Find returns the value of the first element in the slice that satisfies the provided predict function.
// otherwise false is returned.
func (*Int64) Find(ints []int64, predict func(int64) bool) (int64, bool) {
	for _, v := range ints {
		if ok := predict(v); ok {
			return v, true
		}
	}

	return 0, false
}

// FindIndex returns the index of the first element in the slice that satisfies the provided predict function.
// otherwise -1 is returned.
func (i *Int64) FindIndex(ints []int64, predict func(int64) bool, fromIndex ...int64) int64 {
	var start int64
	var n = int64(len(ints))

	for _, id := range fromIndex {
		start = id
	}

	start = i.Min(start, n)
	if start < 0 {
		start = i.Max(n+start, 0)
	}

	for i := start; i < n; i++ {
		if ok := predict(ints[i]); ok {
			return i
		}
	}

	return -1
}

// FindLastIndex  is like FindIndex except that it iterates over elements from right to left.
func (i *Int64) FindLastIndex(ints []int64, predict func(int64) bool, fromIndex ...int64) int64 {
	var n = int64(len(ints))
	var start = n - 1

	for _, id := range fromIndex {
		start = id
	}

	start = i.Min(start, n)
	if start < 0 {
		start = i.Max(n+start, 0)
	}

	for i := start; i >= 0; i-- {
		if ok := predict(ints[i]); ok {
			return i
		}
	}

	return -1
}

// ForEach executes a provided function once for each slice element.
func (*Int64) ForEach(ints []int64, provided func(int64)) {
	for _, v := range ints {
		provided(v)
	}
}

// Includes determines whether a slice includes a certain element, returning true or false.
func (*Int64) Includes(ints []int64, e int64) bool {
	for _, v := range ints {
		if e == v {
			return true
		}
	}

	return false
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
func (i *Int64) IndexOf(ints []int64, e int64, fromIndex ...int64) int64 {
	var start int64
	var n = int64(len(ints))

	for _, id := range fromIndex {
		start = id
	}

	start = i.Min(start, n)
	if start < 0 {
		start = i.Max(n+start, 0)
	}

	for i := start; i < n; i++ {
		if e == ints[i] {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
func (i *Int64) LastIndexOf(ints []int64, e int64, fromIndex ...int64) int64 {
	var start int64
	var n = int64(len(ints))

	for _, id := range fromIndex {
		start = id
	}

	start = i.Min(start, n)
	if start < 0 {
		start = i.Max(n+start, 0)
	}

	for i := start; i >= 0; i-- {
		if e == ints[i] {
			return i
		}
	}

	return -1
}

// From creates a new, shallow-copied slice.
func (*Int64) From(ints []int64, mapFn ...func(int64) int64) []int64 {
	r := make([]int64, len(ints))
	f := func(i int64) int64 {
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
func (*Int64) Of(ints ...int64) []int64 {
	r := make([]int64, len(ints))

	for i, v := range ints {
		r[i] = v
	}
	return r
}

// Concat merges two or more slices.
func (*Int64) Concat(ints ...[]int64) []int64 {
	var r []int64

	for _, v := range ints {
		r = append(r, v...)
	}

	return r
}

// Every tests whether all elements in the slice pass the predict implemented.
func (*Int64) Every(ints []int64, predict func(int64) bool) bool {
	for _, v := range ints {
		if ok := predict(v); !ok {
			return false
		}
	}

	return true
}

// Some tests whether at least one element in the slice pass the predict implemented.
func (*Int64) Some(ints []int64, predict func(int64) bool) bool {
	for _, v := range ints {
		if ok := predict(v); ok {
			return true
		}
	}

	return false
}

// Fill fills all the elements of a slice from a start index to an end index with a static value.
// the end index is not included.
func (i *Int64) Fill(ints []int64, value int64, startEndIndex ...int64) {
	var n = int64(len(ints))
	var start, end int64 = 0, n

	switch len(startEndIndex) {
	case 1:
		start = startEndIndex[0]
	case 2:
		start, end = startEndIndex[0], startEndIndex[1]
	}

	start = i.Min(start, n)
	if start < 0 {
		start = i.Max(n+start, 0)
	}

	end = i.Min(end, n)
	if end < 0 {
		end = i.Max(n+end, 0)
	}

	for start < end {
		ints[start] = value
		start++
	}
}

// Join joins all elements of a slice into a string.
func (*Int64) Join(ints []int64, sep ...string) string {
	v := make([]string, len(ints))

	for i, t := range ints {
		v[i] = strconv.FormatInt(int64(t), 10)
	}

	s := ","
	for _, se := range sep {
		s = se
	}
	return strings.Join(v, s)
}

// Map creates a new slice with the results of calling a provided function on every element in the calling slice.
func (*Int64) Map(ints []int64, m func(int64) int64) []int64 {
	r := make([]int64, len(ints))

	for i, v := range ints {
		r[i] = m(v)
	}

	return r
}

// Pop removes the last element from a slice and returns that element.
func (*Int64) Pop(ints []int64) (int64, bool, []int64) {
	n := len(ints)
	if n == 0 {
		return 0, false, []int64{}
	}
	return ints[n-1], true, ints[:n-1]
}

// Push adds one or more elements to the end of a slice and returns the new length of the slice.
func (*Int64) Push(ints []int64, es ...int64) (int64, []int64) {
	var m = int64(len(ints))
	n := m + int64(len(es))
	r := make([]int64, n)
	copy(r, ints)
	for i, v := range es {
		r[int64(i)+m] = v
	}
	return n, r
}

// Reduce applies a function against each element in the slice to reduce it to a single value.
func (*Int64) Reduce(ints []int64, accum func(int64, int64) int64, init ...int64) int64 {
	var iv int64
	for _, v := range init {
		iv = v
	}
	for _, v := range ints {
		iv = accum(iv, v)
	}
	return iv
}

// ReduceRight applies a function against each value of the slice (from right-to-left) to reduce it to a single value.
func (*Int64) ReduceRight(ints []int64, accum func(int64, int64) int64, init ...int64) int64 {
	n := int64(len(ints)) - 1
	var iv int64
	for _, v := range init {
		iv = v
	}
	for i := range ints {
		iv = accum(iv, ints[n-int64(i)])
	}
	return iv
}

// Reverse reverses a slice.
func (*Int64) Reverse(ints []int64) []int64 {
	len := len(ints) - 1
	for i := 0; i < len; i++ {
		ints[i], ints[len] = ints[len], ints[i]
		len--
	}
	return ints
}

// Shift removes the first element from a slice and returns that element.
func (*Int64) Shift(ints []int64) (int64, bool, []int64) {
	n := len(ints)
	if n == 0 {
		return 0, false, []int64{}
	}
	return ints[0], true, ints[1:]
}

// Unshift adds one or more elements to the beginning of a slice and returns the new length of the slice.
func (*Int64) Unshift(ints []int64, es ...int64) (int64, []int64) {
	m := int64(len(ints))
	n := int64(len(es))
	r := make([]int64, n+m)

	for i, v := range ints {
		r[int64(i)+n] = v
	}

	for i, v := range es {
		r[i] = v
	}

	return n + m, r
}

// Slice returns a portion of a slice selected from start to end (end not included).
// start and end indexes are optional, default to zero and length of the slice.
func (i *Int64) Slice(ints []int64, startEndIndex ...int64) []int64 {
	n := int64(len(ints))
	var start, end int64 = 0, n

	switch len(startEndIndex) {
	case 1:
		start = startEndIndex[0]
	case 2:
		start, end = startEndIndex[0], startEndIndex[1]
	}

	start = i.Min(start, n)
	if start < 0 {
		start = i.Max(n+start, 0)
	}

	end = i.Min(end, n)
	if end < 0 {
		end = i.Max(n+end, 0)
	}

	r := []int64{}

	for start < end {
		r = append(r, ints[start])
		start++
	}

	return r
}

// Sort sorts elements.
func (*Int64) Sort(ints []int64, less ...func(int, int) bool) {
	swap := func(i, j int) { ints[i], ints[j] = ints[j], ints[i] }
	comp := func(i, j int) bool { return ints[i] < ints[j] }

	if len(less) > 0 {
		comp = less[0]
	}

	sort.Sort(sorter{len: len(ints), swap: swap, less: comp})
}

// Shuffle randomizes the order of elements.
func (*Int64) Shuffle(ints []int64) {
	swap := func(i, j int) { ints[i], ints[j] = ints[j], ints[i] }
	rand.Shuffle(len(ints), swap)
}

// String returns a string representing the specified slice.
func (i *Int64) String(ints []int64) string {
	return i.Join(ints)
}

// Drop creates a slice with n elements dropped from the beginning.
func (*Int64) Drop(ints []int64, n ...int64) []int64 {
	var start int64 = 1
	lenin := int64(len(ints))

	if len(n) > 0 && n[0] >= 0 {
		start = n[0]
	}

	if start >= lenin {
		return []int64{}
	}
	r := make([]int64, lenin-start)
	copy(r, ints[start:])
	return r
}

// Head gets the first element of slice.
func (*Int64) Head(ints []int64) (int64, bool) {
	if len(ints) > 0 {
		return ints[0], true
	}

	return 0, false
}

// Initial gets all but the last element of slice.
func (*Int64) Initial(ints []int64) []int64 {
	n := len(ints)
	if n > 1 {
		return ints[:n-1]
	}

	return []int64{}
}

// Last gets the last element of slice.
func (*Int64) Last(ints []int64) (int64, bool) {
	n := len(ints)
	if n > 0 {
		return ints[n-1], true
	}

	return 0, false
}

// Nth gets the element at index n of slice.
// if n is negative, the nth element from the end is returned.
func (*Int64) Nth(ints []int64, n ...int64) (int64, bool) {
	var start int64
	lenin := int64(len(ints))

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
func (i *Int64) Pull(ints []int64, vs ...int64) []int64 {
	var r []int64
	for _, v := range ints {
		if ok := i.Includes(vs, v); !ok {
			r = append(r, v)
		}
	}
	return r
}

// Difference creates a slice of values not included in the other given slices.
func (i *Int64) Difference(ints []int64, vs ...[]int64) []int64 {
	var r []int64
	var counter int64
	n := int64(len(vs))

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
func (i *Int64) Intersection(vs ...[]int64) []int64 {
	var r []int64
	n := int64(len(vs))
	var counter int64

	if n == 0 {
		return r
	}
	if n == 1 {
		return append(r, vs[0]...)
	}

	for _, t := range vs[0] {
		counter = 0
		for j := int64(1); j < n; j++ {
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
func (i *Int64) Without(ints []int64, vs ...int64) []int64 {
	return i.Pull(ints, vs...)
}

// Remove removes all elements from slice that predicate returns truthy for.
func (*Int64) Remove(ints []int64, predict func(int64) bool) []int64 {
	var r []int64

	for _, v := range ints {
		if ok := predict(v); !ok {
			r = append(r, v)
		}
	}

	return r
}

// Tail gets all but the first element of slice.
func (*Int64) Tail(ints []int64) []int64 {
	n := len(ints)
	if n > 1 {
		return ints[1:]
	}

	return []int64{}
}

// Take creates a slice with n elements taken from the beginning.
func (*Int64) Take(ints []int64, n ...int64) []int64 {
	var start int64 = 1
	lenin := int64(len(ints))

	if len(n) > 0 && n[0] >= 0 {
		start = n[0]
	}

	if start >= lenin {
		return ints
	}

	return ints[:start]
}

// Union creates a slice of unique values, in order, from all given slices.
func (i *Int64) Union(ints ...[]int64) []int64 {
	var r []int64
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
func (i *Int64) Uniq(ints []int64) []int64 {
	var r []int64
	for _, v := range ints {
		if ok := i.Includes(r, v); !ok {
			r = append(r, v)
		}
	}
	return r
}

// Equal performs a deep comparison between two slice values.
func (i *Int64) Equal(s1, s2 []int64) bool {
	if len(s1) != len(s2) {
		return false
	}
	if i.String(s1) != i.String(s2) {
		return false
	}
	return true
}
