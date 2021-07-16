package main

import ("fmt"; "math")

func main() {
  var res int = sub(9,4)
  fmt.Printf("The difference between 9 and 4 is %d\n" ,res)
  sqt := func(x float64) float64{
    return math.Sqrt(x)
  }
  var s9 float64 = sqt(9)
  fmt.Printf("square root of 9 is %f\n" ,s9)
}

func sub(a,b int) int {
  var c int = a-b
  return c;
}

//the sqt() function uses function variable