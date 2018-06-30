package goji

import (
	"math/rand"
	"sort"
	"strings"
)

// Float64 for float64 type slice.
type Float64 struct{}

// Float64Helper is the Float64 interface.
type Float64Helper interface {
	Min(x, y float64) float64
	Max(x, y float64) float64
	Filter(ints []float64, predict func(float64) bool) []float64
	Find(ints []float64, predict func(float64) bool) (float64, bool)
	FindIndex(ints []float64, predict func(float64) bool, fromIndex ...int) int
	FindLastIndex(ints []float64, predict func(float64) bool, fromIndex ...int) int
	ForEach(ints []float64, provided func(float64))
	Includes(ints []float64, e float64) bool
	IndexOf(ints []float64, e float64, fromIndex ...int) int
	LastIndexOf(ints []float64, e float64, fromIndex ...int) int
	From(ints []float64, mapFn ...func(float64) float64) []float64
	Of(ints ...float64) []float64
	Concat(ints ...[]float64) []float64
	Every(ints []float64, predict func(float64) bool) bool
	Some(ints []float64, predict func(float64) bool) bool
	Fill(ints []float64, value float64, startEndIndex ...int)
	Join(ints []float64, sep ...string) string
	Map(ints []float64, m func(float64) float64) []float64
	Pop(ints []float64) (float64, bool, []float64)
	Push(ints []float64, es ...float64) (int, []float64)
	Reduce(ints []float64, accum func(float64, float64) float64, init ...float64) float64
	ReduceRight(ints []float64, accum func(float64, float64) float64, init ...float64) float64
	Reverse(ints []float64) []float64
	Shift(ints []float64) (float64, bool, []float64)
	Unshift(ints []float64, es ...float64) (int, []float64)
	Slice(ints []float64, startEndIndex ...int) []float64
	Sort(ints []float64, less ...func(int, int) bool)
	Shuffle(ints []float64)
	String(ints []float64) string
	Drop(ints []float64, n ...int) []float64
	Head(ints []float64) (float64, bool)
	Initial(ints []float64) []float64
	Last(ints []float64) (float64, bool)
	Nth(ints []float64, n ...int) (float64, bool)
	Pull(ints []float64, vs ...float64) []float64
	Difference(ints []float64, vs ...[]float64) []float64
	Intersection(vs ...[]float64) []float64
	Without(ints []float64, vs ...float64) []float64
	Remove(ints []float64, predict func(float64) bool) []float64
	Tail(ints []float64) []float64
	Take(ints []float64, n ...int) []float64
	Union(ints ...[]float64) []float64
	Uniq(ints []float64) []float64
	Equal(s1, s2 []float64) bool
}

// NewFloat64 creates a new Float64 instance.
func NewFloat64() *Float64 {
	return &Float64{}
}

var _ Float64Helper = NewFloat64()

// Min returns the smaller of x or y.
func (*Float64) Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

