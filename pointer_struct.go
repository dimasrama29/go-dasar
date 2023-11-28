package main 

import "fmt"

type Students struct {
  Name string 
  Age int
}

func main() {
  var s1 = Students {
    Name: "Dimas",
    Age: 24,
  }

  var s2 *Students = &s1

  fmt.Println(s1.Name)
  fmt.Println(s2.Name)

  s2.Name = "Rama"
  fmt.Println(s1.Name)
  fmt.Println(s2.Name)

}

/*
* Property pada object pointer tidak perlu di deferencing, nilai asli tetap bisa diakses*/
