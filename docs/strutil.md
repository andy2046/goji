

# strutil
`import "github.com/andy2046/goji/pkg/strutil"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package strutil provides functional utils for string.




## <a name="pkg-index">Index</a>
* [type IString](#IString)
* [type String](#String)
  * [func NewString() *String](#NewString)
  * [func (*String) CamelCase(input string, pascalCase ...bool) string](#String.CamelCase)
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
  * [func (*String) KebabCase(input string, separateDigit ...bool) string](#String.KebabCase)
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
  * [func (*String) SnakeCase(input string, separateDigit ...bool) string](#String.SnakeCase)
  * [func (*String) Some(ints []string, predict func(string) bool) bool](#String.Some)
  * [func (*String) Sort(ints []string, less ...func(int, int) bool)](#String.Sort)
  * [func (i *String) String(ints []string) string](#String.String)
  * [func (*String) Tail(ints []string) []string](#String.Tail)
  * [func (*String) Take(ints []string, n ...int) []string](#String.Take)
  * [func (*String) Titleize(input string) string](#String.Titleize)
  * [func (i *String) Union(ints ...[]string) []string](#String.Union)
  * [func (i *String) Uniq(ints []string) []string](#String.Uniq)
  * [func (*String) Unshift(ints []string, es ...string) (int, []string)](#String.Unshift)
  * [func (i *String) Without(ints []string, vs ...string) []string](#String.Without)


#### <a name="pkg-files">Package files</a>
[string.go](/src/github.com/andy2046/goji/pkg/strutil/string.go) 






## <a name="IString">type</a> [IString](/src/target/string.go?s=261:2347#L18)
``` go
type IString interface {
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
IString is the String interface.










## <a name="String">type</a> [String](/src/target/string.go?s=203:223#L15)
``` go
type String struct{}
```
String for string type slice.







### <a name="NewString">func</a> [NewString](/src/target/string.go?s=2393:2417#L66)
``` go
func NewString() *String
```
NewString creates a new String instance.





### <a name="String.CamelCase">func</a> (\*String) [CamelCase](/src/target/string.go?s=14929:14994#L677)
``` go
func (*String) CamelCase(input string, pascalCase ...bool) string
```
CamelCase convert dash/dot/underscore/space separated string to camelCase
pascalCase define whether to uppercase the first character




### <a name="String.Concat">func</a> (\*String) [Concat](/src/target/string.go?s=6058:6106#L253)
``` go
func (*String) Concat(ints ...[]string) []string
```
Concat merges two or more slices.




### <a name="String.Difference">func</a> (\*String) [Difference](/src/target/string.go?s=12471:12538#L551)
``` go
func (i *String) Difference(ints []string, vs ...[]string) []string
```
Difference creates a slice of values not included in the other given slices.




### <a name="String.Drop">func</a> (\*String) [Drop](/src/target/string.go?s=11032:11085#L473)
``` go
func (*String) Drop(ints []string, n ...int) []string
```
Drop creates a slice with n elements dropped from the beginning.




### <a name="String.Equal">func</a> (\*String) [Equal](/src/target/string.go?s=14643:14685#L667)
``` go
func (*String) Equal(s1, s2 []string) bool
```
Equal performs a deep comparison between two slice values.




### <a name="String.Every">func</a> (\*String) [Every](/src/target/string.go?s=6270:6337#L264)
``` go
func (*String) Every(ints []string, predict func(string) bool) bool
```
Every tests whether all elements in the slice pass the predict implemented.




### <a name="String.Fill">func</a> (\*String) [Fill](/src/target/string.go?s=6816:6886#L287)
``` go
func (*String) Fill(ints []string, value string, startEndIndex ...int)
```
Fill fills all the elements of a slice from a start index to an end index with a static value.
the end index is not included.




### <a name="String.Filter">func</a> (\*String) [Filter](/src/target/string.go?s=2788:2860#L89)
``` go
func (*String) Filter(ints []string, predict func(string) bool) []string
```
Filter creates a new slice with all elements that pass the predict implemented.




### <a name="String.Find">func</a> (\*String) [Find](/src/target/string.go?s=3113:3189#L103)
``` go
func (*String) Find(ints []string, predict func(string) bool) (string, bool)
```
Find returns the value of the first element in the slice that satisfies the provided predict function.
otherwise false is returned.




### <a name="String.FindIndex">func</a> (\*String) [FindIndex](/src/target/string.go?s=3433:3521#L115)
``` go
func (*String) FindIndex(ints []string, predict func(string) bool, fromIndex ...int) int
```
FindIndex returns the index of the first element in the slice that satisfies the provided predict function.
otherwise -1 is returned.




### <a name="String.FindLastIndex">func</a> (\*String) [FindLastIndex](/src/target/string.go?s=3873:3965#L138)
``` go
func (*String) FindLastIndex(ints []string, predict func(string) bool, fromIndex ...int) int
```
FindLastIndex  is like FindIndex except that it iterates over elements from right to left.




### <a name="String.ForEach">func</a> (\*String) [ForEach](/src/target/string.go?s=4297:4357#L161)
``` go
func (*String) ForEach(ints []string, provided func(string))
```
ForEach executes a provided function once for each slice element.




### <a name="String.From">func</a> (\*String) [From](/src/target/string.go?s=5543:5616#L225)
``` go
func (*String) From(ints []string, mapFn ...func(string) string) []string
```
From creates a new, shallow-copied slice.




### <a name="String.Head">func</a> (\*String) [Head](/src/target/string.go?s=11328:11377#L490)
``` go
func (*String) Head(ints []string) (string, bool)
```
Head gets the first element of slice.




### <a name="String.Includes">func</a> (\*String) [Includes](/src/target/string.go?s=4498:4551#L168)
``` go
func (*String) Includes(ints []string, e string) bool
```
Includes determines whether a slice includes a certain element, returning true or false.




### <a name="String.IndexOf">func</a> (\*String) [IndexOf](/src/target/string.go?s=4749:4818#L179)
``` go
func (*String) IndexOf(ints []string, e string, fromIndex ...int) int
```
IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.




### <a name="String.Initial">func</a> (\*String) [Initial](/src/target/string.go?s=11499:11545#L499)
``` go
func (*String) Initial(ints []string) []string
```
Initial gets all but the last element of slice.




### <a name="String.Intersection">func</a> (\*String) [Intersection](/src/target/string.go?s=12861:12915#L571)
``` go
func (i *String) Intersection(vs ...[]string) []string
```
Intersection creates a slice of unique values that are included in all given slices.




### <a name="String.Join">func</a> (\*String) [Join](/src/target/string.go?s=7310:7366#L315)
``` go
func (*String) Join(ints []string, sep ...string) string
```
Join joins all elements of a slice into a string.




### <a name="String.KebabCase">func</a> (\*String) [KebabCase](/src/target/string.go?s=16826:16894#L788)
``` go
func (*String) KebabCase(input string, separateDigit ...bool) string
```
KebabCase convert string to kebab-case
separateDigit define whether to separate digit from letter




### <a name="String.Last">func</a> (\*String) [Last](/src/target/string.go?s=11662:11711#L509)
``` go
func (*String) Last(ints []string) (string, bool)
```
Last gets the last element of slice.




### <a name="String.LastIndexOf">func</a> (\*String) [LastIndexOf](/src/target/string.go?s=5180:5253#L202)
``` go
func (*String) LastIndexOf(ints []string, e string, fromIndex ...int) int
```
LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.




### <a name="String.Map">func</a> (\*String) [Map](/src/target/string.go?s=7564:7629#L324)
``` go
func (*String) Map(ints []string, m func(string) string) []string
```
Map creates a new slice with the results of calling a provided function on every element in the calling slice.




### <a name="String.Max">func</a> (\*String) [Max](/src/target/string.go?s=2625:2663#L81)
``` go
func (*String) Max(x, y string) string
```
Max returns the larger of x or y.




### <a name="String.Min">func</a> (\*String) [Min](/src/target/string.go?s=2508:2546#L73)
``` go
func (*String) Min(x, y string) string
```
Min returns the smaller of x or y.




### <a name="String.Nth">func</a> (\*String) [Nth](/src/target/string.go?s=11900:11958#L520)
``` go
func (*String) Nth(ints []string, n ...int) (string, bool)
```
Nth gets the element at index n of slice.
if n is negative, the nth element from the end is returned.




### <a name="String.Of">func</a> (\*String) [Of](/src/target/string.go?s=5890:5932#L243)
``` go
func (*String) Of(ints ...string) []string
```
Of creates a new slice instance with a variable number of arguments.




### <a name="String.Pop">func</a> (\*String) [Pop](/src/target/string.go?s=7793:7851#L335)
``` go
func (*String) Pop(ints []string) (string, bool, []string)
```
Pop removes the last element from a slice and returns that element.




### <a name="String.Pull">func</a> (\*String) [Pull](/src/target/string.go?s=12211:12270#L540)
``` go
func (i *String) Pull(ints []string, vs ...string) []string
```
Pull removes all given values from slice.




### <a name="String.Push">func</a> (\*String) [Push](/src/target/string.go?s=8053:8117#L344)
``` go
func (*String) Push(ints []string, es ...string) (int, []string)
```
Push adds one or more elements to the end of a slice and returns the new length of the slice.




### <a name="String.Reduce">func</a> (\*String) [Reduce](/src/target/string.go?s=8347:8441#L356)
``` go
func (*String) Reduce(ints []string, accum func(string, string) string, init ...string) string
```
Reduce applies a function against each element in the slice to reduce it to a single value.




### <a name="String.ReduceRight">func</a> (\*String) [ReduceRight](/src/target/string.go?s=8679:8778#L368)
``` go
func (*String) ReduceRight(ints []string, accum func(string, string) string, init ...string) string
```
ReduceRight applies a function against each value of the slice (from right-to-left) to reduce it to a single value.




### <a name="String.Remove">func</a> (\*String) [Remove](/src/target/string.go?s=13460:13532#L604)
``` go
func (*String) Remove(ints []string, predict func(string) bool) []string
```
Remove removes all elements from slice that predicate returns truthy for.




### <a name="String.Reverse">func</a> (\*String) [Reverse](/src/target/string.go?s=8951:8997#L381)
``` go
func (*String) Reverse(ints []string) []string
```
Reverse reverses a slice.




### <a name="String.Shift">func</a> (\*String) [Shift](/src/target/string.go?s=9193:9253#L391)
``` go
func (*String) Shift(ints []string) (string, bool, []string)
```
Shift removes the first element from a slice and returns that element.




### <a name="String.Shuffle">func</a> (\*String) [Shuffle](/src/target/string.go?s=10693:10730#L462)
``` go
func (*String) Shuffle(ints []string)
```
Shuffle randomizes the order of elements.




### <a name="String.Slice">func</a> (\*String) [Slice](/src/target/string.go?s=9854:9920#L418)
``` go
func (*String) Slice(ints []string, startEndIndex ...int) []string
```
Slice returns a portion of a slice selected from start to end (end not included).
start and end indexes are optional, default to zero and length of the slice.




### <a name="String.SnakeCase">func</a> (\*String) [SnakeCase](/src/target/string.go?s=16048:16116#L740)
``` go
func (*String) SnakeCase(input string, separateDigit ...bool) string
```
SnakeCase convert string to snake_case
separateDigit define whether to separate digit from letter




### <a name="String.Some">func</a> (\*String) [Some](/src/target/string.go?s=6521:6587#L275)
``` go
func (*String) Some(ints []string, predict func(string) bool) bool
```
Some tests whether at least one element in the slice pass the predict implemented.




### <a name="String.Sort">func</a> (\*String) [Sort](/src/target/string.go?s=10348:10411#L450)
``` go
func (*String) Sort(ints []string, less ...func(int, int) bool)
```
Sort sorts elements.




### <a name="String.String">func</a> (\*String) [String](/src/target/string.go?s=10892:10937#L468)
``` go
func (i *String) String(ints []string) string
```
String returns a string representing the specified slice.




### <a name="String.Tail">func</a> (\*String) [Tail](/src/target/string.go?s=13697:13740#L617)
``` go
func (*String) Tail(ints []string) []string
```
Tail gets all but the first element of slice.




### <a name="String.Take">func</a> (\*String) [Take](/src/target/string.go?s=13881:13934#L627)
``` go
func (*String) Take(ints []string, n ...int) []string
```
Take creates a slice with n elements taken from the beginning.




### <a name="String.Titleize">func</a> (\*String) [Titleize](/src/target/string.go?s=15741:15785#L731)
``` go
func (*String) Titleize(input string) string
```
Titleize capitalize every word in string




### <a name="String.Union">func</a> (\*String) [Union](/src/target/string.go?s=14161:14210#L643)
``` go
func (i *String) Union(ints ...[]string) []string
```
Union creates a slice of unique values, in order, from all given slices.




### <a name="String.Uniq">func</a> (\*String) [Uniq](/src/target/string.go?s=14416:14461#L656)
``` go
func (i *String) Uniq(ints []string) []string
```
Uniq creates a duplicate-free version of a slice.




### <a name="String.Unshift">func</a> (\*String) [Unshift](/src/target/string.go?s=9460:9527#L400)
``` go
func (*String) Unshift(ints []string, es ...string) (int, []string)
```
Unshift adds one or more elements to the beginning of a slice and returns the new length of the slice.




### <a name="String.Without">func</a> (\*String) [Without](/src/target/string.go?s=13287:13349#L599)
``` go
func (i *String) Without(ints []string, vs ...string) []string
```
Without creates a slice excluding all given values.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
