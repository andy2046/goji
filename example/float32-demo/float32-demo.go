package main

import (
	"fmt"

	G "github.com/andy2046/goji"
)

func main() {
	i := G.NewFloat32()

	f := func(i float32) float32 { return i + 1 }
	preFilter := []float32{5, 4, 3}
	preFilter2 := []float32{2, 1}

	r := i.From(preFilter)
	fmt.Println(r)

	r = i.From(preFilter, f)
	fmt.Println(r)

	r = i.From(append(preFilter, preFilter2...), f)
	fmt.Println(r)

	n := []float32{1, 22, 3}
	fmt.Println(n)
	i.Sort(n, func(i, j int) bool { return n[i] > n[j] })
	fmt.Println(n)

	r = i.Of(1, 3, 4)
	fmt.Println(r)

	r = i.Filter([]float32{1, 3, 4}, func(n float32) bool { return n > 2 })
	fmt.Println(r)

	id1, ok := i.Find([]float32{1, 3, 4}, func(n float32) bool { return n > 1 })
	fmt.Println(id1, ok)

	id2 := i.FindIndex([]float32{1, 3, 1, 5}, func(n float32) bool { return n > 1 }, 2)
	fmt.Println(id2)

	id3 := i.FindLastIndex([]float32{1, 3, 1, 5}, func(n float32) bool { return n > 1 }, 2)
	fmt.Println(id3)

	i.ForEach([]float32{1, 3, 5}, func(n float32) { fmt.Print(n) })
	fmt.Println()

	fmt.Println(i.Includes([]float32{1, 2, 3}, 3))

	fmt.Println(i.IndexOf([]float32{1, 3, 1, 5}, 1, 1))
	fmt.Println(i.LastIndexOf([]float32{1, 3, 1, 5}, 1, 1))

	fmt.Println(i.Concat([]float32{1, 3}, []float32{2, 4}))

	fmt.Println(i.Every([]float32{2, 4, 6}, func(n float32) bool { return n >= 2 }))
	fmt.Println(i.Some([]float32{2, 4, 6}, func(n float32) bool { return n > 5 }))

	fs := make([]float32, 3)
	i.Fill(fs, 2)
	fmt.Println(fs)

	fmt.Println(i.Join([]float32{3, 5, 7}))

	fmt.Println(i.Map([]float32{1, 3, 5}, func(n float32) float32 { return n * 2 }))

	fmt.Println(i.Reduce([]float32{1, 2, 3}, func(acc, cur float32) float32 { return acc + cur }, 4))

	fmt.Println(i.Reverse([]float32{1, 2, 3}))

	fmt.Println(i.Slice([]float32{1, 2, 3}, 1, 3))

	ss := []float32{1, 2, 3, 4, 5}
	i.Shuffle(ss)
	fmt.Println(ss)
	i.Shuffle(ss)
	fmt.Println(ss)

	fmt.Println(i.Pull([]float32{1, 2, 3, 4, 5}, 3, 5))

	fmt.Println(i.Difference([]float32{1, 2, 3}, []float32{2, 4}, []float32{3, 5}))
	fmt.Println(i.Intersection([]float32{1, 2, 3}, []float32{2, 4}, []float32{2, 5}))

	fmt.Println(i.Remove([]float32{1, 2, 3, 4, 5}, func(n float32) bool { return n > 2 }))

	fmt.Println(i.Union([]float32{1, 2}, []float32{2, 3}, []float32{3, 5}))
	fmt.Println(i.Uniq([]float32{1, 2, 1, 3, 2}))

	fmt.Println(i.Equal([]float32{1, 2}, []float32{1, 2}))

}
