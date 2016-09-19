// 1. 为什么需要 interface ?
// 答：优先选择组合而非继承（设计模式原则）
//
// 2. 怎么去做选择？一定要做选择么？
// 答：一个程序只做一件事，并且做好它（UNIX 编程艺术）

package main

import (
    "fmt"
    "math"
)

// 定义一个 interface
type geometry interface {
    area() float64
}

// 定义一个 rect 结构体
type rect struct {
    width, height float64
}

// 定义一个 circle 结构体
type circle struct {
    radius float64
}

// 实现 rect 的 area() 接口
func (r rect) area() float64 {
    return r.width * r.height
}

// 实现 circle 的 area() 接口
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}
