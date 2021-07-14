/* testing operators in golang */
package main

import "fmt"

func main() {
  var a int = 10
  var b = 2
  fmt.Println("Sum: ", a+b)
  fmt.Println("Diff: ", a-b)
  fmt.Println("Product: ", a*b)
  fmt.Println("Quotient: ", a/b)
  fmt.Println("Remainder: ", a%b)
}