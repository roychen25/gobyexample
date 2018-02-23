package gobyexample

import "fmt"
import str "strings"

var p = fmt.Println

// StringFunctionsDemo - demonstrates Go's builtin string functions
func StringFunctionsDemo() {
	// Sample of the functions available in the `strings` package.
	p("Contains:", str.Contains("test", "es"))
	p("Count:", str.Count("test", "t"))
	p("HasPrefix:", str.HasPrefix("test", "te"))
	p("HasSuffix:", str.HasSuffix("test", "st"))
	p("Index:", str.Index("test", "e"))
	p("Join:", str.Join([]string{"a", "b"}, "-"))
	p("Repeat:", str.Repeat("a", 5))
	p("Replace:", str.Replace("foo", "o", "0", -1))
	p("Replace:", str.Replace("foo", "o", "0", 1))
	p("Split:", str.Split("a-b-c-d-e", "-"))
	p("ToLower:", str.ToLower("TEST"))
	p("ToUpper:", str.ToUpper("test"))
	p()

	// Not part of the `strings` package, but worth mentioning here
	p("Len:", len("hello"))
	p("Char:", "hello"[1])
}
