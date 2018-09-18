

# goji
`import "github.com/andy2046/goji"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package goji provides functional utils for int / float / string Slice.




## <a name="pkg-index">Index</a>
* [func Concat()](#Concat)
* [func Difference(lenI, lenJ int, comparator func(int, int) bool, handlers ...func(...int)) []int](#Difference)
* [func Every(n int, predict func(int) bool) bool](#Every)
* [func Fill(fill func(int), n int, startEndIndex ...int)](#Fill)
* [func Filter(n int, predict func(int) bool, handlers ...func(...int)) []int](#Filter)
* [func FindIndex(n int, testing func(int) bool, fromIndex ...int) int](#FindIndex)
* [func FindLastIndex(n int, testing func(int) bool, fromIndex ...int) int](#FindLastIndex)
* [func Includes(n int, comp func(int) bool) bool](#Includes)
* [func Intersection(lenI, lenJ int, comparator func(int, int) bool, handlers ...func(...int)) []int](#Intersection)
* [func Join(join func(int) string, n int, sep ...string) string](#Join)
* [func Map(n int, m func(int))](#Map)
* [func Max(x, y int) int](#Max)
* [func Min(x, y int) int](#Min)
* [func Range(end int) []struct{}](#Range)
* [func Reverse(n int, swap func(int, int))](#Reverse)
* [func Shuffle(n int, swap func(int, int))](#Shuffle)
* [func Slice(slice func(int, int), n int, startEndIndex ...int)](#Slice)
* [func Some(n int, predict func(int) bool) bool](#Some)
* [func Sort(n int, swap func(int, int), less func(int, int) bool)](#Sort)
* [type Sorter](#Sorter)
  * [func (x Sorter) Len() int](#Sorter.Len)
  * [func (x Sorter) Less(i, j int) bool](#Sorter.Less)
  * [func (x Sorter) Swap(i, j int)](#Sorter.Swap)


#### <a name="pkg-files">Package files</a>
[goji.go](/src/github.com/andy2046/goji/goji.go) 





## <a name="Concat">func</a> [Concat](/src/target/goji.go?s=2476:2489#L105)
``` go
func Concat()
```
Concat merges two slices.



## <a name="Difference">func</a> [Difference](/src/target/goji.go?s=5757:5852#L254)
``` go
func Difference(lenI, lenJ int, comparator func(int, int) bool, handlers ...func(...int)) []int
```
Difference returns difference of array/slice I's values not included in the other given array J.
lenI, lenJ is the length of array/slice I, J.
comparator compares values from array/slice I, J.



## <a name="Every">func</a> [Every](/src/target/goji.go?s=2671:2717#L111)
``` go
func Every(n int, predict func(int) bool) bool
```
Every tests whether all elements pass the test implemented by predict.
n is the number of elements. j is the index of array/slice.



## <a name="Fill">func</a> [Fill](/src/target/goji.go?s=4361:4415#L180)
``` go
func Fill(fill func(int), n int, startEndIndex ...int)
```
Fill fills all elements from a start index to an end index with a static value.
the end index is not included.
n is the number of elements. start and end index are optional, default to zero and n.



## <a name="Filter">func</a> [Filter](/src/target/goji.go?s=1495:1569#L63)
``` go
func Filter(n int, predict func(int) bool, handlers ...func(...int)) []int
```
Filter filters elements, n is the number of elements.
predict check element with index i and returns true or false.
handlers receive accumulated indexes idx from predict, if any.



## <a name="FindIndex">func</a> [FindIndex](/src/target/goji.go?s=4900:4967#L208)
``` go
func FindIndex(n int, testing func(int) bool, fromIndex ...int) int
```
FindIndex returns the index of the first element that satisfies the provided testing function.
Otherwise -1 is returned. n is the number of elements.



## <a name="FindLastIndex">func</a> [FindLastIndex](/src/target/goji.go?s=5253:5324#L230)
``` go
func FindLastIndex(n int, testing func(int) bool, fromIndex ...int) int
```
FindLastIndex iterates over elements from right to left.



## <a name="Includes">func</a> [Includes](/src/target/goji.go?s=1943:1989#L81)
``` go
func Includes(n int, comp func(int) bool) bool
```
Includes determines whether an array/slice includes a certain element, returning true or false.
n is the number of elements. comp determines whether the given element with index j is found.



## <a name="Intersection">func</a> [Intersection](/src/target/goji.go?s=6222:6319#L276)
``` go
func Intersection(lenI, lenJ int, comparator func(int, int) bool, handlers ...func(...int)) []int
```
Intersection returns values that are included in both given array/slice.



## <a name="Join">func</a> [Join](/src/target/goji.go?s=3941:4002#L164)
``` go
func Join(join func(int) string, n int, sep ...string) string
```
Join joins all elements of array/slice into a string.
The separator string sep is placed between elements in the resulting string.
n is the number of elements. func join receive element index i and return a string.



## <a name="Map">func</a> [Map](/src/target/goji.go?s=2185:2213#L92)
``` go
func Map(n int, m func(int))
```
Map calls a provided function m on every element.
n is the number of elements. m map element with index i.



## <a name="Max">func</a> [Max](/src/target/goji.go?s=278:300#L20)
``` go
func Max(x, y int) int
```
Max returns the larger of x or y.



## <a name="Min">func</a> [Min](/src/target/goji.go?s=177:199#L12)
``` go
func Min(x, y int) int
```
Min returns the smaller of x or y.



## <a name="Range">func</a> [Range](/src/target/goji.go?s=430:460#L28)
``` go
func Range(end int) []struct{}
```
Range creates a range of numbers progressing from zero up to, but not including end.



## <a name="Reverse">func</a> [Reverse](/src/target/goji.go?s=1193:1233#L52)
``` go
func Reverse(n int, swap func(int, int))
```
Reverse reverses an array/slice, n is the number of elements.
swap swaps the elements with indexes i and j.



## <a name="Shuffle">func</a> [Shuffle](/src/target/goji.go?s=2378:2418#L100)
``` go
func Shuffle(n int, swap func(int, int))
```
Shuffle randomizes the order of elements.
n is the number of elements. swap swaps the elements with indexes i and j.



## <a name="Slice">func</a> [Slice](/src/target/goji.go?s=3332:3393#L136)
``` go
func Slice(slice func(int, int), n int, startEndIndex ...int)
```
Slice returns a portion of array/slice selected from start to end (end not included).
n is the number of elements. start and end indexes are optional, default to zero and n.
start and end indexes are passed to slice func.



## <a name="Some">func</a> [Some](/src/target/goji.go?s=2959:3004#L123)
``` go
func Some(n int, predict func(int) bool) bool
```
Some tests whether at least one element pass the test implemented by predict.
n is the number of elements. j is the index of array/slice.



## <a name="Sort">func</a> [Sort](/src/target/goji.go?s=953:1016#L46)
``` go
func Sort(n int, swap func(int, int), less func(int, int) bool)
```
Sort sorts elements, n is the number of elements.
less compares the elements with indexes i and j.
swap swaps the elements with indexes i and j.




## <a name="Sorter">type</a> [Sorter](/src/target/goji.go?s=533:624#L33)
``` go
type Sorter struct {
    N        int
    SwapFunc func(i, j int)
    LessFunc func(i, j int) bool
}
```
Sorter implements Sort interface.










### <a name="Sorter.Len">func</a> (Sorter) [Len](/src/target/goji.go?s=626:651#L39)
``` go
func (x Sorter) Len() int
```



### <a name="Sorter.Less">func</a> (Sorter) [Less](/src/target/goji.go?s=734:769#L41)
``` go
func (x Sorter) Less(i, j int) bool
```



### <a name="Sorter.Swap">func</a> (Sorter) [Swap](/src/target/goji.go?s=677:707#L40)
``` go
func (x Sorter) Swap(i, j int)
```







- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
