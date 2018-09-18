

# intutil
`import "github.com/andy2046/goji/pkg/intutil"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package intutil provides functional utils for int.




## <a name="pkg-index">Index</a>
* [type IInt](#IInt)
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


#### <a name="pkg-files">Package files</a>
[int.go](/src/github.com/andy2046/goji/pkg/intutil/int.go) 






## <a name="IInt">type</a> [IInt](/src/target/int.go?s=233:1963#L17)
``` go
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
```
IInt is the Int interface.










## <a name="Int">type</a> [Int](/src/target/int.go?s=184:201#L14)
``` go
type Int struct{}
```
Int for int type slice.







### <a name="NewInt">func</a> [NewInt](/src/target/int.go?s=2003:2021#L63)
``` go
func NewInt() *Int
```
NewInt creates a new Int instance.





### <a name="Int.Concat">func</a> (\*Int) [Concat](/src/target/int.go?s=5292:5331#L235)
``` go
func (*Int) Concat(ints ...[]int) []int
```
Concat merges two or more slices.




### <a name="Int.Difference">func</a> (\*Int) [Difference](/src/target/int.go?s=11483:11538#L539)
``` go
func (i *Int) Difference(ints []int, vs ...[]int) []int
```
Difference creates a slice of values not included in the other given slices.




### <a name="Int.Drop">func</a> (\*Int) [Drop](/src/target/int.go?s=10119:10163#L461)
``` go
func (*Int) Drop(ints []int, n ...int) []int
```
Drop creates a slice with n elements dropped from the beginning.




### <a name="Int.Equal">func</a> (\*Int) [Equal](/src/target/int.go?s=13553:13589#L655)
``` go
func (*Int) Equal(s1, s2 []int) bool
```
Equal performs a deep comparison between two slice values.




### <a name="Int.Every">func</a> (\*Int) [Every](/src/target/int.go?s=5492:5550#L246)
``` go
func (*Int) Every(ints []int, predict func(int) bool) bool
```
Every tests whether all elements in the slice pass the predict implemented.




### <a name="Int.Fill">func</a> (\*Int) [Fill](/src/target/int.go?s=6020:6081#L269)
``` go
func (*Int) Fill(ints []int, value int, startEndIndex ...int)
```
Fill fills all the elements of a slice from a start index to an end index with a static value.
the end index is not included.




### <a name="Int.Filter">func</a> (\*Int) [Filter](/src/target/int.go?s=2155:2215#L71)
``` go
func (*Int) Filter(ints []int, predict func(int) bool) []int
```
Filter creates a new slice with all elements that pass the predict implemented.




### <a name="Int.Find">func</a> (\*Int) [Find](/src/target/int.go?s=2465:2529#L85)
``` go
func (*Int) Find(ints []int, predict func(int) bool) (int, bool)
```
Find returns the value of the first element in the slice that satisfies the provided predict function.
otherwise false is returned.




### <a name="Int.FindIndex">func</a> (\*Int) [FindIndex](/src/target/int.go?s=2772:2851#L97)
``` go
func (*Int) FindIndex(ints []int, predict func(int) bool, fromIndex ...int) int
```
FindIndex returns the index of the first element in the slice that satisfies the provided predict function.
otherwise -1 is returned.




### <a name="Int.FindLastIndex">func</a> (\*Int) [FindLastIndex](/src/target/int.go?s=3200:3283#L120)
``` go
func (*Int) FindLastIndex(ints []int, predict func(int) bool, fromIndex ...int) int
```
FindLastIndex  is like FindIndex except that it iterates over elements from right to left.




### <a name="Int.ForEach">func</a> (\*Int) [ForEach](/src/target/int.go?s=3609:3660#L143)
``` go
func (*Int) ForEach(ints []int, provided func(int))
```
ForEach executes a provided function once for each slice element.




### <a name="Int.From">func</a> (\*Int) [From](/src/target/int.go?s=4813:4871#L207)
``` go
func (*Int) From(ints []int, mapFn ...func(int) int) []int
```
From creates a new, shallow-copied slice.




### <a name="Int.Head">func</a> (\*Int) [Head](/src/target/int.go?s=10397:10437#L478)
``` go
func (*Int) Head(ints []int) (int, bool)
```
Head gets the first element of slice.




### <a name="Int.Includes">func</a> (\*Int) [Includes](/src/target/int.go?s=3801:3845#L150)
``` go
func (*Int) Includes(ints []int, e int) bool
```
Includes determines whether a slice includes a certain element, returning true or false.




### <a name="Int.IndexOf">func</a> (\*Int) [IndexOf](/src/target/int.go?s=4043:4103#L161)
``` go
func (*Int) IndexOf(ints []int, e int, fromIndex ...int) int
```
IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.




### <a name="Int.Initial">func</a> (\*Int) [Initial](/src/target/int.go?s=10558:10595#L487)
``` go
func (*Int) Initial(ints []int) []int
```
Initial gets all but the last element of slice.




### <a name="Int.Intersection">func</a> (\*Int) [Intersection](/src/target/int.go?s=11858:11903#L559)
``` go
func (i *Int) Intersection(vs ...[]int) []int
```
Intersection creates a slice of unique values that are included in all given slices.




### <a name="Int.Join">func</a> (\*Int) [Join](/src/target/int.go?s=6495:6545#L297)
``` go
func (*Int) Join(ints []int, sep ...string) string
```
Join joins all elements of a slice into a string.




### <a name="Int.Last">func</a> (\*Int) [Last](/src/target/int.go?s=10709:10749#L497)
``` go
func (*Int) Last(ints []int) (int, bool)
```
Last gets the last element of slice.




### <a name="Int.LastIndexOf">func</a> (\*Int) [LastIndexOf](/src/target/int.go?s=4462:4526#L184)
``` go
func (*Int) LastIndexOf(ints []int, e int, fromIndex ...int) int
```
LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.




### <a name="Int.Map">func</a> (\*Int) [Map](/src/target/int.go?s=6828:6878#L312)
``` go
func (*Int) Map(ints []int, m func(int) int) []int
```
Map creates a new slice with the results of calling a provided function on every element in the calling slice.




### <a name="Int.Nth">func</a> (\*Int) [Nth](/src/target/int.go?s=10937:10986#L508)
``` go
func (*Int) Nth(ints []int, n ...int) (int, bool)
```
Nth gets the element at index n of slice.
if n is negative, the nth element from the end is returned.




### <a name="Int.Of">func</a> (\*Int) [Of](/src/target/int.go?s=5136:5169#L225)
``` go
func (*Int) Of(ints ...int) []int
```
Of creates a new slice instance with a variable number of arguments.




### <a name="Int.Pop">func</a> (\*Int) [Pop](/src/target/int.go?s=7039:7085#L323)
``` go
func (*Int) Pop(ints []int) (int, bool, []int)
```
Pop removes the last element from a slice and returns that element.




### <a name="Int.Pull">func</a> (\*Int) [Pull](/src/target/int.go?s=11238:11285#L528)
``` go
func (i *Int) Pull(ints []int, vs ...int) []int
```
Pull removes all given values from slice.




### <a name="Int.Push">func</a> (\*Int) [Push](/src/target/int.go?s=7283:7335#L332)
``` go
func (*Int) Push(ints []int, es ...int) (int, []int)
```
Push adds one or more elements to the end of a slice and returns the new length of the slice.




### <a name="Int.Reduce">func</a> (\*Int) [Reduce](/src/target/int.go?s=7559:7632#L344)
``` go
func (*Int) Reduce(ints []int, accum func(int, int) int, init ...int) int
```
Reduce applies a function against each element in the slice to reduce it to a single value.




### <a name="Int.ReduceRight">func</a> (\*Int) [ReduceRight](/src/target/int.go?s=7867:7945#L356)
``` go
func (*Int) ReduceRight(ints []int, accum func(int, int) int, init ...int) int
```
ReduceRight applies a function against each value of the slice (from right-to-left) to reduce it to a single value.




### <a name="Int.Remove">func</a> (\*Int) [Remove](/src/target/int.go?s=12433:12493#L592)
``` go
func (*Int) Remove(ints []int, predict func(int) bool) []int
```
Remove removes all elements from slice that predicate returns truthy for.




### <a name="Int.Reverse">func</a> (\*Int) [Reverse](/src/target/int.go?s=8115:8152#L369)
``` go
func (*Int) Reverse(ints []int) []int
```
Reverse reverses a slice.




### <a name="Int.Shift">func</a> (\*Int) [Shift](/src/target/int.go?s=8348:8396#L379)
``` go
func (*Int) Shift(ints []int) (int, bool, []int)
```
Shift removes the first element from a slice and returns that element.




### <a name="Int.Shuffle">func</a> (\*Int) [Shuffle](/src/target/int.go?s=9792:9823#L450)
``` go
func (*Int) Shuffle(ints []int)
```
Shuffle randomizes the order of elements.




### <a name="Int.Slice">func</a> (\*Int) [Slice](/src/target/int.go?s=8978:9035#L406)
``` go
func (*Int) Slice(ints []int, startEndIndex ...int) []int
```
Slice returns a portion of a slice selected from start to end (end not included).
start and end indexes are optional, default to zero and length of the slice.




### <a name="Int.Some">func</a> (\*Int) [Some](/src/target/int.go?s=5734:5791#L257)
``` go
func (*Int) Some(ints []int, predict func(int) bool) bool
```
Some tests whether at least one element in the slice pass the predict implemented.




### <a name="Int.Sort">func</a> (\*Int) [Sort](/src/target/int.go?s=9453:9510#L438)
``` go
func (*Int) Sort(ints []int, less ...func(int, int) bool)
```
Sort sorts elements.




### <a name="Int.String">func</a> (\*Int) [String](/src/target/int.go?s=9985:10024#L456)
``` go
func (i *Int) String(ints []int) string
```
String returns a string representing the specified slice.




### <a name="Int.Tail">func</a> (\*Int) [Tail](/src/target/int.go?s=12655:12689#L605)
``` go
func (*Int) Tail(ints []int) []int
```
Tail gets all but the first element of slice.




### <a name="Int.Take">func</a> (\*Int) [Take](/src/target/int.go?s=12827:12871#L615)
``` go
func (*Int) Take(ints []int, n ...int) []int
```
Take creates a slice with n elements taken from the beginning.




### <a name="Int.Union">func</a> (\*Int) [Union](/src/target/int.go?s=13095:13135#L631)
``` go
func (i *Int) Union(ints ...[]int) []int
```
Union creates a slice of unique values, in order, from all given slices.




### <a name="Int.Uniq">func</a> (\*Int) [Uniq](/src/target/int.go?s=13338:13374#L644)
``` go
func (i *Int) Uniq(ints []int) []int
```
Uniq creates a duplicate-free version of a slice.




### <a name="Int.Unshift">func</a> (\*Int) [Unshift](/src/target/int.go?s=8599:8654#L388)
``` go
func (*Int) Unshift(ints []int, es ...int) (int, []int)
```
Unshift adds one or more elements to the beginning of a slice and returns the new length of the slice.




### <a name="Int.Without">func</a> (\*Int) [Without](/src/target/int.go?s=12272:12322#L587)
``` go
func (i *Int) Without(ints []int, vs ...int) []int
```
Without creates a slice excluding all given values.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
