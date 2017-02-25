package main

import "fmt"

type person struct {
  fname string // lowercase nt visible outside package uppercase for visible outside pkg
  lname string
}

func main() {
  xi := []int{2, 4, 7, 9, 42}
  fmt.Println(xi)

  m := map[string]int{
    "Todd": 45,
    "Frank": 33,
  }
  fmt.Println(m)

  p1 := person{
      "Brusif",
      "Franklernisher",
  }

  fmt.Println(p1)
}