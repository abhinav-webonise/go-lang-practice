package main

import "fmt"
import "os"

type point struct {
  x,y int
}

var w = fmt.Printf

func main() {
  p := point{1, 2}

  w("%v\n", p)
  w("%+v\n", p)
  w("%#v\n", p)
  w("%T\n", p)
  w("%t\n", true)
  w("%d\n", 123)
  w("%b\n", 12)
  w("%c\n", 97)
  w("%x\n", 456)
  w("%f", 78.2)
  w("%e\n", 128500000.0)
  w("%E\n", 128500000.0)
  w("%s\n", "\"string\"" )
  w("%q\n", "\"string\"")
  w("%x\n", "hex me")
  w("%p\n", &p)
  w("|%-6d|%10d|\n", 23, 6000)
  w("|%-6.2f|%6.2f|\n", 52.6, 23.84545)
  w("|%-6s|%6s|\n", "bro", "abchde")

  s := fmt.Sprintf("a %s\n", "string")
  w(s)

  fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

// {1 2}
// {x:1 y:2}
// main.point{x:1, y:2}
// main.point
// true
// 123
// 1110
// !
// 1c8
// 78.900000
// 1.234000e+08
// 1.234000E+08
// "string"
// "\"string\""
// 6865782074686973
// 0x42135100
// |    12|   345|
// |  1.20|  3.45|
// |1.20  |3.45  |
// |   foo|     b|
// |foo   |b     |
// a string
// an error
