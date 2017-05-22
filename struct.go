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

type Vertex struct {
	X int
	Y int
}

func main() {
    p := pair{3, 4}
    fmt.Println(p.String()) // 调用 pair 类型 p 的 string 方法

    v1 := Vertex{1, 2} // has type Vertex
    v2 := Vertex{X: 1} // Y:0 is implicit
    v3 := Vertex{}     // X:0 and Y:0
    q  := &Vertex{1, 2} // has type *Vertex

    fmt.Println(v1,  v2, v3, q)
}
