

# goji
`import "github.com/andy2046/goji"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)


## <a name="pkg-overview">Overview</a>
Package goji provides functional utils for Slice.




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
* [type Int](#Int)
  * [func NewInt() *Int](#NewInt)
  * [func (*Int) Concat(ints ...[]int) []int](#Int.Concat)
  * [func (*Int) Difference(ints []int, vs ...[]int) []int](#Int.Difference)
  * [func (*Int) Drop(ints []int, n ...int) []int](#Int.Drop)
  * [func (*Int) Equal(s1, s2 []int) bool](#Int.Equal)
  * [func (*Int) Every(ints []int, predict func(int) bool) bool](#Int.Every)
  * [func (*Int) Fill(ints []int, value int, startEndIndex ...int)](#Int.Fill)
  * [func (*Int) Filter(ints []int, predict func(int) bool) []int](#Int.Filter)
  * [func (*Int) Find(ints []int, predict func(int) bool) (int, bool)](#Int.Find)
  * [func (*Int) FindIndex(ints []int, predict func(int) bool, fromIndex ...int) int](#Int.FindIndex)
  * [func (*Int) FindLastIndex(ints []int, predict func(int) bool, fromIndex ...int) int](#Int.FindLastIndex)
  * [func (*Int) ForEach(ints []int, provided func(int))](#Int.ForEach)
  * [func (*Int) From(ints []int, mapFn ...func(int) int) []int](#Int.From)
  * [func (*Int) Head(ints []int) (int, bool)](#Int.Head)
  * [func (*Int) Includes(ints []int, e int) bool](#Int.Includes)
  * [func (*Int) IndexOf(ints []int, e int, fromIndex ...int) int](#Int.IndexOf)
  * [func (*Int) Initial(ints []int) []int](#Int.Initial)
  * [func (*Int) Intersection(vs ...[]int) []int](#Int.Intersection)
  * [func (*Int) Join(ints []int, sep ...string) string](#Int.Join)
  * [func (*Int) Last(ints []int) (int, bool)](#Int.Last)
  * [func (*Int) LastIndexOf(ints []int, e int, fromIndex ...int) int](#Int.LastIndexOf)
  * [func (*Int) Map(ints []int, m func(int) int) []int](#Int.Map)
  * [func (*Int) Nth(ints []int, n ...int) (int, bool)](#Int.Nth)
  * [func (*Int) Of(ints ...int) []int](#Int.Of)
  * [func (*Int) Pop(ints []int) (int, bool, []int)](#Int.Pop)
  * [func (*Int) Pull(ints []int, vs ...int) []int](#Int.Pull)
  * [func (*Int) Push(ints []int, es ...int) (int, []int)](#Int.Push)
  * [func (*Int) Reduce(ints []int, accum func(int, int) int, init ...int) int](#Int.Reduce)
  * [func (*Int) ReduceRight(ints []int, accum func(int, int) int, init ...int) int](#Int.ReduceRight)
  * [func (*Int) Remove(ints []int, predict func(int) bool) []int](#Int.Remove)
  * [func (*Int) Reverse(ints []int) []int](#Int.Reverse)
  * [func (*Int) Shift(ints []int) (int, bool, []int)](#Int.Shift)
  * [func (*Int) Shuffle(ints []int)](#Int.Shuffle)
  * [func (*Int) Slice(ints []int, startEndIndex ...int) []int](#Int.Slice)
  * [func (*Int) Some(ints []int, predict func(int) bool) bool](#Int.Some)
  * [func (*Int) Sort(ints []int, less ...func(int, int) bool)](#Int.Sort)
  * [func (*Int) String(ints []int) string](#Int.String)
  * [func (*Int) Tail(ints []int) []int](#Int.Tail)
  * [func (*Int) Take(ints []int, n ...int) []int](#Int.Take)
  * [func (*Int) Union(ints ...[]int) []int](#Int.Union)
  * [func (*Int) Uniq(ints []int) []int](#Int.Uniq)
  * [func (*Int) Unshift(ints []int, es ...int) (int, []int)](#Int.Unshift)
  * [func (*Int) Without(ints []int, vs ...int) []int](#Int.Without)
* [type IntHelper](#IntHelper)


#### <a name="pkg-files">Package files</a>
[goji.go](./goji.go) [int.go](./int.go) 





## <a name="Concat">func</a> [Concat](./goji.go?s=2395:2408#L104)
``` go
func Concat()
```
Concat merges two slices.



## <a name="Difference">func</a> [Difference](./goji.go?s=5676:5771#L253)
``` go
func Difference(lenI, lenJ int, comparator func(int, int) bool, handlers ...func(...int)) []int
```
Difference returns difference of array/slice I's values not included in the other given array J.
lenI, lenJ is the length of array/slice I, J.
comparator compares values from array/slice I, J.



## <a name="Every">func</a> [Every](./goji.go?s=2590:2636#L110)
``` go
func Every(n int, predict func(int) bool) bool
```
Every tests whether all elements pass the test implemented by predict.
n is the number of elements. j is the index of array/slice.



## <a name="Fill">func</a> [Fill](./goji.go?s=4280:4334#L179)
``` go
func Fill(fill func(int), n int, startEndIndex ...int)
```
Fill fills all elements from a start index to an end index with a static value.
the end index is not included.
n is the number of elements. start and end index are optional, default to zero and n.



## <a name="Filter">func</a> [Filter](./goji.go?s=1414:1488#L62)
``` go
func Filter(n int, predict func(int) bool, handlers ...func(...int)) []int
```
Filter filters elements, n is the number of elements.
predict check element with index i and returns true or false.
handlers receive accumulated indexes idx from predict, if any.



## <a name="FindIndex">func</a> [FindIndex](./goji.go?s=4819:4886#L207)
``` go
func FindIndex(n int, testing func(int) bool, fromIndex ...int) int
```
FindIndex returns the index of the first element that satisfies the provided testing function.
Otherwise -1 is returned. n is the number of elements.



## <a name="FindLastIndex">func</a> [FindLastIndex](./goji.go?s=5172:5243#L229)
``` go
func FindLastIndex(n int, testing func(int) bool, fromIndex ...int) int
```
FindLastIndex iterates over elements from right to left.



## <a name="Includes">func</a> [Includes](./goji.go?s=1862:1908#L80)
``` go
func Includes(n int, comp func(int) bool) bool
```
Includes determines whether an array/slice includes a certain element, returning true or false.
n is the number of elements. comp determines whether the given element with index j is found.



## <a name="Intersection">func</a> [Intersection](./goji.go?s=6141:6238#L275)
``` go
func Intersection(lenI, lenJ int, comparator func(int, int) bool, handlers ...func(...int)) []int
```
Intersection returns values that are included in both given array/slice.



## <a name="Join">func</a> [Join](./goji.go?s=3860:3921#L163)
``` go
func Join(join func(int) string, n int, sep ...string) string
```
Join joins all elements of array/slice into a string.
The separator string sep is placed between elements in the resulting string.
n is the number of elements. func join receive element index i and return a string.



## <a name="Map">func</a> [Map](./goji.go?s=2104:2132#L91)
``` go
func Map(n int, m func(int))
```
Map calls a provided function m on every element.
n is the number of elements. m map element with index i.



## <a name="Max">func</a> [Max](./goji.go?s=258:280#L20)
``` go
func Max(x, y int) int
```
Max returns the larger of x or y.



## <a name="Min">func</a> [Min](./goji.go?s=157:179#L12)
``` go
func Min(x, y int) int
```
Min returns the smaller of x or y.



## <a name="Range">func</a> [Range](./goji.go?s=410:440#L28)
``` go
func Range(end int) []struct{}
```
Range creates a range of numbers progressing from zero up to, but not including end.



## <a name="Reverse">func</a> [Reverse](./goji.go?s=1112:1152#L51)
``` go
func Reverse(n int, swap func(int, int))
```
Reverse reverses an array/slice, n is the number of elements.
swap swaps the elements with indexes i and j.



## <a name="Shuffle">func</a> [Shuffle](./goji.go?s=2297:2337#L99)
``` go
func Shuffle(n int, swap func(int, int))
```
Shuffle randomizes the order of elements.
n is the number of elements. swap swaps the elements with indexes i and j.



## <a name="Slice">func</a> [Slice](./goji.go?s=3251:3312#L135)
``` go
func Slice(slice func(int, int), n int, startEndIndex ...int)
```
Slice returns a portion of array/slice selected from start to end (end not included).
n is the number of elements. start and end indexes are optional, default to zero and n.
start and end indexes are passed to slice func.



## <a name="Some">func</a> [Some](./goji.go?s=2878:2923#L122)
``` go
func Some(n int, predict func(int) bool) bool
```
Some tests whether at least one element pass the test implemented by predict.
n is the number of elements. j is the index of array/slice.



## <a name="Sort">func</a> [Sort](./goji.go?s=878:941#L45)
``` go
func Sort(n int, swap func(int, int), less func(int, int) bool)
```
Sort sorts elements, n is the number of elements.
less compares the elements with indexes i and j.
swap swaps the elements with indexes i and j.




## <a name="Int">type</a> [Int](./int.go?s=96:113#L11)
``` go
type Int struct{}
```
Int for int type slice.







### <a name="NewInt">func</a> [NewInt](./int.go?s=289:307#L20)
``` go
func NewInt() *Int
```
NewInt creates a new Int instance.





### <a name="Int.Concat">func</a> (\*Int) [Concat](./int.go?s=3523:3562#L189)
``` go
func (*Int) Concat(ints ...[]int) []int
```
Concat merges two or more slices.




### <a name="Int.Difference">func</a> (\*Int) [Difference](./int.go?s=9681:9736#L493)
``` go
func (i *Int) Difference(ints []int, vs ...[]int) []int
```
Difference creates a slice of values not included in the other given slices.




### <a name="Int.Drop">func</a> (\*Int) [Drop](./int.go?s=8320:8364#L415)
``` go
func (*Int) Drop(ints []int, n ...int) []int
```
Drop creates a slice with n elements dropped from the beginning.




### <a name="Int.Equal">func</a> (\*Int) [Equal](./int.go?s=11751:11789#L609)
``` go
func (i *Int) Equal(s1, s2 []int) bool
```
Equal performs a deep comparison between two slice values.




### <a name="Int.Every">func</a> (\*Int) [Every](./int.go?s=3723:3781#L200)
``` go
func (*Int) Every(ints []int, predict func(int) bool) bool
```
Every tests whether all elements in the slice pass the predict implemented.




### <a name="Int.Fill">func</a> (\*Int) [Fill](./int.go?s=4251:4312#L223)
``` go
func (*Int) Fill(ints []int, value int, startEndIndex ...int)
```
Fill fills all the elements of a slice from a start index to an end index with a static value.
the end index is not included.




### <a name="Int.Filter">func</a> (\*Int) [Filter](./int.go?s=411:471#L25)
``` go
func (*Int) Filter(ints []int, predict func(int) bool) []int
```
Filter creates a new slice with all elements that pass the predict implemented.




### <a name="Int.Find">func</a> (\*Int) [Find](./int.go?s=721:785#L39)
``` go
func (*Int) Find(ints []int, predict func(int) bool) (int, bool)
```
Find returns the value of the first element in the slice that satisfies the provided predict function.
otherwise false is returned.




### <a name="Int.FindIndex">func</a> (\*Int) [FindIndex](./int.go?s=1028:1107#L51)
``` go
func (*Int) FindIndex(ints []int, predict func(int) bool, fromIndex ...int) int
```
FindIndex returns the index of the first element in the slice that satisfies the provided predict function.
otherwise -1 is returned.




### <a name="Int.FindLastIndex">func</a> (\*Int) [FindLastIndex](./int.go?s=1449:1532#L74)
``` go
func (*Int) FindLastIndex(ints []int, predict func(int) bool, fromIndex ...int) int
```
FindLastIndex  is like FindIndex except that it iterates over elements from right to left.




### <a name="Int.ForEach">func</a> (\*Int) [ForEach](./int.go?s=1854:1905#L97)
``` go
func (*Int) ForEach(ints []int, provided func(int))
```
ForEach executes a provided function once for each slice element.




### <a name="Int.From">func</a> (\*Int) [From](./int.go?s=3044:3102#L161)
``` go
func (*Int) From(ints []int, mapFn ...func(int) int) []int
```
From creates a new, shallow-copied slice.




### <a name="Int.Head">func</a> (\*Int) [Head](./int.go?s=8598:8638#L432)
``` go
func (*Int) Head(ints []int) (int, bool)
```
Head gets the first element of slice.




### <a name="Int.Includes">func</a> (\*Int) [Includes](./int.go?s=2046:2090#L104)
``` go
func (*Int) Includes(ints []int, e int) bool
```
Includes determines whether a slice includes a certain element, returning true or false.




### <a name="Int.IndexOf">func</a> (\*Int) [IndexOf](./int.go?s=2288:2348#L115)
``` go
func (*Int) IndexOf(ints []int, e int, fromIndex ...int) int
```
IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.




### <a name="Int.Initial">func</a> (\*Int) [Initial](./int.go?s=8759:8796#L441)
``` go
func (*Int) Initial(ints []int) []int
```
Initial gets all but the last element of slice.




### <a name="Int.Intersection">func</a> (\*Int) [Intersection](./int.go?s=10056:10101#L513)
``` go
func (i *Int) Intersection(vs ...[]int) []int
```
Intersection creates a slice of unique values that are included in all given slices.




### <a name="Int.Join">func</a> (\*Int) [Join](./int.go?s=4718:4768#L251)
``` go
func (*Int) Join(ints []int, sep ...string) string
```
Join joins all elements of a slice into a string.




### <a name="Int.Last">func</a> (\*Int) [Last](./int.go?s=8910:8950#L451)
``` go
func (*Int) Last(ints []int) (int, bool)
```
Last gets the last element of slice.




### <a name="Int.LastIndexOf">func</a> (\*Int) [LastIndexOf](./int.go?s=2700:2764#L138)
``` go
func (*Int) LastIndexOf(ints []int, e int, fromIndex ...int) int
```
LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.




### <a name="Int.Map">func</a> (\*Int) [Map](./int.go?s=5051:5101#L266)
``` go
func (*Int) Map(ints []int, m func(int) int) []int
```
Map creates a new slice with the results of calling a provided function on every element in the calling slice.




### <a name="Int.Nth">func</a> (\*Int) [Nth](./int.go?s=9138:9187#L462)
``` go
func (*Int) Nth(ints []int, n ...int) (int, bool)
```
Nth gets the element at index n of slice.
if n is negative, the nth element from the end is returned.




### <a name="Int.Of">func</a> (\*Int) [Of](./int.go?s=3367:3400#L179)
``` go
func (*Int) Of(ints ...int) []int
```
Of creates a new slice instance with a variable number of arguments.




### <a name="Int.Pop">func</a> (\*Int) [Pop](./int.go?s=5262:5308#L277)
``` go
func (*Int) Pop(ints []int) (int, bool, []int)
```
Pop removes the last element from a slice and returns that element.




### <a name="Int.Pull">func</a> (\*Int) [Pull](./int.go?s=9436:9483#L482)
``` go
func (i *Int) Pull(ints []int, vs ...int) []int
```
Pull removes all given values from slice.




### <a name="Int.Push">func</a> (\*Int) [Push](./int.go?s=5506:5558#L286)
``` go
func (*Int) Push(ints []int, es ...int) (int, []int)
```
Push adds one or more elements to the end of a slice and returns the new length of the slice.




### <a name="Int.Reduce">func</a> (\*Int) [Reduce](./int.go?s=5782:5855#L298)
``` go
func (*Int) Reduce(ints []int, accum func(int, int) int, init ...int) int
```
Reduce applies a function against each element in the slice to reduce it to a single value.




### <a name="Int.ReduceRight">func</a> (\*Int) [ReduceRight](./int.go?s=6087:6165#L310)
``` go
func (*Int) ReduceRight(ints []int, accum func(int, int) int, init ...int) int
```
ReduceRight applies a function against each value of the slice (from right-to-left) to reduce it to a single value.




### <a name="Int.Remove">func</a> (\*Int) [Remove](./int.go?s=10631:10691#L546)
``` go
func (*Int) Remove(ints []int, predict func(int) bool) []int
```
Remove removes all elements from slice that predicate returns truthy for.




### <a name="Int.Reverse">func</a> (\*Int) [Reverse](./int.go?s=6332:6369#L323)
``` go
func (*Int) Reverse(ints []int) []int
```
Reverse reverses a slice.




### <a name="Int.Shift">func</a> (\*Int) [Shift](./int.go?s=6565:6613#L333)
``` go
func (*Int) Shift(ints []int) (int, bool, []int)
```
Shift removes the first element from a slice and returns that element.




### <a name="Int.Shuffle">func</a> (\*Int) [Shuffle](./int.go?s=7993:8024#L404)
``` go
func (*Int) Shuffle(ints []int)
```
Shuffle randomizes the order of elements.




### <a name="Int.Slice">func</a> (\*Int) [Slice](./int.go?s=7195:7252#L360)
``` go
func (*Int) Slice(ints []int, startEndIndex ...int) []int
```
Slice returns a portion of a slice selected from start to end (end not included).
start and end indexes are optional, default to zero and length of the slice.




### <a name="Int.Some">func</a> (\*Int) [Some](./int.go?s=3965:4022#L211)
``` go
func (*Int) Some(ints []int, predict func(int) bool) bool
```
Some tests whether at least one element in the slice pass the predict implemented.




### <a name="Int.Sort">func</a> (\*Int) [Sort](./int.go?s=7662:7719#L392)
``` go
func (*Int) Sort(ints []int, less ...func(int, int) bool)
```
Sort sorts elements.




### <a name="Int.String">func</a> (\*Int) [String](./int.go?s=8186:8225#L410)
``` go
func (i *Int) String(ints []int) string
```
String returns a string representing the specified slice.




### <a name="Int.Tail">func</a> (\*Int) [Tail](./int.go?s=10853:10887#L559)
``` go
func (*Int) Tail(ints []int) []int
```
Tail gets all but the first element of slice.




### <a name="Int.Take">func</a> (\*Int) [Take](./int.go?s=11025:11069#L569)
``` go
func (*Int) Take(ints []int, n ...int) []int
```
Take creates a slice with n elements taken from the beginning.




### <a name="Int.Union">func</a> (\*Int) [Union](./int.go?s=11293:11333#L585)
``` go
func (i *Int) Union(ints ...[]int) []int
```
Union creates a slice of unique values, in order, from all given slices.




### <a name="Int.Uniq">func</a> (\*Int) [Uniq](./int.go?s=11536:11572#L598)
``` go
func (i *Int) Uniq(ints []int) []int
```
Uniq creates a duplicate-free version of a slice.




### <a name="Int.Unshift">func</a> (\*Int) [Unshift](./int.go?s=6816:6871#L342)
``` go
func (*Int) Unshift(ints []int, es ...int) (int, []int)
```
Unshift adds one or more elements to the beginning of a slice and returns the new length of the slice.




### <a name="Int.Without">func</a> (\*Int) [Without](./int.go?s=10470:10520#L541)
``` go
func (i *Int) Without(ints []int, vs ...int) []int
```
Without creates a slice excluding all given values.




## <a name="IntHelper">type</a> [IntHelper](./int.go?s=150:249#L14)
``` go
type IntHelper interface {
    Filter(ints []int, predict func(int) bool) []int
    Find(ints []int, predict func(int) bool) (int, bool)
    FindIndex(ints []int, predict func(int) bool, fromIndex ...int) int
    Includes(ints []int, e int) bool
    IndexOf(ints []int, e int, fromIndex ...int) int
    From(ints []int, mapFn ...func(int) int) []int
    Of(ints ...int) []int
    Concat(ints ...[]int) []int
    Every(ints []int, predict func(int) bool) bool
    Some(ints []int, predict func(int) bool) bool
    Fill(ints []int, value int, startEndIndex ...int)
    Join(ints []int, sep ...string) string
    Map(ints []int, m func(int) int) []int
    Reduce(ints []int, accum func(int, int) int, init ...int) int
    // ...
}
```
IntHelper is the Int interface.





