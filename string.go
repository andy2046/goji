package goji

import (
	"math/rand"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

// String for string type slice.
type String struct{}

// StringHelper is the String interface.
type StringHelper interface {
	Min(x, y string) string
	Max(x, y string) string
	Filter(ints []string, predict func(string) bool) []string
	Find(ints []string, predict func(string) bool) (string, bool)
	FindIndex(ints []string, predict func(string) bool, fromIndex ...int) int
	FindLastIndex(ints []string, predict func(string) bool, fromIndex ...int) int
	ForEach(ints []string, provided func(string))
	Includes(ints []string, e string) bool
	IndexOf(ints []string, e string, fromIndex ...int) int
	LastIndexOf(ints []string, e string, fromIndex ...int) int
	From(ints []string, mapFn ...func(string) string) []string
	Of(ints ...string) []string
	Concat(ints ...[]string) []string
	Every(ints []string, predict func(string) bool) bool
	Some(ints []string, predict func(string) bool) bool
	Fill(ints []string, value string, startEndIndex ...int)
	Join(ints []string, sep ...string) string
	Map(ints []string, m func(string) string) []string
	Pop(ints []string) (string, bool, []string)
	Push(ints []string, es ...string) (int, []string)
	Reduce(ints []string, accum func(string, string) string, init ...string) string
	ReduceRight(ints []string, accum func(string, string) string, init ...string) string
	Reverse(ints []string) []string
	Shift(ints []string) (string, bool, []string)
	Unshift(ints []string, es ...string) (int, []string)
	Slice(ints []string, startEndIndex ...int) []string
	Sort(ints []string, less ...func(int, int) bool)
	Shuffle(ints []string)
	String(ints []string) string
	Drop(ints []string, n ...int) []string
	Head(ints []string) (string, bool)
	Initial(ints []string) []string
	Last(ints []string) (string, bool)
	Nth(ints []string, n ...int) (string, bool)
	Pull(ints []string, vs ...string) []string
	Difference(ints []string, vs ...[]string) []string
	Intersection(vs ...[]string) []string
	Without(ints []string, vs ...string) []string
	Remove(ints []string, predict func(string) bool) []string
	Tail(ints []string) []string
	Take(ints []string, n ...int) []string
	Union(ints ...[]string) []string
	Uniq(ints []string) []string
	Equal(s1, s2 []string) bool
}

// NewString creates a new String instance.
func NewString() *String {
	return &String{}
}

var _ StringHelper = NewString()

// Min returns the smaller of x or y.
func (*String) Min(x, y string) string {
	if x < y {
		return x
	}
	return y
}

// Max returns the larger of x or y.
func (*String) Max(x, y string) string {
	if x > y {
		return x
	}
	return y
}

// Filter creates a new slice with all elements that pass the predict implemented.
func (*String) Filter(ints []string, predict func(string) bool) []string {
	var r []string

	for _, v := range ints {
		if ok := predict(v); ok {
			r = append(r, v)
		}
	}

	return r
}

// Find returns the value of the first element in the slice that satisfies the provided predict function.
// otherwise false is returned.
func (*String) Find(ints []string, predict func(string) bool) (string, bool) {
	for _, v := range ints {
		if ok := predict(v); ok {
			return v, true
		}
	}

	return "", false
}

// FindIndex returns the index of the first element in the slice that satisfies the provided predict function.
// otherwise -1 is returned.
func (*String) FindIndex(ints []string, predict func(string) bool, fromIndex ...int) int {
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
func (*String) FindLastIndex(ints []string, predict func(string) bool, fromIndex ...int) int {
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
func (*String) ForEach(ints []string, provided func(string)) {
	for _, v := range ints {
		provided(v)
	}
}

// Includes determines whether a slice includes a certain element, returning true or false.
func (*String) Includes(ints []string, e string) bool {
	for _, v := range ints {
		if e == v {
			return true
		}
	}

	return false
}

// IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.
func (*String) IndexOf(ints []string, e string, fromIndex ...int) int {
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
func (*String) LastIndexOf(ints []string, e string, fromIndex ...int) int {
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
func (*String) From(ints []string, mapFn ...func(string) string) []string {
	r := make([]string, len(ints))
	f := func(i string) string {
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
func (*String) Of(ints ...string) []string {
	r := make([]string, len(ints))

	for i, v := range ints {
		r[i] = v
	}
	return r
}

// Concat merges two or more slices.
func (*String) Concat(ints ...[]string) []string {
	var r []string

	for _, v := range ints {
		r = append(r, v...)
	}

	return r
}

// Every tests whether all elements in the slice pass the predict implemented.
func (*String) Every(ints []string, predict func(string) bool) bool {
	for _, v := range ints {
		if ok := predict(v); !ok {
			return false
		}
	}

	return true
}

// Some tests whether at least one element in the slice pass the predict implemented.
func (*String) Some(ints []string, predict func(string) bool) bool {
	for _, v := range ints {
		if ok := predict(v); ok {
			return true
		}
	}

	return false
}

// Fill fills all the elements of a slice from a start index to an end index with a static value.
// the end index is not included.
func (*String) Fill(ints []string, value string, startEndIndex ...int) {
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
func (*String) Join(ints []string, sep ...string) string {
	s := ","
	for _, se := range sep {
		s = se
	}
	return strings.Join(ints, s)
}

// Map creates a new slice with the results of calling a provided function on every element in the calling slice.
func (*String) Map(ints []string, m func(string) string) []string {
	r := make([]string, len(ints))

	for i, v := range ints {
		r[i] = m(v)
	}

	return r
}

// Pop removes the last element from a slice and returns that element.
func (*String) Pop(ints []string) (string, bool, []string) {
	n := len(ints)
	if n == 0 {
		return "", false, []string{}
	}
	return ints[n-1], true, ints[:n-1]
}

// Push adds one or more elements to the end of a slice and returns the new length of the slice.
func (*String) Push(ints []string, es ...string) (int, []string) {
	var m = len(ints)
	n := m + len(es)
	r := make([]string, n)
	copy(r, ints)
	for i, v := range es {
		r[i+m] = v
	}
	return n, r
}

// Reduce applies a function against each element in the slice to reduce it to a single value.
func (*String) Reduce(ints []string, accum func(string, string) string, init ...string) string {
	var iv string
	for _, v := range init {
		iv = v
	}
	for _, v := range ints {
		iv = accum(iv, v)
	}
	return iv
}

// ReduceRight applies a function against each value of the slice (from right-to-left) to reduce it to a single value.
func (*String) ReduceRight(ints []string, accum func(string, string) string, init ...string) string {
	n := len(ints) - 1
	var iv string
	for _, v := range init {
		iv = v
	}
	for i := range ints {
		iv = accum(iv, ints[n-i])
	}
	return iv
}

// Reverse reverses a slice.
func (*String) Reverse(ints []string) []string {
	len := len(ints) - 1
	for i := 0; i < len; i++ {
		ints[i], ints[len] = ints[len], ints[i]
		len--
	}
	return ints
}

// Shift removes the first element from a slice and returns that element.
func (*String) Shift(ints []string) (string, bool, []string) {
	n := len(ints)
	if n == 0 {
		return "", false, []string{}
	}
	return ints[0], true, ints[1:]
}

// Unshift adds one or more elements to the beginning of a slice and returns the new length of the slice.
func (*String) Unshift(ints []string, es ...string) (int, []string) {
	m := len(ints)
	n := len(es)
	r := make([]string, n+m)

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
func (*String) Slice(ints []string, startEndIndex ...int) []string {
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

	r := []string{}

	for start < end {
		r = append(r, ints[start])
		start++
	}

	return r
}

// Sort sorts elements.
func (*String) Sort(ints []string, less ...func(int, int) bool) {
	swap := func(i, j int) { ints[i], ints[j] = ints[j], ints[i] }
	comp := func(i, j int) bool { return ints[i] < ints[j] }

	if len(less) > 0 {
		comp = less[0]
	}

	sort.Sort(sorter{len: len(ints), swap: swap, less: comp})
}

// Shuffle randomizes the order of elements.
func (*String) Shuffle(ints []string) {
	swap := func(i, j int) { ints[i], ints[j] = ints[j], ints[i] }
	rand.Shuffle(len(ints), swap)
}

// String returns a string representing the specified slice.
func (i *String) String(ints []string) string {
	return i.Join(ints)
}

// Drop creates a slice with n elements dropped from the beginning.
func (*String) Drop(ints []string, n ...int) []string {
	var start = 1
	lenin := len(ints)

	if len(n) > 0 && n[0] >= 0 {
		start = n[0]
	}

	if start >= lenin {
		return []string{}
	}
	r := make([]string, lenin-start)
	copy(r, ints[start:])
	return r
}

// Head gets the first element of slice.
func (*String) Head(ints []string) (string, bool) {
	if len(ints) > 0 {
		return ints[0], true
	}

	return "", false
}

// Initial gets all but the last element of slice.
func (*String) Initial(ints []string) []string {
	n := len(ints)
	if n > 1 {
		return ints[:n-1]
	}

	return []string{}
}

// Last gets the last element of slice.
func (*String) Last(ints []string) (string, bool) {
	n := len(ints)
	if n > 0 {
		return ints[n-1], true
	}

	return "", false
}

// Nth gets the element at index n of slice.
// if n is negative, the nth element from the end is returned.
func (*String) Nth(ints []string, n ...int) (string, bool) {
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

	return "", false
}

// Pull removes all given values from slice.
func (i *String) Pull(ints []string, vs ...string) []string {
	var r []string
	for _, v := range ints {
		if ok := i.Includes(vs, v); !ok {
			r = append(r, v)
		}
	}
	return r
}

// Difference creates a slice of values not included in the other given slices.
func (i *String) Difference(ints []string, vs ...[]string) []string {
	var r []string
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
func (i *String) Intersection(vs ...[]string) []string {
	var r []string
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
func (i *String) Without(ints []string, vs ...string) []string {
	return i.Pull(ints, vs...)
}

// Remove removes all elements from slice that predicate returns truthy for.
func (*String) Remove(ints []string, predict func(string) bool) []string {
	var r []string

	for _, v := range ints {
		if ok := predict(v); !ok {
			r = append(r, v)
		}
	}

	return r
}

// Tail gets all but the first element of slice.
func (*String) Tail(ints []string) []string {
	n := len(ints)
	if n > 1 {
		return ints[1:]
	}

	return []string{}
}

// Take creates a slice with n elements taken from the beginning.
func (*String) Take(ints []string, n ...int) []string {
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
func (i *String) Union(ints ...[]string) []string {
	var r []string
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
func (i *String) Uniq(ints []string) []string {
	var r []string
	for _, v := range ints {
		if ok := i.Includes(r, v); !ok {
			r = append(r, v)
		}
	}
	return r
}

// Equal performs a deep comparison between two slice values.
func (*String) Equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return strings.Join(s1, ",") == strings.Join(s2, ",")
}

func (*String) preserveCamelCase(input string) string {
	var isLastCharLower, isLastCharUpper, isLastLastCharUpper bool
	inputAsRunes := []rune(input)

	for i, c := range inputAsRunes {
		sc := string(c)
		matched, _ := regexp.MatchString(`[a-zA-Z]`, sc)

		switch {
		case isLastCharLower && matched && unicode.ToUpper(c) == c:
			temp := append([]rune{'-'}, inputAsRunes[i:]...)
			inputAsRunes = append(inputAsRunes[:i], temp...)
			isLastCharLower = false
			isLastLastCharUpper = isLastCharUpper
			isLastCharUpper = true
		case isLastCharUpper && isLastLastCharUpper && matched && unicode.ToLower(c) == c:
			temp := append([]rune{'-'}, inputAsRunes[i-1:]...)
			inputAsRunes = append(inputAsRunes[:i-1], temp...)
			isLastLastCharUpper = isLastCharUpper
			isLastCharUpper = false
			isLastCharLower = true
		default:
			isLastCharLower = (unicode.ToLower(c) == c)
			isLastLastCharUpper = isLastCharUpper
			isLastCharUpper = (unicode.ToUpper(c) == c)
		}
	}

	return string(inputAsRunes)
}

// CamelCase convert dash/dot/underscore/space separated string to camelCase
// pascalCase define whether to uppercase the first character
func (i *String) CamelCase(input string, pascalCase ...bool) string {
	var pCase bool
	for _, p := range pascalCase {
		pCase = p
	}

	postProcess := func(x string) string {
		if !pCase {
			return x
		}
		xAsRunes := []rune(x)
		xAsRunes[0] = unicode.ToUpper(xAsRunes[0])
		return string(xAsRunes)
	}

	inputX := strings.TrimSpace(input)

	if len(inputX) == 0 {
		return ""
	}

	if len(inputX) == 1 {
		if pCase {
			return strings.ToUpper(inputX)
		}
		return strings.ToLower(inputX)
	}

	matched, _ := regexp.MatchString(`^[a-z\d]+$`, inputX)
	if matched {
		return postProcess(inputX)
	}

	if inputX != strings.ToLower(inputX) {
		inputX = i.preserveCamelCase(inputX)
	}

	re, _ := regexp.Compile(`^[_.\- ]+`)
	inputX = strings.ToLower(re.ReplaceAllString(inputX, ""))

	re2, _ := regexp.Compile(`[_.\- ]+(\w|$)`)
	inputX = re2.ReplaceAllStringFunc(inputX, func(a string) string {
		return strings.ToUpper(re.ReplaceAllString(a, ""))
	})

	return postProcess(inputX)
}
