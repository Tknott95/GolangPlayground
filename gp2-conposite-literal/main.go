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

func (p person) speak() {
    fmt.Println(p.fname, p.lname, `says, "Good morning, Trevor Knott`)
}

func (sa secretAgent) speak() {
    fmt.Println(sa.fname, sa.lname, `says, "Kick right to your face.. SNAP`)
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

  sa1 := secretAgent{
      person{
          "Juade Claude",
          "Van Dam",
      },
      true,
  }

  p1.speak()
  sa1.speak()
  sa1.person.speak()
  fmt.Println(p1)

}