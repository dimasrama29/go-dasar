package main 

import (
  "fmt"
  "math"
)

type hitung interface {
  keliling() float64
  luas() float64
}

type lingkaran struct {
  diameter float64
}

func (l lingkaran) jariJari() float64 {
  return l.diameter / 2
}

func (l lingkaran) keliling() float64 {
  return math.Pi * l.diameter 
}

func (l lingkaran) luas() float64 {
  return math.Pi * math.Pow(l.jariJari(), 2)
}

func main() {
  var bangunDatar hitung 
  bangunDatar = lingkaran{70}

  fmt.Println(bangunDatar.keliling())
  fmt.Println(bangunDatar.luas())
}

/*
* Interface: Kumpulan definisi method yang tidak memiliki isi
* Interface adalah tipe data
*   Zero value: nil*/
