package main

import "fmt"

func main()  {

    nums := []int{2, 3, 4}

    sum := 0

    for _, nums := range nums {
      sum += nums
    }
    fmt.Println("sum: ", sum)

    for i, nums := range nums {
      if nums == 3 {
        fmt.Println("index: ", i)
      }
    }

    kvs := map[string]string {"a": "apple", "b": "banana"}
    for k,v := range kvs {
      fmt.Println(k, ":", v)
    }

    for k:= range kvs {
      fmt.Println("key:", k)
    }

    for i,c := range "go" {
      fmt.Println(i, c)
    }
}
