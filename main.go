package main

import (
  "fmt"
  "math"
  "math/rand"
  "math/cmplx"
)

func add(x, y int) int {
  return x + y
}

func swap(x, y string) (string, string) {
  return y,x
}

// Named return values
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println("My favorite number is", rand.Intn(10))
  fmt.Println(math.Pi)
  fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

  fmt.Println(add(42, 13))

  a, b := swap("hello", "world")
  fmt.Println(a, b)

  fmt.Println(split(17))

  // The var statement declares a list of variables; as in function argument lists, the type is last.
  var c, python, java bool
  var i int

  fmt.Println(i, c, python, java)

  var (
    ToBe   bool       = false
    MaxInt uint64     = 1<<64 - 1
    z      complex128 = cmplx.Sqrt(-5 + 12i)
  )

  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value: %v\n", z, z)
}
