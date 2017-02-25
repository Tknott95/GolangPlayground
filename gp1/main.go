package main

import "fmt";

var x  int

// in golang value of x here is 0 instead of undefined

func main() {
    x := 7 // Block level scope (sort of like a let in typescript)
    fmt.Println("%T", x)
    // fmt.Println("yo")
}