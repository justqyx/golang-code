package main

import (
    "fmt"
)

// 定义pair为一个结构体，有x和y两个int型变量。
type pair struct {
    x, y int
}

// 定义pair类型的方法，实现Stringer接口。
func (p pair) String() string {
    return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func main() {
    p := pair{3, 4}
    fmt.Println(p.String()) // 调用 pair 类型 p 的 string 方法
}
