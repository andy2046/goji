package goji

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

// Int32 for int32 type slice.
type Int32 struct{}

// Int32Helper is the Int32 interface.
type Int32Helper interface {
	Min(x, y int32) int32
	Max(x, y int32) int32
	Filter(ints []int32, predict func(int32) bool) []int32
	Find(ints []int32, predict func(int32) bool) (int32, bool)
	FindIndex(ints []int32, predict func(int32) bool, fromIndex ...int32) int32
	FindLastIndex(ints []int32, predict func(int32) bool, fromIndex ...int32) int32
	ForEach(ints []int32, provided func(int32))
	Includes(ints []int32, e int32) bool
	IndexOf(ints []int32, e int32, fromIndex ...int32) int32
	LastIndexOf(ints []int32, e int32, fromIndex ...int32) int32
	From(ints []int32, mapFn ...func(int32) int32) []int32
	Of(ints ...int32) []int32
	Concat(ints ...[]int32) []int32
	Every(ints []int32, predict func(int32) bool) bool
	Some(ints []int32, predict func(int32) bool) bool
	Fill(ints []int32, value int32, startEndIndex ...int32)
	Join(ints []int32, sep ...string) string
	Map(ints []int32, m func(int32) int32) []int32
	Pop(ints []int32) (int32, bool, []int32)
	Push(ints []int32, es ...int32) (int32, []int32)
	Reduce(ints []int32, accum func(int32, int32) int32, init ...int32) int32
	ReduceRight(ints []int32, accum func(int32, int32) int32, init ...int32) int32
	Reverse(ints []int32) []int32
	Shift(ints []int32) (int32, bool, []int32)
	Unshift(ints []int32, es ...int32) (int32, []int32)
	Slice(ints []int32, startEndIndex ...int32) []int32
	Sort(ints []int32, less ...func(int, int) bool)
	Shuffle(ints []int32)
	String(ints []int32) string
	Drop(ints []int32, n ...int32) []int32
	Head(ints []int32) (int32, bool)
	Initial(ints []int32) []int32
	Last(ints []int32) (int32, bool)
	Nth(ints []int32, n ...int32) (int32, bool)
	Pull(ints []int32, vs ...int32) []int32
	Difference(ints []int32, vs ...[]int32) []int32
	Intersection(vs ...[]int32) []int32
	Without(ints []int32, vs ...int32) []int32
	Remove(ints []int32, predict func(int32) bool) []int32
	Tail(ints []int32) []int32
	Take(ints []int32, n ...int32) []int32
	Union(ints ...[]int32) []int32
	Uniq(ints []int32) []int32
	Equal(s1, s2 []int32) bool
}

// NewInt32 creates a new Int32 instance.
func NewInt32() *Int32 {
	var p *Int32
	return p
}

var _ Int32Helper = NewInt32()

