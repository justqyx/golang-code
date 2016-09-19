package main

import (
    "fmt"
)

func main() {

    var x map[string]int
    fmt.Println(x) // map[]

    y := map[string]int{"age": 30, "face": 80}
    fmt.Println("face =", y["face"]) // face = 80

    // 初始化 size
    z := make(map[string]string, 10)
    z["name"] = "Pike"
    fmt.Println(z) // map[name:Pike]

    // 不确定 key 访问
    age, ok := z["age"]
    fmt.Printf("age=%s, age exist=%v\n", age, ok) // age=, age exist=false

    // 遍历
    for k, v := range y {
        fmt.Printf("k=%v, v=%v \n", k, v)
        // k=age, v=30
        // k=face, v=80
    }

    // 删除
    delete(y, "face")
    fmt.Println(y) // map[age:30]
}
