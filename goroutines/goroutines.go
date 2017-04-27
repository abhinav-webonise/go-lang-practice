package main

import "fmt"

func f(from string) {
    for i :=  0; i < 20; i++ {
      fmt.Println(from, ":", i)
    }
}

func main() {
  f("direct")

  go f("goroutine")

  go func(msg string) {
    fmt.Println(msg)
  }("going")

  go f("abhinav")

  var input string
  fmt.Scanln(&input)

  fmt.Println("done")
}


// go routines work asynchronously
