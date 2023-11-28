package main 

import "fmt"

type Person struct {
  Name string 
  Age int
}

func main() {
  var allStudents = []struct {
    Person
    Grade int 
  } {
    {Person: Person{Name: "Dimas", Age: 24}, Grade: 2}, 
    {Person: Person{Name: "Rama", Age: 22}, Grade: 3}, 
  }

  for _, v := range allStudents {
    fmt.Println(v)
  }
}
