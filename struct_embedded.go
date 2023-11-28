package main 

import "fmt"

type Person struct {
  Name string 
  Age string
}

type Student struct { 
  Grade int
  Age int
  Person // Embedded struct 

}

func main() {
  var s1 = Student{}

  s1.Name = "Dimas"
  s1.Age = 24
  s1.Person.Age = "25"
  s1.Grade = 3

  fmt.Println(s1.Name)
  fmt.Println(s1.Age) // Age Student
  fmt.Println(s1.Person.Age) // Harus jelas jika memiliki property yang sama (Age Person)
  fmt.Println(s1.Grade)

  // Pengisian sub-struct
  var p2 = Person{Name: "Ira", Age: "21"}
  var s2 = Student{Grade: 2, Age: 21, Person: p2}

  fmt.Println(p2.Name)
  fmt.Println(s2.Person)
}

/*
* Embedded: Tempel struct sebagai property struct lain*/
