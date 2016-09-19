package main

import (
    "fmt"
)

func add(x int, y int) int {
    return x + y
}

func swap(x, y string) (string, string)  {
    return y, x
}

func intSeq() func() int {
    i := 0

    return func () int {
        i += 1
        return i
    }
}

func partialSum(x int) func(int) int {
    // 返回一个匿名函数
    return func (y int) int  {
        return x + y
    }
}

func main() {
    fmt.Println(add(42, 13))
    fmt.Println(swap("Hello", "world"))

    message := "Hello golangcode.com readers!"

    // Lambda 语法
    func (m string)  {
        fmt.Println(m)
    }(message)

    nextInt := intSeq()
    fmt.Println(nextInt()) // 1
    fmt.Println(nextInt()) // 2
    fmt.Println(nextInt()) // 3

    // 结合匿名函数实现柯里化
    partial := partialSum(3)
    fmt.Println(partial(4))       // prints 7
    fmt.Println(partialSum(3)(4)) // prints 7
}
