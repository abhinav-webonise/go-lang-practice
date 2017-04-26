package main

import "fmt"

type person struct {
  name string
  age int
}

type rect struct {
  width, height int
}

func(r *rect) area() int {
  return r.width * r.height
}

func (r rect) perimeter() int {
  return 2 *(r.width + r.height)
}

func main()  {
    fmt.Println(person{"Bob", 20})
    // fmt.Println(person{"Abhinav"}) // this is not a valid statement
    // fmt.Println(person{22})  // -do-

    fmt.Println(person{name: "Varun", age: 30})
    fmt.Println(person{name: "Afridi"})
    fmt.Println(person{age: 40})

    fmt.Println(&person{name: "Anshu", age: 29})
    fmt.Println(&person{"Taha", 23})

    s := person{name: "Jay", age: 50}
    fmt.Println(s.name)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(s.age)        // Structs are mutable

    // Methods
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perimeter: ", r.perimeter())

    rp := &r

    fmt.Println("area: ", rp.area())
    fmt.Println("perimeter: ", rp.perimeter())
}
