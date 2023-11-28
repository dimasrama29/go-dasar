package main 

import (
  "fmt"
  "strings"
)

type Students struct {
  Name string
  Age int
}

func (s Students) sayHello() {
  fmt.Println("Halo", s.Name) // fungsi sayHello dideklarasikan sebagai method struct Students
}

func (s Students) getNameAt(i int) string {
  return strings.Split(s.Name, " ")[i - 1]
}

func main() {
  var s1 = Students {
    Name: "Dimas Setyawan Ramadhansyah",
    Age: 24,
  }
  s1.sayHello()
  fmt.Println(s1.getNameAt(2))
} 

/*
* Method: Fungsi yang menempel pada type
*   Keunggulan: Bisa diakses hingga level private*/
