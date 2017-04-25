package main

import "fmt"

func zeroVal(ival int) {
  ival = 0
  fmt.Println("I am inside Function: ", ival)
}

func zeroPtr(iptr *int) {
  *iptr = 0
  fmt.Println("I am inside Pointer function1: ", iptr) // Pointing to the address of i
  fmt.Println("I am inside Pointer function2: ", &iptr) // Pointing to address of iptr
  fmt.Println("I am inside Pointer function3: ", *iptr) // Value at iptr
}

func main() {
  i := 1
  fmt.Println("initially: ", i)

  zeroVal(i)
  fmt.Println("After coming out of func", i)

  zeroPtr(&i)
  fmt.Println("After coming out of pointer func", i)

  fmt.Println("Address of ", &i)

}
