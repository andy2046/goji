

# goji
`import "github.com/andy2046/goji"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)


## <a name="pkg-overview">Overview</a>
Package goji provides functional utils for `int` / `float` / `string` Slice.



## <a name="pkg-index">Index</a>
* [func Concat()](#Concat)
* [func Difference(lenI, lenJ int, comparator func(int, int) bool, handlers ...func(...int)) []int](#Difference)
* [func Every(n int, predict func(int) bool) bool](#Every)
* [func Fill(fill func(int), n int, startEndIndex ...int)](#Fill)
* [func Filter(n int, predict func(int) bool, handlers ...func(...int)) []int](#Filter)
* [func FindIndex(n int, testing func(int) bool, fromIndex ...int) int](#FindIndex)
* [func FindLastIndex(n int, testing func(int) bool, fromIndex ...int) int](#FindLastIndex)
* [func FloatToString(fv float64) string](#FloatToString)
* [func FloatToStringWithPrec(fv float64, prec int) string](#FloatToStringWithPrec)
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
  * [func (i *Int) Difference(ints []int, vs ...[]int) []int](#Int.Difference)
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
  * [func (i *Int) Intersection(vs ...[]int) []int](#Int.Intersection)
  * [func (*Int) Join(ints []int, sep ...string) string](#Int.Join)
  * [func (*Int) Last(ints []int) (int, bool)](#Int.Last)
  * [func (*Int) LastIndexOf(ints []int, e int, fromIndex ...int) int](#Int.LastIndexOf)
  * [func (*Int) Map(ints []int, m func(int) int) []int](#Int.Map)
  * [func (*Int) Nth(ints []int, n ...int) (int, bool)](#Int.Nth)
  * [func (*Int) Of(ints ...int) []int](#Int.Of)
  * [func (*Int) Pop(ints []int) (int, bool, []int)](#Int.Pop)
  * [func (i *Int) Pull(ints []int, vs ...int) []int](#Int.Pull)
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
  * [func (i *Int) String(ints []int) string](#Int.String)
  * [func (*Int) Tail(ints []int) []int](#Int.Tail)
  * [func (*Int) Take(ints []int, n ...int) []int](#Int.Take)
  * [func (i *Int) Union(ints ...[]int) []int](#Int.Union)
  * [func (i *Int) Uniq(ints []int) []int](#Int.Uniq)
  * [func (*Int) Unshift(ints []int, es ...int) (int, []int)](#Int.Unshift)
  * [func (i *Int) Without(ints []int, vs ...int) []int](#Int.Without)

* [type IntHelper](#IntHelper)

* [type String](#String)
  * [func NewString() *String](#NewString)
  * [func (i *String) CamelCase(input string, pascalCase ...bool) string](#String.CamelCase)
  * [func (*String) Concat(ints ...[]string) []string](#String.Concat)
  * [func (i *String) Difference(ints []string, vs ...[]string) []string](#String.Difference)
  * [func (*String) Drop(ints []string, n ...int) []string](#String.Drop)
  * [func (*String) Equal(s1, s2 []string) bool](#String.Equal)
  * [func (*String) Every(ints []string, predict func(string) bool) bool](#String.Every)
  * [func (*String) Fill(ints []string, value string, startEndIndex ...int)](#String.Fill)
  * [func (*String) Filter(ints []string, predict func(string) bool) []string](#String.Filter)
  * [func (*String) Find(ints []string, predict func(string) bool) (string, bool)](#String.Find)
  * [func (*String) FindIndex(ints []string, predict func(string) bool, fromIndex ...int) int](#String.FindIndex)
  * [func (*String) FindLastIndex(ints []string, predict func(string) bool, fromIndex ...int) int](#String.FindLastIndex)
  * [func (*String) ForEach(ints []string, provided func(string))](#String.ForEach)
  * [func (*String) From(ints []string, mapFn ...func(string) string) []string](#String.From)
  * [func (*String) Head(ints []string) (string, bool)](#String.Head)
  * [func (*String) Includes(ints []string, e string) bool](#String.Includes)
  * [func (*String) IndexOf(ints []string, e string, fromIndex ...int) int](#String.IndexOf)
  * [func (*String) Initial(ints []string) []string](#String.Initial)
  * [func (i *String) Intersection(vs ...[]string) []string](#String.Intersection)
  * [func (*String) Join(ints []string, sep ...string) string](#String.Join)
  * [func (*String) Last(ints []string) (string, bool)](#String.Last)
  * [func (*String) LastIndexOf(ints []string, e string, fromIndex ...int) int](#String.LastIndexOf)
  * [func (*String) Map(ints []string, m func(string) string) []string](#String.Map)
  * [func (*String) Max(x, y string) string](#String.Max)
  * [func (*String) Min(x, y string) string](#String.Min)
  * [func (*String) Nth(ints []string, n ...int) (string, bool)](#String.Nth)
  * [func (*String) Of(ints ...string) []string](#String.Of)
  * [func (*String) Pop(ints []string) (string, bool, []string)](#String.Pop)
  * [func (i *String) Pull(ints []string, vs ...string) []string](#String.Pull)
  * [func (*String) Push(ints []string, es ...string) (int, []string)](#String.Push)
  * [func (*String) Reduce(ints []string, accum func(string, string) string, init ...string) string](#String.Reduce)
  * [func (*String) ReduceRight(ints []string, accum func(string, string) string, init ...string) string](#String.ReduceRight)
  * [func (*String) Remove(ints []string, predict func(string) bool) []string](#String.Remove)
  * [func (*String) Reverse(ints []string) []string](#String.Reverse)
  * [func (*String) Shift(ints []string) (string, bool, []string)](#String.Shift)
  * [func (*String) Shuffle(ints []string)](#String.Shuffle)
  * [func (*String) Slice(ints []string, startEndIndex ...int) []string](#String.Slice)
  * [func (*String) Some(ints []string, predict func(string) bool) bool](#String.Some)
  * [func (*String) Sort(ints []string, less ...func(int, int) bool)](#String.Sort)
  * [func (i *String) String(ints []string) string](#String.String)
  * [func (*String) Tail(ints []string) []string](#String.Tail)
  * [func (*String) Take(ints []string, n ...int) []string](#String.Take)
  * [func (i *String) Union(ints ...[]string) []string](#String.Union)
  * [func (i *String) Uniq(ints []string) []string](#String.Uniq)
  * [func (*String) Unshift(ints []string, es ...string) (int, []string)](#String.Unshift)
  * [func (i *String) Without(ints []string, vs ...string) []string](#String.Without)
* [type StringHelper](#StringHelper)


#### <a name="pkg-files">Package files</a>
[float32.go](./float32.go) [float64.go](./float64.go) [goji.go](./goji.go) [int.go](./int.go) [int32.go](./int32.go) [int64.go](./int64.go) [string.go](./string.go) 





## <a name="Concat">func</a> [Concat](./goji.go?s=2415:2428#L104)
``` go
func Concat()
```
Concat merges two slices.



## <a name="Difference">func</a> [Difference](./goji.go?s=5696:5791#L253)
``` go
func Difference(lenI, lenJ int, comparator func(int, int) bool, handlers ...func(...int)) []int
```
Difference returns difference of array/slice I's values not included in the other given array J.
lenI, lenJ is the length of array/slice I, J.
comparator compares values from array/slice I, J.



## <a name="Every">func</a> [Every](./goji.go?s=2610:2656#L110)
``` go
func Every(n int, predict func(int) bool) bool
```
Every tests whether all elements pass the test implemented by predict.
n is the number of elements. j is the index of array/slice.



## <a name="Fill">func</a> [Fill](./goji.go?s=4300:4354#L179)
``` go
func Fill(fill func(int), n int, startEndIndex ...int)
```
Fill fills all elements from a start index to an end index with a static value.
the end index is not included.
n is the number of elements. start and end index are optional, default to zero and n.



## <a name="Filter">func</a> [Filter](./goji.go?s=1434:1508#L62)
``` go
func Filter(n int, predict func(int) bool, handlers ...func(...int)) []int
```
Filter filters elements, n is the number of elements.
predict check element with index i and returns true or false.
handlers receive accumulated indexes idx from predict, if any.



## <a name="FindIndex">func</a> [FindIndex](./goji.go?s=4839:4906#L207)
``` go
func FindIndex(n int, testing func(int) bool, fromIndex ...int) int
```
FindIndex returns the index of the first element that satisfies the provided testing function.
Otherwise -1 is returned. n is the number of elements.



## <a name="FindLastIndex">func</a> [FindLastIndex](./goji.go?s=5192:5263#L229)
``` go
func FindLastIndex(n int, testing func(int) bool, fromIndex ...int) int
```
FindLastIndex iterates over elements from right to left.



## <a name="FloatToString">func</a> [FloatToString](./float32.go?s=286:323#L16)
``` go
func FloatToString(fv float64) string
```
FloatToString converts float to string.



## <a name="FloatToStringWithPrec">func</a> [FloatToStringWithPrec](./float32.go?s=135:190#L11)
``` go
func FloatToStringWithPrec(fv float64, prec int) string
```
FloatToStringWithPrec converts float to string with precision.



## <a name="Includes">func</a> [Includes](./goji.go?s=1882:1928#L80)
``` go
func Includes(n int, comp func(int) bool) bool
```
Includes determines whether an array/slice includes a certain element, returning true or false.
n is the number of elements. comp determines whether the given element with index j is found.



## <a name="Intersection">func</a> [Intersection](./goji.go?s=6161:6258#L275)
``` go
func Intersection(lenI, lenJ int, comparator func(int, int) bool, handlers ...func(...int)) []int
```
Intersection returns values that are included in both given array/slice.



## <a name="Join">func</a> [Join](./goji.go?s=3880:3941#L163)
``` go
func Join(join func(int) string, n int, sep ...string) string
```
Join joins all elements of array/slice into a string.
The separator string sep is placed between elements in the resulting string.
n is the number of elements. func join receive element index i and return a string.



## <a name="Map">func</a> [Map](./goji.go?s=2124:2152#L91)
``` go
func Map(n int, m func(int))
```
Map calls a provided function m on every element.
n is the number of elements. m map element with index i.



## <a name="Max">func</a> [Max](./goji.go?s=278:300#L20)
``` go
func Max(x, y int) int
```
Max returns the larger of x or y.



## <a name="Min">func</a> [Min](./goji.go?s=177:199#L12)
``` go
func Min(x, y int) int
```
Min returns the smaller of x or y.



## <a name="Range">func</a> [Range](./goji.go?s=430:460#L28)
``` go
func Range(end int) []struct{}
```
Range creates a range of numbers progressing from zero up to, but not including end.



## <a name="Reverse">func</a> [Reverse](./goji.go?s=1132:1172#L51)
``` go
func Reverse(n int, swap func(int, int))
```
Reverse reverses an array/slice, n is the number of elements.
swap swaps the elements with indexes i and j.



## <a name="Shuffle">func</a> [Shuffle](./goji.go?s=2317:2357#L99)
``` go
func Shuffle(n int, swap func(int, int))
```
Shuffle randomizes the order of elements.
n is the number of elements. swap swaps the elements with indexes i and j.



## <a name="Slice">func</a> [Slice](./goji.go?s=3271:3332#L135)
``` go
func Slice(slice func(int, int), n int, startEndIndex ...int)
```
Slice returns a portion of array/slice selected from start to end (end not included).
n is the number of elements. start and end indexes are optional, default to zero and n.
start and end indexes are passed to slice func.



## <a name="Some">func</a> [Some](./goji.go?s=2898:2943#L122)
``` go
func Some(n int, predict func(int) bool) bool
```
Some tests whether at least one element pass the test implemented by predict.
n is the number of elements. j is the index of array/slice.



## <a name="Sort">func</a> [Sort](./goji.go?s=898:961#L45)
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







### <a name="NewInt">func</a> [NewInt](./int.go?s=1925:1943#L60)
``` go
func NewInt() *Int
```
NewInt creates a new Int instance.





### <a name="Int.Concat">func</a> (\*Int) [Concat](./int.go?s=5203:5242#L232)
``` go
func (*Int) Concat(ints ...[]int) []int
```
Concat merges two or more slices.




### <a name="Int.Difference">func</a> (\*Int) [Difference](./int.go?s=11370:11425#L536)
``` go
func (i *Int) Difference(ints []int, vs ...[]int) []int
```
Difference creates a slice of values not included in the other given slices.




### <a name="Int.Drop">func</a> (\*Int) [Drop](./int.go?s=10006:10050#L458)
``` go
func (*Int) Drop(ints []int, n ...int) []int
```
Drop creates a slice with n elements dropped from the beginning.




### <a name="Int.Equal">func</a> (\*Int) [Equal](./int.go?s=13440:13476#L652)
``` go
func (*Int) Equal(s1, s2 []int) bool
```
Equal performs a deep comparison between two slice values.




### <a name="Int.Every">func</a> (\*Int) [Every](./int.go?s=5403:5461#L243)
``` go
func (*Int) Every(ints []int, predict func(int) bool) bool
```
Every tests whether all elements in the slice pass the predict implemented.




### <a name="Int.Fill">func</a> (\*Int) [Fill](./int.go?s=5931:5992#L266)
``` go
func (*Int) Fill(ints []int, value int, startEndIndex ...int)
```
Fill fills all the elements of a slice from a start index to an end index with a static value.
the end index is not included.




### <a name="Int.Filter">func</a> (\*Int) [Filter](./int.go?s=2082:2142#L68)
``` go
func (*Int) Filter(ints []int, predict func(int) bool) []int
```
Filter creates a new slice with all elements that pass the predict implemented.




### <a name="Int.Find">func</a> (\*Int) [Find](./int.go?s=2392:2456#L82)
``` go
func (*Int) Find(ints []int, predict func(int) bool) (int, bool)
```
Find returns the value of the first element in the slice that satisfies the provided predict function.
otherwise false is returned.




### <a name="Int.FindIndex">func</a> (\*Int) [FindIndex](./int.go?s=2699:2778#L94)
``` go
func (*Int) FindIndex(ints []int, predict func(int) bool, fromIndex ...int) int
```
FindIndex returns the index of the first element in the slice that satisfies the provided predict function.
otherwise -1 is returned.




### <a name="Int.FindLastIndex">func</a> (\*Int) [FindLastIndex](./int.go?s=3123:3206#L117)
``` go
func (*Int) FindLastIndex(ints []int, predict func(int) bool, fromIndex ...int) int
```
FindLastIndex  is like FindIndex except that it iterates over elements from right to left.




### <a name="Int.ForEach">func</a> (\*Int) [ForEach](./int.go?s=3528:3579#L140)
``` go
func (*Int) ForEach(ints []int, provided func(int))
```
ForEach executes a provided function once for each slice element.




### <a name="Int.From">func</a> (\*Int) [From](./int.go?s=4724:4782#L204)
``` go
func (*Int) From(ints []int, mapFn ...func(int) int) []int
```
From creates a new, shallow-copied slice.




### <a name="Int.Head">func</a> (\*Int) [Head](./int.go?s=10284:10324#L475)
``` go
func (*Int) Head(ints []int) (int, bool)
```
Head gets the first element of slice.




### <a name="Int.Includes">func</a> (\*Int) [Includes](./int.go?s=3720:3764#L147)
``` go
func (*Int) Includes(ints []int, e int) bool
```
Includes determines whether a slice includes a certain element, returning true or false.




### <a name="Int.IndexOf">func</a> (\*Int) [IndexOf](./int.go?s=3962:4022#L158)
``` go
func (*Int) IndexOf(ints []int, e int, fromIndex ...int) int
```
IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.




### <a name="Int.Initial">func</a> (\*Int) [Initial](./int.go?s=10445:10482#L484)
``` go
func (*Int) Initial(ints []int) []int
```
Initial gets all but the last element of slice.




### <a name="Int.Intersection">func</a> (\*Int) [Intersection](./int.go?s=11745:11790#L556)
``` go
func (i *Int) Intersection(vs ...[]int) []int
```
Intersection creates a slice of unique values that are included in all given slices.




### <a name="Int.Join">func</a> (\*Int) [Join](./int.go?s=6398:6448#L294)
``` go
func (*Int) Join(ints []int, sep ...string) string
```
Join joins all elements of a slice into a string.




### <a name="Int.Last">func</a> (\*Int) [Last](./int.go?s=10596:10636#L494)
``` go
func (*Int) Last(ints []int) (int, bool)
```
Last gets the last element of slice.




### <a name="Int.LastIndexOf">func</a> (\*Int) [LastIndexOf](./int.go?s=4377:4441#L181)
``` go
func (*Int) LastIndexOf(ints []int, e int, fromIndex ...int) int
```
LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.




### <a name="Int.Map">func</a> (\*Int) [Map](./int.go?s=6731:6781#L309)
``` go
func (*Int) Map(ints []int, m func(int) int) []int
```
Map creates a new slice with the results of calling a provided function on every element in the calling slice.




### <a name="Int.Nth">func</a> (\*Int) [Nth](./int.go?s=10824:10873#L505)
``` go
func (*Int) Nth(ints []int, n ...int) (int, bool)
```
Nth gets the element at index n of slice.
if n is negative, the nth element from the end is returned.




### <a name="Int.Of">func</a> (\*Int) [Of](./int.go?s=5047:5080#L222)
``` go
func (*Int) Of(ints ...int) []int
```
Of creates a new slice instance with a variable number of arguments.




### <a name="Int.Pop">func</a> (\*Int) [Pop](./int.go?s=6942:6988#L320)
``` go
func (*Int) Pop(ints []int) (int, bool, []int)
```
Pop removes the last element from a slice and returns that element.




### <a name="Int.Pull">func</a> (\*Int) [Pull](./int.go?s=11125:11172#L525)
``` go
func (i *Int) Pull(ints []int, vs ...int) []int
```
Pull removes all given values from slice.




### <a name="Int.Push">func</a> (\*Int) [Push](./int.go?s=7186:7238#L329)
``` go
func (*Int) Push(ints []int, es ...int) (int, []int)
```
Push adds one or more elements to the end of a slice and returns the new length of the slice.




### <a name="Int.Reduce">func</a> (\*Int) [Reduce](./int.go?s=7462:7535#L341)
``` go
func (*Int) Reduce(ints []int, accum func(int, int) int, init ...int) int
```
Reduce applies a function against each element in the slice to reduce it to a single value.




### <a name="Int.ReduceRight">func</a> (\*Int) [ReduceRight](./int.go?s=7770:7848#L353)
``` go
func (*Int) ReduceRight(ints []int, accum func(int, int) int, init ...int) int
```
ReduceRight applies a function against each value of the slice (from right-to-left) to reduce it to a single value.




### <a name="Int.Remove">func</a> (\*Int) [Remove](./int.go?s=12320:12380#L589)
``` go
func (*Int) Remove(ints []int, predict func(int) bool) []int
```
Remove removes all elements from slice that predicate returns truthy for.




### <a name="Int.Reverse">func</a> (\*Int) [Reverse](./int.go?s=8018:8055#L366)
``` go
func (*Int) Reverse(ints []int) []int
```
Reverse reverses a slice.




### <a name="Int.Shift">func</a> (\*Int) [Shift](./int.go?s=8251:8299#L376)
``` go
func (*Int) Shift(ints []int) (int, bool, []int)
```
Shift removes the first element from a slice and returns that element.




### <a name="Int.Shuffle">func</a> (\*Int) [Shuffle](./int.go?s=9679:9710#L447)
``` go
func (*Int) Shuffle(ints []int)
```
Shuffle randomizes the order of elements.




### <a name="Int.Slice">func</a> (\*Int) [Slice](./int.go?s=8881:8938#L403)
``` go
func (*Int) Slice(ints []int, startEndIndex ...int) []int
```
Slice returns a portion of a slice selected from start to end (end not included).
start and end indexes are optional, default to zero and length of the slice.




### <a name="Int.Some">func</a> (\*Int) [Some](./int.go?s=5645:5702#L254)
``` go
func (*Int) Some(ints []int, predict func(int) bool) bool
```
Some tests whether at least one element in the slice pass the predict implemented.




### <a name="Int.Sort">func</a> (\*Int) [Sort](./int.go?s=9348:9405#L435)
``` go
func (*Int) Sort(ints []int, less ...func(int, int) bool)
```
Sort sorts elements.




### <a name="Int.String">func</a> (\*Int) [String](./int.go?s=9872:9911#L453)
``` go
func (i *Int) String(ints []int) string
```
String returns a string representing the specified slice.




### <a name="Int.Tail">func</a> (\*Int) [Tail](./int.go?s=12542:12576#L602)
``` go
func (*Int) Tail(ints []int) []int
```
Tail gets all but the first element of slice.




### <a name="Int.Take">func</a> (\*Int) [Take](./int.go?s=12714:12758#L612)
``` go
func (*Int) Take(ints []int, n ...int) []int
```
Take creates a slice with n elements taken from the beginning.




### <a name="Int.Union">func</a> (\*Int) [Union](./int.go?s=12982:13022#L628)
``` go
func (i *Int) Union(ints ...[]int) []int
```
Union creates a slice of unique values, in order, from all given slices.




### <a name="Int.Uniq">func</a> (\*Int) [Uniq](./int.go?s=13225:13261#L641)
``` go
func (i *Int) Uniq(ints []int) []int
```
Uniq creates a duplicate-free version of a slice.




### <a name="Int.Unshift">func</a> (\*Int) [Unshift](./int.go?s=8502:8557#L385)
``` go
func (*Int) Unshift(ints []int, es ...int) (int, []int)
```
Unshift adds one or more elements to the beginning of a slice and returns the new length of the slice.




### <a name="Int.Without">func</a> (\*Int) [Without](./int.go?s=12159:12209#L584)
``` go
func (i *Int) Without(ints []int, vs ...int) []int
```
Without creates a slice excluding all given values.



## <a name="IntHelper">type</a> [IntHelper](./int.go?s=150:1885#L14)
``` go
type IntHelper interface {
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
```
IntHelper is the Int interface.










## <a name="String">type</a> [String](./string.go?s=112:132#L12)
``` go
type String struct{}
```
String for string type slice.







### <a name="NewString">func</a> [NewString](./string.go?s=2312:2336#L63)
``` go
func NewString() *String
```
NewString creates a new String instance.





### <a name="String.CamelCase">func</a> (\*String) [CamelCase](./string.go?s=14813:14878#L674)
``` go
func (i *String) CamelCase(input string, pascalCase ...bool) string
```
CamelCase convert dash/dot/underscore/space separated string to camelCase
pascalCase define whether to uppercase the first character




### <a name="String.Concat">func</a> (\*String) [Concat](./string.go?s=5966:6014#L250)
``` go
func (*String) Concat(ints ...[]string) []string
```
Concat merges two or more slices.




### <a name="String.Difference">func</a> (\*String) [Difference](./string.go?s=12355:12422#L548)
``` go
func (i *String) Difference(ints []string, vs ...[]string) []string
```
Difference creates a slice of values not included in the other given slices.




### <a name="String.Drop">func</a> (\*String) [Drop](./string.go?s=10916:10969#L470)
``` go
func (*String) Drop(ints []string, n ...int) []string
```
Drop creates a slice with n elements dropped from the beginning.




### <a name="String.Equal">func</a> (\*String) [Equal](./string.go?s=14527:14569#L664)
``` go
func (*String) Equal(s1, s2 []string) bool
```
Equal performs a deep comparison between two slice values.




### <a name="String.Every">func</a> (\*String) [Every](./string.go?s=6178:6245#L261)
``` go
func (*String) Every(ints []string, predict func(string) bool) bool
```
Every tests whether all elements in the slice pass the predict implemented.




### <a name="String.Fill">func</a> (\*String) [Fill](./string.go?s=6724:6794#L284)
``` go
func (*String) Fill(ints []string, value string, startEndIndex ...int)
```
Fill fills all the elements of a slice from a start index to an end index with a static value.
the end index is not included.




### <a name="String.Filter">func</a> (\*String) [Filter](./string.go?s=2712:2784#L86)
``` go
func (*String) Filter(ints []string, predict func(string) bool) []string
```
Filter creates a new slice with all elements that pass the predict implemented.




### <a name="String.Find">func</a> (\*String) [Find](./string.go?s=3037:3113#L100)
``` go
func (*String) Find(ints []string, predict func(string) bool) (string, bool)
```
Find returns the value of the first element in the slice that satisfies the provided predict function.
otherwise false is returned.




### <a name="String.FindIndex">func</a> (\*String) [FindIndex](./string.go?s=3357:3445#L112)
``` go
func (*String) FindIndex(ints []string, predict func(string) bool, fromIndex ...int) int
```
FindIndex returns the index of the first element in the slice that satisfies the provided predict function.
otherwise -1 is returned.




### <a name="String.FindLastIndex">func</a> (\*String) [FindLastIndex](./string.go?s=3793:3885#L135)
``` go
func (*String) FindLastIndex(ints []string, predict func(string) bool, fromIndex ...int) int
```
FindLastIndex  is like FindIndex except that it iterates over elements from right to left.




### <a name="String.ForEach">func</a> (\*String) [ForEach](./string.go?s=4213:4273#L158)
``` go
func (*String) ForEach(ints []string, provided func(string))
```
ForEach executes a provided function once for each slice element.




### <a name="String.From">func</a> (\*String) [From](./string.go?s=5451:5524#L222)
``` go
func (*String) From(ints []string, mapFn ...func(string) string) []string
```
From creates a new, shallow-copied slice.




### <a name="String.Head">func</a> (\*String) [Head](./string.go?s=11212:11261#L487)
``` go
func (*String) Head(ints []string) (string, bool)
```
Head gets the first element of slice.




### <a name="String.Includes">func</a> (\*String) [Includes](./string.go?s=4414:4467#L165)
``` go
func (*String) Includes(ints []string, e string) bool
```
Includes determines whether a slice includes a certain element, returning true or false.




### <a name="String.IndexOf">func</a> (\*String) [IndexOf](./string.go?s=4665:4734#L176)
``` go
func (*String) IndexOf(ints []string, e string, fromIndex ...int) int
```
IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.




### <a name="String.Initial">func</a> (\*String) [Initial](./string.go?s=11383:11429#L496)
``` go
func (*String) Initial(ints []string) []string
```
Initial gets all but the last element of slice.




### <a name="String.Intersection">func</a> (\*String) [Intersection](./string.go?s=12745:12799#L568)
``` go
func (i *String) Intersection(vs ...[]string) []string
```
Intersection creates a slice of unique values that are included in all given slices.




### <a name="String.Join">func</a> (\*String) [Join](./string.go?s=7210:7266#L312)
``` go
func (*String) Join(ints []string, sep ...string) string
```
Join joins all elements of a slice into a string.




### <a name="String.Last">func</a> (\*String) [Last](./string.go?s=11546:11595#L506)
``` go
func (*String) Last(ints []string) (string, bool)
```
Last gets the last element of slice.




### <a name="String.LastIndexOf">func</a> (\*String) [LastIndexOf](./string.go?s=5092:5165#L199)
``` go
func (*String) LastIndexOf(ints []string, e string, fromIndex ...int) int
```
LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.




### <a name="String.Map">func</a> (\*String) [Map](./string.go?s=7464:7529#L321)
``` go
func (*String) Map(ints []string, m func(string) string) []string
```
Map creates a new slice with the results of calling a provided function on every element in the calling slice.




### <a name="String.Max">func</a> (\*String) [Max](./string.go?s=2549:2587#L78)
``` go
func (*String) Max(x, y string) string
```
Max returns the larger of x or y.




### <a name="String.Min">func</a> (\*String) [Min](./string.go?s=2432:2470#L70)
``` go
func (*String) Min(x, y string) string
```
Min returns the smaller of x or y.




### <a name="String.Nth">func</a> (\*String) [Nth](./string.go?s=11784:11842#L517)
``` go
func (*String) Nth(ints []string, n ...int) (string, bool)
```
Nth gets the element at index n of slice.
if n is negative, the nth element from the end is returned.




### <a name="String.Of">func</a> (\*String) [Of](./string.go?s=5798:5840#L240)
``` go
func (*String) Of(ints ...string) []string
```
Of creates a new slice instance with a variable number of arguments.




### <a name="String.Pop">func</a> (\*String) [Pop](./string.go?s=7693:7751#L332)
``` go
func (*String) Pop(ints []string) (string, bool, []string)
```
Pop removes the last element from a slice and returns that element.




### <a name="String.Pull">func</a> (\*String) [Pull](./string.go?s=12095:12154#L537)
``` go
func (i *String) Pull(ints []string, vs ...string) []string
```
Pull removes all given values from slice.




### <a name="String.Push">func</a> (\*String) [Push](./string.go?s=7953:8017#L341)
``` go
func (*String) Push(ints []string, es ...string) (int, []string)
```
Push adds one or more elements to the end of a slice and returns the new length of the slice.




### <a name="String.Reduce">func</a> (\*String) [Reduce](./string.go?s=8247:8341#L353)
``` go
func (*String) Reduce(ints []string, accum func(string, string) string, init ...string) string
```
Reduce applies a function against each element in the slice to reduce it to a single value.




### <a name="String.ReduceRight">func</a> (\*String) [ReduceRight](./string.go?s=8579:8678#L365)
``` go
func (*String) ReduceRight(ints []string, accum func(string, string) string, init ...string) string
```
ReduceRight applies a function against each value of the slice (from right-to-left) to reduce it to a single value.




### <a name="String.Remove">func</a> (\*String) [Remove](./string.go?s=13344:13416#L601)
``` go
func (*String) Remove(ints []string, predict func(string) bool) []string
```
Remove removes all elements from slice that predicate returns truthy for.




### <a name="String.Reverse">func</a> (\*String) [Reverse](./string.go?s=8851:8897#L378)
``` go
func (*String) Reverse(ints []string) []string
```
Reverse reverses a slice.




### <a name="String.Shift">func</a> (\*String) [Shift](./string.go?s=9093:9153#L388)
``` go
func (*String) Shift(ints []string) (string, bool, []string)
```
Shift removes the first element from a slice and returns that element.




### <a name="String.Shuffle">func</a> (\*String) [Shuffle](./string.go?s=10577:10614#L459)
``` go
func (*String) Shuffle(ints []string)
```
Shuffle randomizes the order of elements.




### <a name="String.Slice">func</a> (\*String) [Slice](./string.go?s=9754:9820#L415)
``` go
func (*String) Slice(ints []string, startEndIndex ...int) []string
```
Slice returns a portion of a slice selected from start to end (end not included).
start and end indexes are optional, default to zero and length of the slice.




### <a name="String.Some">func</a> (\*String) [Some](./string.go?s=6429:6495#L272)
``` go
func (*String) Some(ints []string, predict func(string) bool) bool
```
Some tests whether at least one element in the slice pass the predict implemented.




### <a name="String.Sort">func</a> (\*String) [Sort](./string.go?s=10240:10303#L447)
``` go
func (*String) Sort(ints []string, less ...func(int, int) bool)
```
Sort sorts elements.




### <a name="String.String">func</a> (\*String) [String](./string.go?s=10776:10821#L465)
``` go
func (i *String) String(ints []string) string
```
String returns a string representing the specified slice.




### <a name="String.Tail">func</a> (\*String) [Tail](./string.go?s=13581:13624#L614)
``` go
func (*String) Tail(ints []string) []string
```
Tail gets all but the first element of slice.




### <a name="String.Take">func</a> (\*String) [Take](./string.go?s=13765:13818#L624)
``` go
func (*String) Take(ints []string, n ...int) []string
```
Take creates a slice with n elements taken from the beginning.




### <a name="String.Union">func</a> (\*String) [Union](./string.go?s=14045:14094#L640)
``` go
func (i *String) Union(ints ...[]string) []string
```
Union creates a slice of unique values, in order, from all given slices.




### <a name="String.Uniq">func</a> (\*String) [Uniq](./string.go?s=14300:14345#L653)
``` go
func (i *String) Uniq(ints []string) []string
```
Uniq creates a duplicate-free version of a slice.




### <a name="String.Unshift">func</a> (\*String) [Unshift](./string.go?s=9360:9427#L397)
``` go
func (*String) Unshift(ints []string, es ...string) (int, []string)
```
Unshift adds one or more elements to the beginning of a slice and returns the new length of the slice.




### <a name="String.Without">func</a> (\*String) [Without](./string.go?s=13171:13233#L596)
``` go
func (i *String) Without(ints []string, vs ...string) []string
```
Without creates a slice excluding all given values.




## <a name="StringHelper">type</a> [StringHelper](./string.go?s=175:2266#L15)
``` go
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
```
StringHelper is the String interface.








