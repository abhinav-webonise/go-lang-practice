package main

import "fmt"
import s "strings"

var p = fmt.Println

func main() {
    p("Contains", s.Contains("test","es"))
    p("Count", s.Count("test","t"))
    p("HasPrefix", s.HasPrefix("test"," "))
    p("HasSuffix", s.HasSuffix("test", "st"))
    p("Index:", s.Index("test", "t"))
    p("Join:", s.Join([]string{"a","b"}, "-"))
    p("Repeat:", s.Repeat("test", 5))
    p("Replace", s.Replace("fooo", "o", "0", -1))
    p("Replace", s.Replace("fooo", "o", "0", 1))
    p("Split", s.Split("a-b-c-d", "-"))
    p("ToLower", s.ToLower("Abhinav"))
    p("ToUpper", s.ToUpper("abhinav"))
    p()

    p("Length", len("hello"))
    p("Char", "hello"[2])

}
