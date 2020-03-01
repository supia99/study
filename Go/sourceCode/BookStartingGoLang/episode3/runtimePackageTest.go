package main

import (
  "fmt"
  "runtime"
)

func runtimePackageTest() {
  go fmt.Printf("Yeah!")
  fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
  fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
  fmt.Printf("Version: %s\n", runtime.Version())

}


// func main() {
//
// }
