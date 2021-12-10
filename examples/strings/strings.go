package main

import (
	"fmt"
	"strconv"
	"strings"
)

func mapper(c rune) rune {
	if c == 'o' {
		return '0'
	} else {
		return c
	}
}

func main() {

	s1 := "quick brown fox"
	s2 := "green neon nights on the moon"

	// basic operations
	fmt.Println(strings.EqualFold("bear", "BeaR"))
	fmt.Println(strings.Title(s1))   // every word
	fmt.Println(strings.ToTitle(s1)) // all words uppercased
	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.ToLower(s1))

	// search
	fmt.Println(strings.Contains(s1, "fox"))
	fmt.Println(strings.ContainsAny(s2, "ae"))
	fmt.Println(strings.Index(s1, "brown"))
	fmt.Println(strings.IndexAny(s2, "nr")) // first position of any matching character
	fmt.Println(strings.HasPrefix(s2, "green"))
	fmt.Println(strings.HasSuffix(s2, "moon"))
	fmt.Println(strings.Count(s2, "on"))

	// manipulation
	fmt.Println(strings.Split(s2, "n"))
	fmt.Println(strings.Join([]string{"big", "bird", "says", "hey"}, "-"))

	fmt.Println(strings.Fields(s2))
	fmt.Println(strings.FieldsFunc(s2, func(c rune) bool { return c == 'o' }))
	fmt.Println(strings.ReplaceAll(s2, " ", "~"))

	fmt.Println(strings.Map(mapper, s2))

	// builder
	var sb strings.Builder
	for _, v := range []string{"one", "two", "three"} {
		sb.WriteString(v)
		sb.WriteRune(',')
	}

	fmt.Println(sb.Len())
	fmt.Println(sb.Cap())
	sb.Grow(5)
	fmt.Println(sb.Len())
	fmt.Println(sb.String())
	sb.Reset()
	fmt.Println("|" + sb.String() + "|")

	// parsing
	fmt.Println(strconv.Atoi("5"))
	fmt.Println(strconv.Itoa(33))
	fmt.Println(strconv.ParseBool("T"))
	f, _ := strconv.ParseFloat("6.28", 64)
	fmt.Printf("%[1]f %.4[1]f\n", f) // repeated use of the argument

}
