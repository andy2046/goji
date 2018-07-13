package main

import (
	"fmt"

	G "github.com/andy2046/goji"
)

func main() {
	i := G.NewString()

	f := func(i string) string { return i + "-" }
	preFilter := []string{"a", "b", "c"}
	preFilter2 := []string{"d", "e"}

	r := i.From(preFilter)
	fmt.Println(r)

	r = i.From(preFilter, f)
	fmt.Println(r)

	r = i.From(append(preFilter, preFilter2...), f)
	fmt.Println(r)

	n := []string{"eee", "aaa", "ddd"}
	fmt.Println(n)
	i.Sort(n, func(i, j int) bool { return n[i] > n[j] })
	fmt.Println(n)

	r = i.Of("d", "n", "a")
	fmt.Println(r)

	r = i.Filter([]string{"d", "n", "a"}, func(n string) bool { return n > "e" })
	fmt.Println(r)

	id1, ok := i.Find([]string{"d", "n", "a"}, func(n string) bool { return n > "d" })
	fmt.Println(id1, ok)

	id2 := i.FindIndex([]string{"d", "n", "a", "s"}, func(n string) bool { return n > "a" }, 2)
	fmt.Println(id2)

	id3 := i.FindLastIndex([]string{"d", "n", "a", "s"}, func(n string) bool { return n > "d" }, 2)
	fmt.Println(id3)

	i.ForEach([]string{"d", "n", "a"}, func(n string) { fmt.Print(n) })
	fmt.Println()

	fmt.Println(i.Includes([]string{"d", "n", "a"}, "d"))

	fmt.Println(i.IndexOf([]string{"d", "n", "a"}, "a", 1))
	fmt.Println(i.LastIndexOf([]string{"d", "n", "a", "n"}, "n", 2))

	fmt.Println(i.Concat([]string{"d", "n"}, []string{"a"}))

	fmt.Println(i.Every([]string{"d", "n", "s"}, func(n string) bool { return n > "a" }))
	fmt.Println(i.Some([]string{"d", "n", "a"}, func(n string) bool { return n > "d" }))

	fs := make([]string, 3)
	i.Fill(fs, "c")
	fmt.Println(fs)

	fmt.Println(i.Join([]string{"d", "n", "a"}))

	fmt.Println(i.Map([]string{"d", "n", "a"}, func(n string) string { return n + n }))

	fmt.Println(i.Reduce([]string{"d", "n", "a"}, func(acc, cur string) string { return acc + cur }, "i"))

	fmt.Println(i.Reverse([]string{"d", "n", "a"}))

	fmt.Println(i.Slice([]string{"d", "n", "a"}, 1, 3))

	ss := []string{"d", "n", "a", "s", "x"}
	i.Shuffle(ss)
	fmt.Println(ss)
	i.Shuffle(ss)
	fmt.Println(ss)

	fmt.Println(i.Pull([]string{"d", "n", "a", "s", "x"}, "s", "x"))

	fmt.Println(i.Difference([]string{"d", "n", "a"}, []string{"d", "s"}, []string{"n", "s"}))
	fmt.Println(i.Intersection([]string{"d", "n", "a"}, []string{"d", "s"}, []string{"d"}))

	fmt.Println(i.Remove([]string{"d", "n", "a", "s", "x"}, func(n string) bool { return n > "o" }))

	fmt.Println(i.Union([]string{"d", "n", "a"}, []string{"d", "s"}, []string{"n", "x"}))
	fmt.Println(i.Uniq([]string{"d", "n", "a", "n", "d"}))

	fmt.Println(i.Equal([]string{"d", "n", "a"}, []string{"d", "n", "a"}))

	camelCase := i.CamelCase

	fmt.Println(camelCase("foo-bar-baz"))
	// fooBarBaz
	fmt.Println(camelCase("foo_bar"))
	// fooBar
	fmt.Println(camelCase("Foo-Bar"))
	// fooBar
	fmt.Println(camelCase("Foo-Bar", true))
	// FooBar
	fmt.Println(camelCase("--foo.bar"))
	// fooBar
	fmt.Println(camelCase("foo bar"))
	// fooBar
}
