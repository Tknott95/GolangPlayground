package main

import "fmt"

type person struct {
  fname string // lowercase nt visible outside package uppercase for visible outside pkg
  lname string
}

type secretAgent struct {
  person
  lic2Kill bool
}

// func (p person) speak() {
//     fmt.Println(p.fname, p.lname, `says, "Good morning, Trevor Knott`)
// }

func (sa secretAgent) speak() {
    fmt.Println(sa.fname, sa.lname, `says, "Good morning, Trevor Knott`)
}

func main() {
//   xi := []int{2, 4, 7, 9, 42}
//   fmt.Println(xi)

//   m := map[string]int{
//     "Todd": 45,
//     "Frank": 33,
//   }
//   fmt.Println(m)

  p1 := person{
      "Brusif",
      "Franklernisher",
  }

  fmt.Println(p1)

//   p1.speak()

  sa1 := secretAgent{
      person{
          "Juade Claude",
          "Van Dam",
      },
      true,
  }
  sa1.speak()
}