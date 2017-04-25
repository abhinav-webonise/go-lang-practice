package main

import "fmt"

func vals() (int, int) {
  return 3, 7
}

func add ( a int, b int) int {
  return a + b
}

func addThreeNumbers (a, b , c int) int {
  return a + b + c
}

func sum(nums ...int) {
    fmt.Println(nums, " ")
    total := 0
    for _, num := range nums {
      total += num
    }
    fmt.Println(total)
}

func main() {
  res := add (1, 2)

  fmt.Println(res)

  res = addThreeNumbers (1, 2, 3)
  fmt.Println(res)

  // multiple retun values
  a, b := vals()
  fmt.Println(a, b)

  _, c := vals()
  fmt.Println(c)

  // variadic functions
  sum(1, 2, 3, 4)
  sum(5, 6)

  nums := []int {1, 5, 3, 8}
  sum(nums...)

}
