package goji

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

// FloatToStringWithPrec converts float to string with precision.
func FloatToStringWithPrec(fv float64, prec int) string {
	return strconv.FormatFloat(fv, 'f', prec, 64)
}

// FloatToString converts float to string.
func FloatToString(fv float64) string {
	return strconv.FormatFloat(fv, 'f', 3, 64)
}

// Float32 for float32 type slice.
type Float32 struct{}

// Float32Helper is the Float32 interface.
type Float32Helper interface {
	Min(x, y float32) float32
	Max(x, y float32) float32
	Filter(ints []float32, predict func(float32) bool) []float32
	Find(ints []float32, predict func(float32) bool) (float32, bool)
	FindIndex(ints []float32, predict func(float32) bool, fromIndex ...int) int
	FindLastIndex(ints []float32, predict func(float32) bool, fromIndex ...int) int
	ForEach(ints []float32, provided func(float32))
	Includes(ints []float32, e float32) bool
	IndexOf(ints []float32, e float32, fromIndex ...int) int
	LastIndexOf(ints []float32, e float32, fromIndex ...int) int
	From(ints []float32, mapFn ...func(float32) float32) []float32
	Of(ints ...float32) []float32
	Concat(ints ...[]float32) []float32
	Every(ints []float32, predict func(float32) bool) bool
	Some(ints []float32, predict func(float32) bool) bool
	Fill(ints []float32, value float32, startEndIndex ...int)
	Join(ints []float32, sep ...string) string
	Map(ints []float32, m func(float32) float32) []float32
	Pop(ints []float32) (float32, bool, []float32)
	Push(ints []float32, es ...float32) (int, []float32)
	Reduce(ints []float32, accum func(float32, float32) float32, init ...float32) float32
	ReduceRight(ints []float32, accum func(float32, float32) float32, init ...float32) float32
	Reverse(ints []float32) []float32
	Shift(ints []float32) (float32, bool, []float32)
	Unshift(ints []float32, es ...float32) (int, []float32)
	Slice(ints []float32, startEndIndex ...int) []float32
	Sort(ints []float32, less ...func(int, int) bool)
	Shuffle(ints []float32)
	String(ints []float32) string
	Drop(ints []float32, n ...int) []float32
	Head(ints []float32) (float32, bool)
	Initial(ints []float32) []float32
	Last(ints []float32) (float32, bool)
	Nth(ints []float32, n ...int) (float32, bool)
	Pull(ints []float32, vs ...float32) []float32
	Difference(ints []float32, vs ...[]float32) []float32
	Intersection(vs ...[]float32) []float32
	Without(ints []float32, vs ...float32) []float32
	Remove(ints []float32, predict func(float32) bool) []float32
	Tail(ints []float32) []float32
	Take(ints []float32, n ...int) []float32
	Union(ints ...[]float32) []float32
	Uniq(ints []float32) []float32
	Equal(s1, s2 []float32) bool
}

// NewFloat32 creates a new Float32 instance.
func NewFloat32() *Float32 {
	return &Float32{}
}

var _ Float32Helper = NewFloat32()

// Min returns the smaller of x or y.
func (*Float32) Min(x, y float32) float32 {
	if x < y {
		return x
	}
	return y
}

// Max returns the larger of x or y.
func (*Float32) Max(x, y float32) float32 {
	if x > y {
		return x
	}
	return y
}

// Filter creates a new slice with all elements that pass the predict implemented.
func (*Float32) Filter(ints []float32, predict func(float32) bool) []float32 {
	var r []float32

	for _, v := range ints {
		if ok := predict(v); ok {
			r = append(r, v)
		}
	}

	return r
}

// Find returns the value of the first element in the slice that satisfies the provided predict function.
// otherwise false is returned.
func (*Float32) Find(ints []float32, predict func(float32) bool) (float32, bool) {
	for _, v := range ints {
		if ok := predict(v); ok {
			return v, true
		}
	}

	return 0, false
}

