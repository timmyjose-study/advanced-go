package main

// we can alias packages so
import (
	"fmt"
	s "strings"
)

// we can have aliases for functions as well
var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es"))
	p("Count: ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index: ", s.Index("test", "e"))
	p("Join: ", s.Join([]string{"hello", ", ", "world"}, ""))
	p("Repeat: ", s.Repeat("hello", 10))
	p("Replace: ", s.Replace("fooo", "o", "0", -1)) // -1 means "all"
	p("Replace: ", s.Replace("fooo", "o", "0", 2))  // only 2 occurrences
	p("Split: ", s.Split("a-b-c-d-e", "-"))
	p("ToLower: ", s.ToLower("Hello, world"))
	p("ToUpper: ", s.ToUpper("Hello, world"))
	p()
	p("Len: ", len("Hello"))
	p("Char: ", "Hola"[1])
}
