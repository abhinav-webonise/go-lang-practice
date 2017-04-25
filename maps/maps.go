package main

import "fmt"

func main() {
  m := make(map[string] int)
  m["k1"] = 7
  m["k2"] = 13

  fmt.Println("map", m)

  v1 := m["k1"]
  fmt.Println("v1:", v1)

  fmt.Println("len:", len(m))

  delete(m, "k2")
  fmt.Println("map:", m)

  _, prs := m["k2"]

  // Basically what it does is _ will skip it's value  and prs will return whether that value exists in that map or not based on that it will retun true or false

  fmt.Println("prs:", prs)
  n := map[string]int {"foo": 1, "bar": 2}
  fmt.Println("map", n)
}
