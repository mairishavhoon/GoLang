package main

import "fmt"

func main() {
  var a = 17

  if(a>20){
    if(a>20 && a<=25){
      fmt.Println("A is less or equal to 25")
    } else {
      fmt.Println("A is more than 25")
    }
  } else {
    fmt.Println("A is not more than 20 UwU")
  }
}