package main 

import "fmt"

type Person struct {
  Name string 
  Age int
}

func main() {
  var allStudents = []Person {
    {Name: "Dimas", Age: 24},
    {Name: "Setyawan", Age: 23},
    {Name: "Ramadhansyah", Age: 22},
  }

  for _, v := range allStudents {
    fmt.Println(v)
  }
  
}