// FindIndex returns the index of the first element in the slice that satisfies the provided predict function.
// otherwise -1 is returned.
func (*Float32) FindIndex(ints []float32, predict func(float32) bool, fromIndex ...int) int {
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
func (*Float32) FindLastIndex(ints []float32, predict func(float32) bool, fromIndex ...int) int {
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
func (*Float32) ForEach(ints []float32, provided func(float32)) {
	for _, v := range ints {
		provided(v)
	}
}

// Includes determines whether a slice includes a certain element, returning true or false.
func (*Float32) Includes(ints []float32, e float32) bool {
	for _, v := range ints {
		if e == v {
			return true
		}
	}

	return false
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
func (*Float32) IndexOf(ints []float32, e float32, fromIndex ...int) int {
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
func (*Float32) LastIndexOf(ints []float32, e float32, fromIndex ...int) int {
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
func (*Float32) From(ints []float32, mapFn ...func(float32) float32) []float32 {
	r := make([]float32, len(ints))
	f := func(i float32) float32 {
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
func (*Float32) Of(ints ...float32) []float32 {
	r := make([]float32, len(ints))

	for i, v := range ints {
		r[i] = v
	}
	return r
}

// Concat merges two or more slices.
func (*Float32) Concat(ints ...[]float32) []float32 {
	var r []float32

	for _, v := range ints {
		r = append(r, v...)
	}

	return r
}

// Every tests whether all elements in the slice pass the predict implemented.
func (*Float32) Every(ints []float32, predict func(float32) bool) bool {
	for _, v := range ints {
		if ok := predict(v); !ok {
			return false
		}
	}

	return true
}

// Some tests whether at least one element in the slice pass the predict implemented.
func (*Float32) Some(ints []float32, predict func(float32) bool) bool {
	for _, v := range ints {
		if ok := predict(v); ok {
			return true
		}
	}

	return false
}

// Fill fills all the elements of a slice from a start index to an end index with a static value.
// the end index is not included.
func (*Float32) Fill(ints []float32, value float32, startEndIndex ...int) {
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
func (*Float32) Join(ints []float32, sep ...string) string {
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
func (*Float32) JoinWithPrec(ints []float32, prec int, sep ...string) string {
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
func (*Float32) Map(ints []float32, m func(float32) float32) []float32 {
	r := make([]float32, len(ints))

	for i, v := range ints {
		r[i] = m(v)
	}

	return r
}

// Pop removes the last element from a slice and returns that element.
func (*Float32) Pop(ints []float32) (float32, bool, []float32) {
	n := len(ints)
	if n == 0 {
		return 0, false, []float32{}
	}
	return ints[n-1], true, ints[:n-1]
}

// Push adds one or more elements to the end of a slice and returns the new length of the slice.
func (*Float32) Push(ints []float32, es ...float32) (int, []float32) {
	var m = len(ints)
	n := m + len(es)
	r := make([]float32, n)
	copy(r, ints)
	for i, v := range es {
		r[i+m] = v
	}
	return n, r
}

// Reduce applies a function against each element in the slice to reduce it to a single value.
func (*Float32) Reduce(ints []float32, accum func(float32, float32) float32, init ...float32) float32 {
	var iv float32
	for _, v := range init {
		iv = v
	}
	for _, v := range ints {
		iv = accum(iv, v)
	}
	return iv
}

// ReduceRight applies a function against each value of the slice (from right-to-left) to reduce it to a single value.
func (*Float32) ReduceRight(ints []float32, accum func(float32, float32) float32, init ...float32) float32 {
	n := len(ints) - 1
	var iv float32
	for _, v := range init {
		iv = v
	}
	for i := range ints {
		iv = accum(iv, ints[n-i])
	}
	return iv
}

// Reverse reverses a slice.
func (*Float32) Reverse(ints []float32) []float32 {
	len := len(ints) - 1
	for i := 0; i < len; i++ {
		ints[i], ints[len] = ints[len], ints[i]
		len--
	}
	return ints
}

// Shift removes the first element from a slice and returns that element.
func (*Float32) Shift(ints []float32) (float32, bool, []float32) {
	n := len(ints)
	if n == 0 {
		return 0, false, []float32{}
	}
	return ints[0], true, ints[1:]
}

// Unshift adds one or more elements to the beginning of a slice and returns the new length of the slice.
func (*Float32) Unshift(ints []float32, es ...float32) (int, []float32) {
	m := len(ints)
	n := len(es)
	r := make([]float32, n+m)

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
func (*Float32) Slice(ints []float32, startEndIndex ...int) []float32 {
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

	r := []float32{}

	for start < end {
		r = append(r, ints[start])
		start++
	}

	return r
}

// Sort sorts elements.
func (*Float32) Sort(ints []float32, less ...func(int, int) bool) {
	swap := func(i, j int) { ints[i], ints[j] = ints[j], ints[i] }
	comp := func(i, j int) bool { return ints[i] < ints[j] }

	if len(less) > 0 {
		comp = less[0]
	}

	sort.Sort(sorter{len: len(ints), swap: swap, less: comp})
}

// Shuffle randomizes the order of elements.
func (*Float32) Shuffle(ints []float32) {
	swap := func(i, j int) { ints[i], ints[j] = ints[j], ints[i] }
	rand.Shuffle(len(ints), swap)
}

// String returns a string representing the specified slice.
func (i *Float32) String(ints []float32) string {
	return i.Join(ints)
}

// Drop creates a slice with n elements dropped from the beginning.
func (*Float32) Drop(ints []float32, n ...int) []float32 {
	var start = 1
	lenin := len(ints)

	if len(n) > 0 && n[0] >= 0 {
		start = n[0]
	}

	if start >= lenin {
		return []float32{}
	}
	r := make([]float32, lenin-start)
	copy(r, ints[start:])
	return r
}

// Head gets the first element of slice.
func (*Float32) Head(ints []float32) (float32, bool) {
	if len(ints) > 0 {
		return ints[0], true
	}

	return 0, false
}

// Initial gets all but the last element of slice.
func (*Float32) Initial(ints []float32) []float32 {
	n := len(ints)
	if n > 1 {
		return ints[:n-1]
	}

	return []float32{}
}

// Last gets the last element of slice.
func (*Float32) Last(ints []float32) (float32, bool) {
	n := len(ints)
	if n > 0 {
		return ints[n-1], true
	}

	return 0, false
}

// Nth gets the element at index n of slice.
// if n is negative, the nth element from the end is returned.
func (*Float32) Nth(ints []float32, n ...int) (float32, bool) {
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
func (i *Float32) Pull(ints []float32, vs ...float32) []float32 {
	var r []float32
	for _, v := range ints {
		if ok := i.Includes(vs, v); !ok {
			r = append(r, v)
		}
	}
	return r
}

// Difference creates a slice of values not included in the other given slices.
func (i *Float32) Difference(ints []float32, vs ...[]float32) []float32 {
	var r []float32
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
func (i *Float32) Intersection(vs ...[]float32) []float32 {
	var r []float32
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
func (i *Float32) Without(ints []float32, vs ...float32) []float32 {
	return i.Pull(ints, vs...)
}

// Remove removes all elements from slice that predicate returns truthy for.
func (*Float32) Remove(ints []float32, predict func(float32) bool) []float32 {
	var r []float32

	for _, v := range ints {
		if ok := predict(v); !ok {
			r = append(r, v)
		}
	}

	return r
}

// Tail gets all but the first element of slice.
func (*Float32) Tail(ints []float32) []float32 {
	n := len(ints)
	if n > 1 {
		return ints[1:]
	}

	return []float32{}
}

// Take creates a slice with n elements taken from the beginning.
func (*Float32) Take(ints []float32, n ...int) []float32 {
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
func (i *Float32) Union(ints ...[]float32) []float32 {
	var r []float32
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
func (i *Float32) Uniq(ints []float32) []float32 {
	var r []float32
	for _, v := range ints {
		if ok := i.Includes(r, v); !ok {
			r = append(r, v)
		}
	}
	return r
}

// Equal performs a deep comparison between two slice values.
func (*Float32) Equal(s1, s2 []float32) bool {
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
