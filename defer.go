package main

import (
    "fmt"
)

func sum(x, y int) int {
    fmt.Println("sum func start ...")
    defer fmt.Println("sum func end ...")
    return x + y
}

func div(x, y int) int {
    // 捕捉该函数可能出现的错误，recover()
    defer func () {
        if err := recover(); err != nil {
            fmt.Println("recover", err)
        }
    }()

    return x / y
}

func main() {
    // Defer
    fmt.Println("main start ...")

    defer fmt.Println("main first defer ...")
    defer fmt.Println("main second defer ...")

    sum(1, 2)

    fmt.Println("main end ...")

    // Recover
    div(1, 0)
}

// Defer Result
//
// main start ...
// sum func start ...
// sum func end ...
// main end ...
// main second defer ...
// main first defer ...


// Recover Result
//
// recover runtime error: integer divide by zero

// 其他的应用场景
// 在 Web 开发中通过 middleware 的方式将错误异步地打印出来
//
// func (mw loggingMiddleware) Count(s string) (n int) {
//     defer func (begin time.Time) {
//         mw.logger.Log(
//             "method", "count",
//             "input", s,
//             "n", n,
//             "took", time.Since(begin)
//         )
//     }(time.Now())
//
//     n = mw.StringService.Count(s)
//     return
// }
