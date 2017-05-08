package main

import "fmt"
import "sort"

func main() {
  strs := []string{"c", "b", "d", "a"}
  sort.Strings(strs)
  fmt.Println("Strings: ", strs)

  ints := []int{4, 5, 1, 0, 2}
  sort.Ints(ints)
  fmt.Println("Ints: ", ints)

  s := sort.IntsAreSorted(ints)
  fmt.Println("Sorted: ", s)

}
