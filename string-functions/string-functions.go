package main

import (
	"fmt"
	"strings"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.Index("test", "e"))
	p("Join: ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat: ", s.Repeat("a", 5))
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	p("Replace: ", s.Replace("foo", "o", "0", 1))
	//n大于0，就是替换几次；n小于0，就是全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	p("Split: ", s.Split("a-b-c-d-e", "-"))
	p("ToLower: ", s.ToLower("TEST"))
	p("ToUpper: ", s.ToUpper("test"))
}
