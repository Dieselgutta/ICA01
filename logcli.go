package main

import (
  "fmt"
  "math"
  )

func main() {
  fmt.Print("Number: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := math.Log2(input)

  fmt.Println(output)
}
