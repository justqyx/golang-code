package main

import (
    "fmt"
)

// func main() {
//     c := make(chan int)
//
//     fmt.Println("send ...")
//
//     c <- 1
//
//     fmt.Println("reciver ...")
//
//     n := <-c
//     fmt.Println(n)
// }

// 只有一个 channel 的情况下，无法同时压入和取出
// Result
//
// send ...
// fatal error: all goroutines are asleep - deadlock!
//
// goroutine 1 [chan send]:
// main.main()
// 	/Users/yuanxin/Learn/golang/channel.go:12 +0x146
// exit status 2

// func send(c chan int, num int)  {
//     fmt.Println("send ...")
//     c <- num
// }

func main() {
    c := make(chan int, 2)

    // go send(c, 3)
    // go send(c, 4)

    // // 同步方法
    fmt.Println("send ...")
    c <- 2
    //
    // // 异步方法
    // // go is a statement that starts a new goroutine
    go func () {
        fmt.Println("send ...")
        c <- 1
    }()

    fmt.Println("recive ...")
    fmt.Println(<-c, <-c)
}

// 一个 channel 用来发送，一个用来接收
// Reuslt
//
// send ...
// recive ...
// send ...
// 1 1
