package main

import "fmt"
import "math"
import "time"

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

  const n = 500000000

  const f = 3e20 / n
  fmt.Println(f)

  fmt.Println(int64(f))

  fmt.Println(math.Sin(n))

  // FOR LOOP

  // Syntax similar to while loop
  i := 1
  for i<=3 {
    fmt.Println(i)
    i = i + 1
  }

  for j:= 0; j <= 9; j++ {
    fmt.Println(j)
  }

  // infinite loop if break statement isn't present here
  for {
    fmt.Println("loop")
    break
  }

  for l := 0; l < 15; l++ {
    if l%2 == 0 {
      continue
    }
    fmt.Println(l)
  }

  // IF-ELSE
  // Switch

  swi := 2
  fmt.Println("Write", swi, " as ")
  switch swi {
    case 1:
      fmt.Println("one")
    case 2:
      fmt.Println("two")
    case 3:
      fmt.Println("three")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's Weekend")
  default:
    fmt.Println("It's Weekday")
  }

  current_time := time.Now()
  switch  {
  case current_time.Hour() < 12:
    fmt.Println("It's before Noon")
  default:
    fmt.Println("It's after noon")
  }

  whatAmI := func (i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("I am Bool type")
    case int:
      fmt.Println("I am int type")
    default:
      fmt.Println("Don't know type \n", t) // returns string name
      fmt.Printf("Don't know type %T\n", t) // returns type of format
    }
  }

  whatAmI(true)
  whatAmI(55)
  whatAmI("hello")

  //fmt.Println(time.Now().Day())  // Playing with time object is fun :)
  //fmt.Println(time.Saturday)    // time.Saturday is constant in time class


  // Arrays

  var arr [5]int
  fmt.Println("emp:", arr)

  arr[4] = 100
  fmt.Println("set:", arr)
  fmt.Println("get:", arr[4])

  fmt.Println("len:", len(arr))

  barr := [5]int{1,2,3,4,5}
  fmt.Println("dcl:", barr)

  //  barr1 := [5]int{1}       // If not initialised then auto set to zero
  //  fmt.Println("dcl1:", barr1)


  var twoD [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0;j < 3; j++ {
        twoD[i][j] = i + j
    }
  }
  fmt.Println("2D", twoD)

}
