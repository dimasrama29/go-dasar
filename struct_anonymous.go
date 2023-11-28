package main 

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
  // Anonymous struct tanpa nilai property
  var s1 = struct {
    Person 
    Grade int
  }{}

  s1.Grade = 2
  s1.Person = Person {Name: "Dimas", Age: 24}

  // Anonymous struct dengan nilai property
  var s2 = struct {
    Person 
    Grade int
  }{
    Grade: 3,
    Person: Person{Name: "Ira", Age: 21},
  }

  fmt.Println(s1)
  fmt.Println(s2)

  // Anonymous dengan var | Deklarasi
  var mahasiswa struct {
    Person 
    Grade int
  }

  mahasiswa.Person = Person{"Rama", 22}
  mahasiswa.Grade = 2

  // Inisialisasi
  var mhs2 = struct {
    Grade int 
  }{
    3,
  }

  fmt.Println(mhs2)
}
