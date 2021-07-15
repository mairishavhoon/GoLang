package main

import "fmt"

func main() {
  var a = 90
  switch a {
    case 80 : fmt.Println("Its 80")

    case 90 : fmt.Println("Its 90")

    default : fmt.Println("I dont know main")
  }
  fmt.Printf("a's value ij %d\n", a );
}