// Max returns the larger of x or y.
func (*Float64) Max(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

// Filter creates a new slice with all elements that pass the predict implemented.
func (*Float64) Filter(ints []float64, predict func(float64) bool) []float64 {
	var r []float64

	for _, v := range ints {
		if ok := predict(v); ok {
			r = append(r, v)
		}
	}

	return r
}

// Find returns the value of the first element in the slice that satisfies the provided predict function.
// otherwise false is returned.
func (*Float64) Find(ints []float64, predict func(float64) bool) (float64, bool) {
	for _, v := range ints {
		if ok := predict(v); ok {
			return v, true
		}
	}

	return 0, false
}

// FindIndex returns the index of the first element in the slice that satisfies the provided predict function.
// otherwise -1 is returned.
func (*Float64) FindIndex(ints []float64, predict func(float64) bool, fromIndex ...int) int {
	var start int
	var n = len(ints)

	for _, id := range fromIndex {
		start = id
	}

	start = Min(start, n)
	if start < 0 {
		start = Max(n+start, 0)
	}

	for i := start; i < n; i++ {
		if ok := predict(ints[i]); ok {
			return i
		}
	}

	return -1
}

// FindLastIndex  is like FindIndex except that it iterates over elements from right to left.
func (*Float64) FindLastIndex(ints []float64, predict func(float64) bool, fromIndex ...int) int {
	var n = len(ints)
	var start = n - 1

	for _, id := range fromIndex {
		start = id
	}

	start = Min(start, n)
	if start < 0 {
		start = Max(n+start, 0)
	}

	for i := start; i >= 0; i-- {
		if ok := predict(ints[i]); ok {
			return i
		}
	}

	return -1
}

// ForEach executes a provided function once for each slice element.
func (*Float64) ForEach(ints []float64, provided func(float64)) {
	for _, v := range ints {
		provided(v)
	}
}

// Includes determines whether a slice includes a certain element, returning true or false.
func (*Float64) Includes(ints []float64, e float64) bool {
	for _, v := range ints {
		if e == v {
			return true
		}
	}

	return false
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
func (*Float64) IndexOf(ints []float64, e float64, fromIndex ...int) int {
	var start int
	var n = len(ints)

	for _, id := range fromIndex {
		start = id
	}

	start = Min(start, n)
	if start < 0 {
		start = Max(n+start, 0)
	}

	for i := start; i < n; i++ {
		if e == ints[i] {
			return i
		}
	}

	return -1
}

// LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.
func (*Float64) LastIndexOf(ints []float64, e float64, fromIndex ...int) int {
	var start int
	var n = len(ints)

	for _, id := range fromIndex {
		start = id
	}

	start = Min(start, n)
	if start < 0 {
		start = Max(n+start, 0)
	}

	for i := start; i >= 0; i-- {
		if e == ints[i] {
			return i
		}
	}

	return -1
}

// From creates a new, shallow-copied slice.
func (*Float64) From(ints []float64, mapFn ...func(float64) float64) []float64 {
	r := make([]float64, len(ints))
	f := func(i float64) float64 {
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
func (*Float64) Of(ints ...float64) []float64 {
	r := make([]float64, len(ints))

	for i, v := range ints {
		r[i] = v
	}
	return r
}

// Concat merges two or more slices.
func (*Float64) Concat(ints ...[]float64) []float64 {
	var r []float64

	for _, v := range ints {
		r = append(r, v...)
	}

	return r
}

// Every tests whether all elements in the slice pass the predict implemented.
func (*Float64) Every(ints []float64, predict func(float64) bool) bool {
	for _, v := range ints {
		if ok := predict(v); !ok {
			return false
		}
	}

	return true
}

// Some tests whether at least one element in the slice pass the predict implemented.
func (*Float64) Some(ints []float64, predict func(float64) bool) bool {
	for _, v := range ints {
		if ok := predict(v); ok {
			return true
		}
	}

	return false
}

// Fill fills all the elements of a slice from a start index to an end index with a static value.
// the end index is not included.
func (*Float64) Fill(ints []float64, value float64, startEndIndex ...int) {
	var n = len(ints)
	var start, end int = 0, n

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
		ints[start] = value
		start++
	}
}

// Join joins all elements of a slice into a string.
func (*Float64) Join(ints []float64, sep ...string) string {
	v := make([]string, len(ints))

	for i, t := range ints {
		v[i] = FloatToString(float64(t))
	}

	s := ","
	for _, se := range sep {
		s = se
	}
	return strings.Join(v, s)
}

// JoinWithPrec joins all elements of a slice into a string with precision.
func (*Float64) JoinWithPrec(ints []float64, prec int, sep ...string) string {
	v := make([]string, len(ints))

	for i, t := range ints {
		v[i] = FloatToStringWithPrec(float64(t), prec)
	}

	s := ","
	for _, se := range sep {
		s = se
	}
	return strings.Join(v, s)
}

// Map creates a new slice with the results of calling a provided function on every element in the calling slice.
func (*Float64) Map(ints []float64, m func(float64) float64) []float64 {
	r := make([]float64, len(ints))

	for i, v := range ints {
		r[i] = m(v)
	}

	return r
}

// Pop removes the last element from a slice and returns that element.
func (*Float64) Pop(ints []float64) (float64, bool, []float64) {
	n := len(ints)
	if n == 0 {
		return 0, false, []float64{}
	}
	return ints[n-1], true, ints[:n-1]
}

// Push adds one or more elements to the end of a slice and returns the new length of the slice.
func (*Float64) Push(ints []float64, es ...float64) (int, []float64) {
	var m = len(ints)
	n := m + len(es)
	r := make([]float64, n)
	copy(r, ints)
	for i, v := range es {
		r[i+m] = v
	}
	return n, r
}

// Reduce applies a function against each element in the slice to reduce it to a single value.
func (*Float64) Reduce(ints []float64, accum func(float64, float64) float64, init ...float64) float64 {
	var iv float64
	for _, v := range init {
		iv = v
	}
	for _, v := range ints {
		iv = accum(iv, v)
	}
	return iv
}

// ReduceRight applies a function against each value of the slice (from right-to-left) to reduce it to a single value.
func (*Float64) ReduceRight(ints []float64, accum func(float64, float64) float64, init ...float64) float64 {
	n := len(ints) - 1
	var iv float64
	for _, v := range init {
		iv = v
	}
	for i := range ints {
		iv = accum(iv, ints[n-i])
	}
	return iv
}

// Reverse reverses a slice.
func (*Float64) Reverse(ints []float64) []float64 {
	len := len(ints) - 1
	for i := 0; i < len; i++ {
		ints[i], ints[len] = ints[len], ints[i]
		len--
	}
	return ints
}

// Shift removes the first element from a slice and returns that element.
func (*Float64) Shift(ints []float64) (float64, bool, []float64) {
	n := len(ints)
	if n == 0 {
		return 0, false, []float64{}
	}
	return ints[0], true, ints[1:]
}

// Unshift adds one or more elements to the beginning of a slice and returns the new length of the slice.
func (*Float64) Unshift(ints []float64, es ...float64) (int, []float64) {
	m := len(ints)
	n := len(es)
	r := make([]float64, n+m)

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
func (*Float64) Slice(ints []float64, startEndIndex ...int) []float64 {
	n := len(ints)
	var start, end int = 0, n

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

	r := []float64{}

	for start < end {
		r = append(r, ints[start])
		start++
	}

	return r
}

// Sort sorts elements.
func (*Float64) Sort(ints []float64, less ...func(int, int) bool) {
	swap := func(i, j int) { ints[i], ints[j] = ints[j], ints[i] }
	comp := func(i, j int) bool { return ints[i] < ints[j] }

	if len(less) > 0 {
		comp = less[0]
	}

	sort.Sort(sorter{len: len(ints), swap: swap, less: comp})
}

// Shuffle randomizes the order of elements.
func (*Float64) Shuffle(ints []float64) {
	swap := func(i, j int) { ints[i], ints[j] = ints[j], ints[i] }
	rand.Shuffle(len(ints), swap)
}

// String returns a string representing the specified slice.
func (i *Float64) String(ints []float64) string {
	return i.Join(ints)
}

// Drop creates a slice with n elements dropped from the beginning.
func (*Float64) Drop(ints []float64, n ...int) []float64 {
	var start = 1
	lenin := len(ints)

	if len(n) > 0 && n[0] >= 0 {
		start = n[0]
	}

	if start >= lenin {
		return []float64{}
	}
	r := make([]float64, lenin-start)
	copy(r, ints[start:])
	return r
}

// Head gets the first element of slice.
func (*Float64) Head(ints []float64) (float64, bool) {
	if len(ints) > 0 {
		return ints[0], true
	}

	return 0, false
}

// Initial gets all but the last element of slice.
func (*Float64) Initial(ints []float64) []float64 {
	n := len(ints)
	if n > 1 {
		return ints[:n-1]
	}

	return []float64{}
}

// Last gets the last element of slice.
func (*Float64) Last(ints []float64) (float64, bool) {
	n := len(ints)
	if n > 0 {
		return ints[n-1], true
	}

	return 0, false
}

// Nth gets the element at index n of slice.
// if n is negative, the nth element from the end is returned.
func (*Float64) Nth(ints []float64, n ...int) (float64, bool) {
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
func (i *Float64) Pull(ints []float64, vs ...float64) []float64 {
	var r []float64
	for _, v := range ints {
		if ok := i.Includes(vs, v); !ok {
			r = append(r, v)
		}
	}
	return r
}

// Difference creates a slice of values not included in the other given slices.
func (i *Float64) Difference(ints []float64, vs ...[]float64) []float64 {
	var r []float64
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
func (i *Float64) Intersection(vs ...[]float64) []float64 {
	var r []float64
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
func (i *Float64) Without(ints []float64, vs ...float64) []float64 {
	return i.Pull(ints, vs...)
}

// Remove removes all elements from slice that predicate returns truthy for.
func (*Float64) Remove(ints []float64, predict func(float64) bool) []float64 {
	var r []float64

	for _, v := range ints {
		if ok := predict(v); !ok {
			r = append(r, v)
		}
	}

	return r
}

// Tail gets all but the first element of slice.
func (*Float64) Tail(ints []float64) []float64 {
	n := len(ints)
	if n > 1 {
		return ints[1:]
	}

	return []float64{}
}

// Take creates a slice with n elements taken from the beginning.
func (*Float64) Take(ints []float64, n ...int) []float64 {
	var start = 1
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
func (i *Float64) Union(ints ...[]float64) []float64 {
	var r []float64
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
func (i *Float64) Uniq(ints []float64) []float64 {
	var r []float64
	for _, v := range ints {
		if ok := i.Includes(r, v); !ok {
			r = append(r, v)
		}
	}
	return r
}

// Equal performs a deep comparison between two slice values.
func (*Float64) Equal(s1, s2 []float64) bool {
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
