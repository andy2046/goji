package main

import (
	"fmt"

	Int64Util "github.com/andy2046/goji/pkg/int64util"
)

func main() {
	i := Int64Util.NewInt64()

	f := func(i int64) int64 { return i + 1 }
	preFilter := []int64{5, 4, 3}
	preFilter2 := []int64{2, 1}

	r := i.From(preFilter)
	fmt.Println(r)

	r = i.From(preFilter, f)
	fmt.Println(r)

	r = i.From(append(preFilter, preFilter2...), f)
	fmt.Println(r)

	n := []int64{1, 22, 3}
	fmt.Println(n)
	i.Sort(n, func(i, j int) bool { return n[i]%2 > n[j]%2 })
	fmt.Println(n)

	r = i.Of(1, 3, 4)
	fmt.Println(r)

	r = i.Filter([]int64{1, 3, 4}, func(n int64) bool { return n%2 == 0 })
	fmt.Println(r)

	idx, ok := i.Find([]int64{1, 3, 4}, func(n int64) bool { return n > 1 })
	fmt.Println(idx, ok)

	idx = i.FindIndex([]int64{1, 3, 1, 5}, func(n int64) bool { return n > 1 }, 2)
	fmt.Println(idx)

	idx = i.FindLastIndex([]int64{1, 3, 1, 5}, func(n int64) bool { return n > 1 }, 2)
	fmt.Println(idx)

	i.ForEach([]int64{1, 3, 5}, func(n int64) { fmt.Print(n) })
	fmt.Println()

	fmt.Println(i.Includes([]int64{1, 2, 3}, 3))

	fmt.Println(i.IndexOf([]int64{1, 3, 1, 5}, 1, 1))
	fmt.Println(i.LastIndexOf([]int64{1, 3, 1, 5}, 1, 1))

	fmt.Println(i.Concat([]int64{1, 3}, []int64{2, 4}))

	fmt.Println(i.Every([]int64{2, 4, 6}, func(n int64) bool { return n%2 == 0 }))
	fmt.Println(i.Some([]int64{2, 4, 6}, func(n int64) bool { return n > 5 }))

	fs := make([]int64, 3)
	i.Fill(fs, 2)
	fmt.Println(fs)

	fmt.Println(i.Join([]int64{3, 5, 7}))

	fmt.Println(i.Map([]int64{1, 3, 5}, func(n int64) int64 { return n * 2 }))

	fmt.Println(i.Reduce([]int64{1, 2, 3}, func(acc, cur int64) int64 { return acc + cur }, 4))

	fmt.Println(i.Reverse([]int64{1, 2, 3}))

	fmt.Println(i.Slice([]int64{1, 2, 3}, 1, 3))

	ss := []int64{1, 2, 3, 4, 5}
	i.Shuffle(ss)
	fmt.Println(ss)
	i.Shuffle(ss)
	fmt.Println(ss)

	fmt.Println(i.Pull([]int64{1, 2, 3, 4, 5}, 3, 5))

	fmt.Println(i.Difference([]int64{1, 2, 3}, []int64{2, 4}, []int64{3, 5}))
	fmt.Println(i.Intersection([]int64{1, 2, 3}, []int64{2, 4}, []int64{2, 5}))

	fmt.Println(i.Remove([]int64{1, 2, 3, 4, 5}, func(n int64) bool { return n%2 == 0 }))

	fmt.Println(i.Union([]int64{1, 2}, []int64{2, 3}, []int64{3, 5}))
	fmt.Println(i.Uniq([]int64{1, 2, 1, 3, 2}))

	fmt.Println(i.Equal([]int64{1, 2}, []int64{1, 2}))

}