// Min returns the smaller of x or y.
func (*Int32) Min(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

// Max returns the larger of x or y.
func (*Int32) Max(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

// Filter creates a new slice with all elements that pass the predict implemented.
func (*Int32) Filter(ints []int32, predict func(int32) bool) []int32 {
	var r []int32

	for _, v := range ints {
		if ok := predict(v); ok {
			r = append(r, v)
		}
	}

	return r
}

// Find returns the value of the first element in the slice that satisfies the provided predict function.
// otherwise false is returned.
func (*Int32) Find(ints []int32, predict func(int32) bool) (int32, bool) {
	for _, v := range ints {
		if ok := predict(v); ok {
			return v, true
		}
	}

	return 0, false
}

// FindIndex returns the index of the first element in the slice that satisfies the provided predict function.
// otherwise -1 is returned.
func (i *Int32) FindIndex(ints []int32, predict func(int32) bool, fromIndex ...int32) int32 {
	var start int32
	var n = int32(len(ints))

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
func (i *Int32) FindLastIndex(ints []int32, predict func(int32) bool, fromIndex ...int32) int32 {
	var n = int32(len(ints))
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
func (*Int32) ForEach(ints []int32, provided func(int32)) {
	for _, v := range ints {
		provided(v)
	}
}

// Includes determines whether a slice includes a certain element, returning true or false.
func (*Int32) Includes(ints []int32, e int32) bool {
	for _, v := range ints {
		if e == v {
			return true
		}
	}

	return false
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
func (i *Int32) IndexOf(ints []int32, e int32, fromIndex ...int32) int32 {
	var start int32
	var n = int32(len(ints))

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
func (i *Int32) LastIndexOf(ints []int32, e int32, fromIndex ...int32) int32 {
	var start int32
	var n = int32(len(ints))

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
func (*Int32) From(ints []int32, mapFn ...func(int32) int32) []int32 {
	r := make([]int32, len(ints))
	f := func(i int32) int32 {
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
func (*Int32) Of(ints ...int32) []int32 {
	r := make([]int32, len(ints))

	for i, v := range ints {
		r[i] = v
	}
	return r
}

// Concat merges two or more slices.
func (*Int32) Concat(ints ...[]int32) []int32 {
	var r []int32

	for _, v := range ints {
		r = append(r, v...)
	}

	return r
}

// Every tests whether all elements in the slice pass the predict implemented.
func (*Int32) Every(ints []int32, predict func(int32) bool) bool {
	for _, v := range ints {
		if ok := predict(v); !ok {
			return false
		}
	}

	return true
}

// Some tests whether at least one element in the slice pass the predict implemented.
func (*Int32) Some(ints []int32, predict func(int32) bool) bool {
	for _, v := range ints {
		if ok := predict(v); ok {
			return true
		}
	}

	return false
}

// Fill fills all the elements of a slice from a start index to an end index with a static value.
// the end index is not included.
func (i *Int32) Fill(ints []int32, value int32, startEndIndex ...int32) {
	var n = int32(len(ints))
	var start, end int32 = 0, n

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
func (*Int32) Join(ints []int32, sep ...string) string {
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
func (*Int32) Map(ints []int32, m func(int32) int32) []int32 {
	r := make([]int32, len(ints))

	for i, v := range ints {
		r[i] = m(v)
	}

	return r
}

// Pop removes the last element from a slice and returns that element.
func (*Int32) Pop(ints []int32) (int32, bool, []int32) {
	n := len(ints)
	if n == 0 {
		return 0, false, []int32{}
	}
	return ints[n-1], true, ints[:n-1]
}

// Push adds one or more elements to the end of a slice and returns the new length of the slice.
func (*Int32) Push(ints []int32, es ...int32) (int32, []int32) {
	var m = int32(len(ints))
	n := m + int32(len(es))
	r := make([]int32, n)
	copy(r, ints)
	for i, v := range es {
		r[int32(i)+m] = v
	}
	return n, r
}

// Reduce applies a function against each element in the slice to reduce it to a single value.
func (*Int32) Reduce(ints []int32, accum func(int32, int32) int32, init ...int32) int32 {
	var iv int32
	for _, v := range init {
		iv = v
	}
	for _, v := range ints {
		iv = accum(iv, v)
	}
	return iv
}

// ReduceRight applies a function against each value of the slice (from right-to-left) to reduce it to a single value.
func (*Int32) ReduceRight(ints []int32, accum func(int32, int32) int32, init ...int32) int32 {
	n := int32(len(ints)) - 1
	var iv int32
	for _, v := range init {
		iv = v
	}
	for i := range ints {
		iv = accum(iv, ints[n-int32(i)])
	}
	return iv
}

// Reverse reverses a slice.
func (*Int32) Reverse(ints []int32) []int32 {
	len := len(ints) - 1
	for i := 0; i < len; i++ {
		ints[i], ints[len] = ints[len], ints[i]
		len--
	}
	return ints
}

// Shift removes the first element from a slice and returns that element.
func (*Int32) Shift(ints []int32) (int32, bool, []int32) {
	n := len(ints)
	if n == 0 {
		return 0, false, []int32{}
	}
	return ints[0], true, ints[1:]
}

// Unshift adds one or more elements to the beginning of a slice and returns the new length of the slice.
func (*Int32) Unshift(ints []int32, es ...int32) (int32, []int32) {
	m := int32(len(ints))
	n := int32(len(es))
	r := make([]int32, n+m)

	for i, v := range ints {
		r[int32(i)+n] = v
	}

	for i, v := range es {
		r[i] = v
	}

	return n + m, r
}

// Slice returns a portion of a slice selected from start to end (end not included).
// start and end indexes are optional, default to zero and length of the slice.
func (i *Int32) Slice(ints []int32, startEndIndex ...int32) []int32 {
	n := int32(len(ints))
	var start, end int32 = 0, n

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

	r := []int32{}

	for start < end {
		r = append(r, ints[start])
		start++
	}

	return r
}

// Sort sorts elements.
func (*Int32) Sort(ints []int32, less ...func(int, int) bool) {
	swap := func(i, j int) { ints[i], ints[j] = ints[j], ints[i] }
	comp := func(i, j int) bool { return ints[i] < ints[j] }

	if len(less) > 0 {
		comp = less[0]
	}

	sort.Sort(sorter{len: len(ints), swap: swap, less: comp})
}

// Shuffle randomizes the order of elements.
func (*Int32) Shuffle(ints []int32) {
	swap := func(i, j int) { ints[i], ints[j] = ints[j], ints[i] }
	rand.Shuffle(len(ints), swap)
}

// String returns a string representing the specified slice.
func (i *Int32) String(ints []int32) string {
	return i.Join(ints)
}

// Drop creates a slice with n elements dropped from the beginning.
func (*Int32) Drop(ints []int32, n ...int32) []int32 {
	var start int32 = 1
	lenin := int32(len(ints))

	if len(n) > 0 && n[0] >= 0 {
		start = n[0]
	}

	if start >= lenin {
		return []int32{}
	}
	r := make([]int32, lenin-start)
	copy(r, ints[start:])
	return r
}

// Head gets the first element of slice.
func (*Int32) Head(ints []int32) (int32, bool) {
	if len(ints) > 0 {
		return ints[0], true
	}

	return 0, false
}

// Initial gets all but the last element of slice.
func (*Int32) Initial(ints []int32) []int32 {
	n := len(ints)
	if n > 1 {
		return ints[:n-1]
	}

	return []int32{}
}

// Last gets the last element of slice.
func (*Int32) Last(ints []int32) (int32, bool) {
	n := len(ints)
	if n > 0 {
		return ints[n-1], true
	}

	return 0, false
}

// Nth gets the element at index n of slice.
// if n is negative, the nth element from the end is returned.
func (*Int32) Nth(ints []int32, n ...int32) (int32, bool) {
	var start int32
	lenin := int32(len(ints))

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
func (i *Int32) Pull(ints []int32, vs ...int32) []int32 {
	var r []int32
	for _, v := range ints {
		if ok := i.Includes(vs, v); !ok {
			r = append(r, v)
		}
	}
	return r
}

// Difference creates a slice of values not included in the other given slices.
func (i *Int32) Difference(ints []int32, vs ...[]int32) []int32 {
	var r []int32
	var counter int32
	n := int32(len(vs))

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
func (i *Int32) Intersection(vs ...[]int32) []int32 {
	var r []int32
	n := int32(len(vs))
	var counter int32

	if n == 0 {
		return r
	}
	if n == 1 {
		return append(r, vs[0]...)
	}

	for _, t := range vs[0] {
		counter = 0
		for j := int32(1); j < n; j++ {
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
func (i *Int32) Without(ints []int32, vs ...int32) []int32 {
	return i.Pull(ints, vs...)
}

// Remove removes all elements from slice that predicate returns truthy for.
func (*Int32) Remove(ints []int32, predict func(int32) bool) []int32 {
	var r []int32

	for _, v := range ints {
		if ok := predict(v); !ok {
			r = append(r, v)
		}
	}

	return r
}

// Tail gets all but the first element of slice.
func (*Int32) Tail(ints []int32) []int32 {
	n := len(ints)
	if n > 1 {
		return ints[1:]
	}

	return []int32{}
}

// Take creates a slice with n elements taken from the beginning.
func (*Int32) Take(ints []int32, n ...int32) []int32 {
	var start int32 = 1
	lenin := int32(len(ints))

	if len(n) > 0 && n[0] >= 0 {
		start = n[0]
	}

	if start >= lenin {
		return ints
	}

	return ints[:start]
}

// Union creates a slice of unique values, in order, from all given slices.
func (i *Int32) Union(ints ...[]int32) []int32 {
	var r []int32
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
func (i *Int32) Uniq(ints []int32) []int32 {
	var r []int32
	for _, v := range ints {
		if ok := i.Includes(r, v); !ok {
			r = append(r, v)
		}
	}
	return r
}

// Equal performs a deep comparison between two slice values.
func (*Int32) Equal(s1, s2 []int32) bool {
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
