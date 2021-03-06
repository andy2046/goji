package main

import (
	"fmt"

	IntUtil "github.com/andy2046/goji/pkg/intutil"
)

func main() {
	i := IntUtil.NewInt()

	f := func(i int) int { return i + 1 }
	preFilter := []int{5, 4, 3}
	preFilter2 := []int{2, 1}

	r := i.From(preFilter)
	fmt.Println(r)

	r = i.From(preFilter, f)
	fmt.Println(r)

	r = i.From(append(preFilter, preFilter2...), f)
	fmt.Println(r)

	n := []int{1, 22, 3}
	fmt.Println(n)
	i.Sort(n, func(i, j int) bool { return n[i]%2 > n[j]%2 })
	fmt.Println(n)

	r = i.Of(1, 3, 4)
	fmt.Println(r)

	r = i.Filter([]int{1, 3, 4}, func(n int) bool { return n%2 == 0 })
	fmt.Println(r)

	idx, ok := i.Find([]int{1, 3, 4}, func(n int) bool { return n > 1 })
	fmt.Println(idx, ok)

	idx = i.FindIndex([]int{1, 3, 1, 5}, func(n int) bool { return n > 1 }, 2)
	fmt.Println(idx)

	idx = i.FindLastIndex([]int{1, 3, 1, 5}, func(n int) bool { return n > 1 }, 2)
	fmt.Println(idx)

	i.ForEach([]int{1, 3, 5}, func(n int) { fmt.Print(n) })
	fmt.Println()

	fmt.Println(i.Includes([]int{1, 2, 3}, 3))

	fmt.Println(i.IndexOf([]int{1, 3, 1, 5}, 1, 1))
	fmt.Println(i.LastIndexOf([]int{1, 3, 1, 5}, 1, 1))

	fmt.Println(i.Concat([]int{1, 3}, []int{2, 4}))

	fmt.Println(i.Every([]int{2, 4, 6}, func(n int) bool { return n%2 == 0 }))
	fmt.Println(i.Some([]int{2, 4, 6}, func(n int) bool { return n > 5 }))

	fs := make([]int, 3)
	i.Fill(fs, 2)
	fmt.Println(fs)

	fmt.Println(i.Join([]int{3, 5, 7}))

	fmt.Println(i.Map([]int{1, 3, 5}, func(n int) int { return n * 2 }))

	fmt.Println(i.Reduce([]int{1, 2, 3}, func(acc, cur int) int { return acc + cur }, 4))

	fmt.Println(i.Reverse([]int{1, 2, 3}))

	fmt.Println(i.Slice([]int{1, 2, 3}, 1, 3))

	ss := []int{1, 2, 3, 4, 5}
	i.Shuffle(ss)
	fmt.Println(ss)
	i.Shuffle(ss)
	fmt.Println(ss)

	fmt.Println(i.Pull([]int{1, 2, 3, 4, 5}, 3, 5))

	fmt.Println(i.Difference([]int{1, 2, 3}, []int{2, 4}, []int{3, 5}))
	fmt.Println(i.Intersection([]int{1, 2, 3}, []int{2, 4}, []int{2, 5}))

	fmt.Println(i.Remove([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 }))

	fmt.Println(i.Union([]int{1, 2}, []int{2, 3}, []int{3, 5}))
	fmt.Println(i.Uniq([]int{1, 2, 1, 3, 2}))

	fmt.Println(i.Equal([]int{1, 2}, []int{1, 2}))

}
