package main

// import "fmt"

type T0 struct {
  T1
}

type T1 struct {
  T0
}

func main() {
  t0 := T0()

}
