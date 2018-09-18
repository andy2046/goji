

# float32util
`import "github.com/andy2046/goji/pkg/float32util"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package float32util provides functional utils for float.




## <a name="pkg-index">Index</a>
* [func FloatToString(fv float64) string](#FloatToString)
* [func FloatToStringWithPrec(fv float64, prec int) string](#FloatToStringWithPrec)
* [type Float32](#Float32)
  * [func NewFloat32() *Float32](#NewFloat32)
  * [func (*Float32) Concat(ints ...[]float32) []float32](#Float32.Concat)
  * [func (i *Float32) Difference(ints []float32, vs ...[]float32) []float32](#Float32.Difference)
  * [func (*Float32) Drop(ints []float32, n ...int) []float32](#Float32.Drop)
  * [func (*Float32) Equal(s1, s2 []float32) bool](#Float32.Equal)
  * [func (*Float32) Every(ints []float32, predict func(float32) bool) bool](#Float32.Every)
  * [func (*Float32) Fill(ints []float32, value float32, startEndIndex ...int)](#Float32.Fill)
  * [func (*Float32) Filter(ints []float32, predict func(float32) bool) []float32](#Float32.Filter)
  * [func (*Float32) Find(ints []float32, predict func(float32) bool) (float32, bool)](#Float32.Find)
  * [func (*Float32) FindIndex(ints []float32, predict func(float32) bool, fromIndex ...int) int](#Float32.FindIndex)
  * [func (*Float32) FindLastIndex(ints []float32, predict func(float32) bool, fromIndex ...int) int](#Float32.FindLastIndex)
  * [func (*Float32) ForEach(ints []float32, provided func(float32))](#Float32.ForEach)
  * [func (*Float32) From(ints []float32, mapFn ...func(float32) float32) []float32](#Float32.From)
  * [func (*Float32) Head(ints []float32) (float32, bool)](#Float32.Head)
  * [func (*Float32) Includes(ints []float32, e float32) bool](#Float32.Includes)
  * [func (*Float32) IndexOf(ints []float32, e float32, fromIndex ...int) int](#Float32.IndexOf)
  * [func (*Float32) Initial(ints []float32) []float32](#Float32.Initial)
  * [func (i *Float32) Intersection(vs ...[]float32) []float32](#Float32.Intersection)
  * [func (*Float32) Join(ints []float32, sep ...string) string](#Float32.Join)
  * [func (*Float32) JoinWithPrec(ints []float32, prec int, sep ...string) string](#Float32.JoinWithPrec)
  * [func (*Float32) Last(ints []float32) (float32, bool)](#Float32.Last)
  * [func (*Float32) LastIndexOf(ints []float32, e float32, fromIndex ...int) int](#Float32.LastIndexOf)
  * [func (*Float32) Map(ints []float32, m func(float32) float32) []float32](#Float32.Map)
  * [func (*Float32) Max(x, y float32) float32](#Float32.Max)
  * [func (*Float32) Min(x, y float32) float32](#Float32.Min)
  * [func (*Float32) Nth(ints []float32, n ...int) (float32, bool)](#Float32.Nth)
  * [func (*Float32) Of(ints ...float32) []float32](#Float32.Of)
  * [func (*Float32) Pop(ints []float32) (float32, bool, []float32)](#Float32.Pop)
  * [func (i *Float32) Pull(ints []float32, vs ...float32) []float32](#Float32.Pull)
  * [func (*Float32) Push(ints []float32, es ...float32) (int, []float32)](#Float32.Push)
  * [func (*Float32) Reduce(ints []float32, accum func(float32, float32) float32, init ...float32) float32](#Float32.Reduce)
  * [func (*Float32) ReduceRight(ints []float32, accum func(float32, float32) float32, init ...float32) float32](#Float32.ReduceRight)
  * [func (*Float32) Remove(ints []float32, predict func(float32) bool) []float32](#Float32.Remove)
  * [func (*Float32) Reverse(ints []float32) []float32](#Float32.Reverse)
  * [func (*Float32) Shift(ints []float32) (float32, bool, []float32)](#Float32.Shift)
  * [func (*Float32) Shuffle(ints []float32)](#Float32.Shuffle)
  * [func (*Float32) Slice(ints []float32, startEndIndex ...int) []float32](#Float32.Slice)
  * [func (*Float32) Some(ints []float32, predict func(float32) bool) bool](#Float32.Some)
  * [func (*Float32) Sort(ints []float32, less ...func(int, int) bool)](#Float32.Sort)
  * [func (i *Float32) String(ints []float32) string](#Float32.String)
  * [func (*Float32) Tail(ints []float32) []float32](#Float32.Tail)
  * [func (*Float32) Take(ints []float32, n ...int) []float32](#Float32.Take)
  * [func (i *Float32) Union(ints ...[]float32) []float32](#Float32.Union)
  * [func (i *Float32) Uniq(ints []float32) []float32](#Float32.Uniq)
  * [func (*Float32) Unshift(ints []float32, es ...float32) (int, []float32)](#Float32.Unshift)
  * [func (i *Float32) Without(ints []float32, vs ...float32) []float32](#Float32.Without)
* [type IFloat32](#IFloat32)


#### <a name="pkg-files">Package files</a>
[float32.go](/src/github.com/andy2046/goji/pkg/float32util/float32.go) 





## <a name="FloatToString">func</a> [FloatToString](/src/target/float32.go?s=384:421#L19)
``` go
func FloatToString(fv float64) string
```
FloatToString converts float to string.



## <a name="FloatToStringWithPrec">func</a> [FloatToStringWithPrec](/src/target/float32.go?s=233:288#L14)
``` go
func FloatToStringWithPrec(fv float64, prec int) string
```
FloatToStringWithPrec converts float to string with precision.




## <a name="Float32">type</a> [Float32](/src/target/float32.go?s=506:527#L24)
``` go
type Float32 struct{}
```
Float32 for float32 type slice.







### <a name="NewFloat32">func</a> [NewFloat32](/src/target/float32.go?s=2807:2833#L75)
``` go
func NewFloat32() *Float32
```
NewFloat32 creates a new Float32 instance.





### <a name="Float32.Concat">func</a> (\*Float32) [Concat](/src/target/float32.go?s=6521:6572#L262)
``` go
func (*Float32) Concat(ints ...[]float32) []float32
```
Concat merges two or more slices.




### <a name="Float32.Difference">func</a> (\*Float32) [Difference](/src/target/float32.go?s=13462:13533#L581)
``` go
func (i *Float32) Difference(ints []float32, vs ...[]float32) []float32
```
Difference creates a slice of values not included in the other given slices.




### <a name="Float32.Drop">func</a> (\*Float32) [Drop](/src/target/float32.go?s=12003:12059#L503)
``` go
func (*Float32) Drop(ints []float32, n ...int) []float32
```
Drop creates a slice with n elements dropped from the beginning.




### <a name="Float32.Equal">func</a> (\*Float32) [Equal](/src/target/float32.go?s=15667:15711#L697)
``` go
func (*Float32) Equal(s1, s2 []float32) bool
```
Equal performs a deep comparison between two slice values.




### <a name="Float32.Every">func</a> (\*Float32) [Every](/src/target/float32.go?s=6737:6807#L273)
``` go
func (*Float32) Every(ints []float32, predict func(float32) bool) bool
```
Every tests whether all elements in the slice pass the predict implemented.




### <a name="Float32.Fill">func</a> (\*Float32) [Fill](/src/target/float32.go?s=7289:7362#L296)
``` go
func (*Float32) Fill(ints []float32, value float32, startEndIndex ...int)
```
Fill fills all the elements of a slice from a start index to an end index with a static value.
the end index is not included.




### <a name="Float32.Filter">func</a> (\*Float32) [Filter](/src/target/float32.go?s=3213:3289#L98)
``` go
func (*Float32) Filter(ints []float32, predict func(float32) bool) []float32
```
Filter creates a new slice with all elements that pass the predict implemented.




### <a name="Float32.Find">func</a> (\*Float32) [Find](/src/target/float32.go?s=3543:3623#L112)
``` go
func (*Float32) Find(ints []float32, predict func(float32) bool) (float32, bool)
```
Find returns the value of the first element in the slice that satisfies the provided predict function.
otherwise false is returned.




### <a name="Float32.FindIndex">func</a> (\*Float32) [FindIndex](/src/target/float32.go?s=3866:3957#L124)
``` go
func (*Float32) FindIndex(ints []float32, predict func(float32) bool, fromIndex ...int) int
```
FindIndex returns the index of the first element in the slice that satisfies the provided predict function.
otherwise -1 is returned.




### <a name="Float32.FindLastIndex">func</a> (\*Float32) [FindLastIndex](/src/target/float32.go?s=4309:4404#L147)
``` go
func (*Float32) FindLastIndex(ints []float32, predict func(float32) bool, fromIndex ...int) int
```
FindLastIndex  is like FindIndex except that it iterates over elements from right to left.




### <a name="Float32.ForEach">func</a> (\*Float32) [ForEach](/src/target/float32.go?s=4736:4799#L170)
``` go
func (*Float32) ForEach(ints []float32, provided func(float32))
```
ForEach executes a provided function once for each slice element.




### <a name="Float32.From">func</a> (\*Float32) [From](/src/target/float32.go?s=5994:6072#L234)
``` go
func (*Float32) From(ints []float32, mapFn ...func(float32) float32) []float32
```
From creates a new, shallow-copied slice.




### <a name="Float32.Head">func</a> (\*Float32) [Head](/src/target/float32.go?s=12304:12356#L520)
``` go
func (*Float32) Head(ints []float32) (float32, bool)
```
Head gets the first element of slice.




### <a name="Float32.Includes">func</a> (\*Float32) [Includes](/src/target/float32.go?s=4940:4996#L177)
``` go
func (*Float32) Includes(ints []float32, e float32) bool
```
Includes determines whether a slice includes a certain element, returning true or false.




### <a name="Float32.IndexOf">func</a> (\*Float32) [IndexOf](/src/target/float32.go?s=5194:5266#L188)
``` go
func (*Float32) IndexOf(ints []float32, e float32, fromIndex ...int) int
```
IndexOf returns the first index at which a given element can be found in the slice, or -1 if it is not present.




### <a name="Float32.Initial">func</a> (\*Float32) [Initial](/src/target/float32.go?s=12477:12526#L529)
``` go
func (*Float32) Initial(ints []float32) []float32
```
Initial gets all but the last element of slice.




### <a name="Float32.Intersection">func</a> (\*Float32) [Intersection](/src/target/float32.go?s=13857:13914#L601)
``` go
func (i *Float32) Intersection(vs ...[]float32) []float32
```
Intersection creates a slice of unique values that are included in all given slices.




### <a name="Float32.Join">func</a> (\*Float32) [Join](/src/target/float32.go?s=7786:7844#L324)
``` go
func (*Float32) Join(ints []float32, sep ...string) string
```
Join joins all elements of a slice into a string.




### <a name="Float32.JoinWithPrec">func</a> (\*Float32) [JoinWithPrec](/src/target/float32.go?s=8099:8175#L339)
``` go
func (*Float32) JoinWithPrec(ints []float32, prec int, sep ...string) string
```
JoinWithPrec joins all elements of a slice into a string with precision.




### <a name="Float32.Last">func</a> (\*Float32) [Last](/src/target/float32.go?s=12644:12696#L539)
``` go
func (*Float32) Last(ints []float32) (float32, bool)
```
Last gets the last element of slice.




### <a name="Float32.LastIndexOf">func</a> (\*Float32) [LastIndexOf](/src/target/float32.go?s=5628:5704#L211)
``` go
func (*Float32) LastIndexOf(ints []float32, e float32, fromIndex ...int) int
```
LastIndexOf returns the last index at which a given element can be found in the slice, or -1 if it is not present.




### <a name="Float32.Map">func</a> (\*Float32) [Map](/src/target/float32.go?s=8482:8552#L354)
``` go
func (*Float32) Map(ints []float32, m func(float32) float32) []float32
```
Map creates a new slice with the results of calling a provided function on every element in the calling slice.




### <a name="Float32.Max">func</a> (\*Float32) [Max](/src/target/float32.go?s=3047:3088#L90)
``` go
func (*Float32) Max(x, y float32) float32
```
Max returns the larger of x or y.




### <a name="Float32.Min">func</a> (\*Float32) [Min](/src/target/float32.go?s=2927:2968#L82)
``` go
func (*Float32) Min(x, y float32) float32
```
Min returns the smaller of x or y.




### <a name="Float32.Nth">func</a> (\*Float32) [Nth](/src/target/float32.go?s=12884:12945#L550)
``` go
func (*Float32) Nth(ints []float32, n ...int) (float32, bool)
```
Nth gets the element at index n of slice.
if n is negative, the nth element from the end is returned.




### <a name="Float32.Of">func</a> (\*Float32) [Of](/src/target/float32.go?s=6349:6394#L252)
``` go
func (*Float32) Of(ints ...float32) []float32
```
Of creates a new slice instance with a variable number of arguments.




### <a name="Float32.Pop">func</a> (\*Float32) [Pop](/src/target/float32.go?s=8717:8779#L365)
``` go
func (*Float32) Pop(ints []float32) (float32, bool, []float32)
```
Pop removes the last element from a slice and returns that element.




### <a name="Float32.Pull">func</a> (\*Float32) [Pull](/src/target/float32.go?s=13197:13260#L570)
``` go
func (i *Float32) Pull(ints []float32, vs ...float32) []float32
```
Pull removes all given values from slice.




### <a name="Float32.Push">func</a> (\*Float32) [Push](/src/target/float32.go?s=8981:9049#L374)
``` go
func (*Float32) Push(ints []float32, es ...float32) (int, []float32)
```
Push adds one or more elements to the end of a slice and returns the new length of the slice.




### <a name="Float32.Reduce">func</a> (\*Float32) [Reduce](/src/target/float32.go?s=9280:9381#L386)
``` go
func (*Float32) Reduce(ints []float32, accum func(float32, float32) float32, init ...float32) float32
```
Reduce applies a function against each element in the slice to reduce it to a single value.




### <a name="Float32.ReduceRight">func</a> (\*Float32) [ReduceRight](/src/target/float32.go?s=9620:9726#L398)
``` go
func (*Float32) ReduceRight(ints []float32, accum func(float32, float32) float32, init ...float32) float32
```
ReduceRight applies a function against each value of the slice (from right-to-left) to reduce it to a single value.




### <a name="Float32.Remove">func</a> (\*Float32) [Remove](/src/target/float32.go?s=14464:14540#L634)
``` go
func (*Float32) Remove(ints []float32, predict func(float32) bool) []float32
```
Remove removes all elements from slice that predicate returns truthy for.




### <a name="Float32.Reverse">func</a> (\*Float32) [Reverse](/src/target/float32.go?s=9900:9949#L411)
``` go
func (*Float32) Reverse(ints []float32) []float32
```
Reverse reverses a slice.




### <a name="Float32.Shift">func</a> (\*Float32) [Shift](/src/target/float32.go?s=10145:10209#L421)
``` go
func (*Float32) Shift(ints []float32) (float32, bool, []float32)
```
Shift removes the first element from a slice and returns that element.




### <a name="Float32.Shuffle">func</a> (\*Float32) [Shuffle](/src/target/float32.go?s=11660:11699#L492)
``` go
func (*Float32) Shuffle(ints []float32)
```
Shuffle randomizes the order of elements.




### <a name="Float32.Slice">func</a> (\*Float32) [Slice](/src/target/float32.go?s=10815:10884#L448)
``` go
func (*Float32) Slice(ints []float32, startEndIndex ...int) []float32
```
Slice returns a portion of a slice selected from start to end (end not included).
start and end indexes are optional, default to zero and length of the slice.




### <a name="Float32.Some">func</a> (\*Float32) [Some](/src/target/float32.go?s=6991:7060#L284)
``` go
func (*Float32) Some(ints []float32, predict func(float32) bool) bool
```
Some tests whether at least one element in the slice pass the predict implemented.




### <a name="Float32.Sort">func</a> (\*Float32) [Sort](/src/target/float32.go?s=11313:11378#L480)
``` go
func (*Float32) Sort(ints []float32, less ...func(int, int) bool)
```
Sort sorts elements.




### <a name="Float32.String">func</a> (\*Float32) [String](/src/target/float32.go?s=11861:11908#L498)
``` go
func (i *Float32) String(ints []float32) string
```
String returns a string representing the specified slice.




### <a name="Float32.Tail">func</a> (\*Float32) [Tail](/src/target/float32.go?s=14706:14752#L647)
``` go
func (*Float32) Tail(ints []float32) []float32
```
Tail gets all but the first element of slice.




### <a name="Float32.Take">func</a> (\*Float32) [Take](/src/target/float32.go?s=14894:14950#L657)
``` go
func (*Float32) Take(ints []float32, n ...int) []float32
```
Take creates a slice with n elements taken from the beginning.




### <a name="Float32.Union">func</a> (\*Float32) [Union](/src/target/float32.go?s=15177:15229#L673)
``` go
func (i *Float32) Union(ints ...[]float32) []float32
```
Union creates a slice of unique values, in order, from all given slices.




### <a name="Float32.Uniq">func</a> (\*Float32) [Uniq](/src/target/float32.go?s=15436:15484#L686)
``` go
func (i *Float32) Uniq(ints []float32) []float32
```
Uniq creates a duplicate-free version of a slice.




### <a name="Float32.Unshift">func</a> (\*Float32) [Unshift](/src/target/float32.go?s=10416:10487#L430)
``` go
func (*Float32) Unshift(ints []float32, es ...float32) (int, []float32)
```
Unshift adds one or more elements to the beginning of a slice and returns the new length of the slice.




### <a name="Float32.Without">func</a> (\*Float32) [Without](/src/target/float32.go?s=14287:14353#L629)
``` go
func (i *Float32) Without(ints []float32, vs ...float32) []float32
```
Without creates a slice excluding all given values.




## <a name="IFloat32">type</a> [IFloat32](/src/target/float32.go?s=567:2759#L27)
``` go
type IFloat32 interface {
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
```
IFloat32 is the Float32 interface.














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
