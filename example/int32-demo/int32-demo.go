package main

import (
	"fmt"

	Int32Util "github.com/andy2046/goji/pkg/int32util"
)

func main() {
	i := Int32Util.NewInt32()

	f := func(i int32) int32 { return i + 1 }
	preFilter := []int32{5, 4, 3}
	preFilter2 := []int32{2, 1}

	r := i.From(preFilter)
	fmt.Println(r)

	r = i.From(preFilter, f)
	fmt.Println(r)

	r = i.From(append(preFilter, preFilter2...), f)
	fmt.Println(r)

	n := []int32{1, 22, 3}
	fmt.Println(n)
	i.Sort(n, func(i, j int) bool { return n[i]%2 > n[j]%2 })
	fmt.Println(n)

	r = i.Of(1, 3, 4)
	fmt.Println(r)

	r = i.Filter([]int32{1, 3, 4}, func(n int32) bool { return n%2 == 0 })
	fmt.Println(r)

	idx, ok := i.Find([]int32{1, 3, 4}, func(n int32) bool { return n > 1 })
	fmt.Println(idx, ok)

	idx = i.FindIndex([]int32{1, 3, 1, 5}, func(n int32) bool { return n > 1 }, 2)
	fmt.Println(idx)

	idx = i.FindLastIndex([]int32{1, 3, 1, 5}, func(n int32) bool { return n > 1 }, 2)
	fmt.Println(idx)

	i.ForEach([]int32{1, 3, 5}, func(n int32) { fmt.Print(n) })
	fmt.Println()

	fmt.Println(i.Includes([]int32{1, 2, 3}, 3))

	fmt.Println(i.IndexOf([]int32{1, 3, 1, 5}, 1, 1))
	fmt.Println(i.LastIndexOf([]int32{1, 3, 1, 5}, 1, 1))

	fmt.Println(i.Concat([]int32{1, 3}, []int32{2, 4}))

	fmt.Println(i.Every([]int32{2, 4, 6}, func(n int32) bool { return n%2 == 0 }))
	fmt.Println(i.Some([]int32{2, 4, 6}, func(n int32) bool { return n > 5 }))

	fs := make([]int32, 3)
	i.Fill(fs, 2)
	fmt.Println(fs)

	fmt.Println(i.Join([]int32{3, 5, 7}))

	fmt.Println(i.Map([]int32{1, 3, 5}, func(n int32) int32 { return n * 2 }))

	fmt.Println(i.Reduce([]int32{1, 2, 3}, func(acc, cur int32) int32 { return acc + cur }, 4))

	fmt.Println(i.Reverse([]int32{1, 2, 3}))

	fmt.Println(i.Slice([]int32{1, 2, 3}, 1, 3))

	ss := []int32{1, 2, 3, 4, 5}
	i.Shuffle(ss)
	fmt.Println(ss)
	i.Shuffle(ss)
	fmt.Println(ss)

	fmt.Println(i.Pull([]int32{1, 2, 3, 4, 5}, 3, 5))

	fmt.Println(i.Difference([]int32{1, 2, 3}, []int32{2, 4}, []int32{3, 5}))
	fmt.Println(i.Intersection([]int32{1, 2, 3}, []int32{2, 4}, []int32{2, 5}))

	fmt.Println(i.Remove([]int32{1, 2, 3, 4, 5}, func(n int32) bool { return n%2 == 0 }))

	fmt.Println(i.Union([]int32{1, 2}, []int32{2, 3}, []int32{3, 5}))
	fmt.Println(i.Uniq([]int32{1, 2, 1, 3, 2}))

	fmt.Println(i.Equal([]int32{1, 2}, []int32{1, 2}))

}
