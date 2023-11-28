package main 

import "fmt"

type Students struct {
  Name string 
  Age int
}

func (s Students) changeName1 (name string) {
  fmt.Println(name)
  s.Name = name 
}

func (s *Students) changeName2 (name string) {
  fmt.Println(name)
  s.Name = name 
}

func main() {
  var s1 = Students {
    Name: "Dimas",
    Age: 24,
  }

  fmt.Println(s1.Name)

  s1.changeName1("Setyawan")
  fmt.Println("Change Name1:", s1.Name) // Tidak berubah

  s1.changeName2("Rama")
  fmt.Println("Change Name2:", s1.Name)
}
