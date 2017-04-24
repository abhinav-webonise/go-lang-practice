package main

import "fmt"
import "math"

const s string = "constantstring"

func main() {
  // Hello World

  fmt.Println("Hello Abhinav!")

  // Values

  fmt.Println("abhi" + "nav")
  fmt.Println(1 + 1) // 2
  fmt.Println(7.0/3.0) // 2.3333
  fmt.Println(8.0/2.0) // 4
  fmt.Println(7/3) // 2

  // variables

  var a string = "initial"
  fmt.Println(a)

  var b,c int = 1, 2
  fmt.Println(b, c)

  var d = true
  fmt.Println(d)

  var e int
  fmt.Println(e) // 0 Zero Valued

  var z string
  fmt.Println(z) // "" empty string

  var x bool
  fmt.Println(x) // false is zero valued in case of boolean

  y, k := "shorthand", 1
  fmt.Println(y, k)

  // Constants

  fmt.Println(s)

  const n = 5500000

  const f = 3e20 / n
  fmt.Println(f)

  fmt.Println(int64(f))

  fmt.Println(math.Sin(n))

}
