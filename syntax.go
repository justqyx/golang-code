package main

import (
    "fmt"
    "math"
)

const MIN = 0.000001

// 浮点数不是精确的表达方式，所以比较使用 == 是不可行的
func IsEqual(f1, f2 float64) bool {
    return math.Dim(f1, f2) < MIN
}

func PrintArray(arr []int) {
    for _, v := range arr {
        fmt.Print(v, " ")
    }

    fmt.Println("")
}

func main() {
	fmt.Println("hello, world.")

  // ---------------------------
  // Boolean
  var b1 bool = true
  b2 := (1 == 2)
  fmt.Println(b1, b2)

  // ---------------------------
  // Integer
  var (
      v1 int8
      v2 uint8
      v3 int16
      v4 uint16
      v5 int32
      v6 uint32
      v7 int64
      v8 uint64
      v9 int      // 平台相关
      v10 uint    // 平台相关
      v11 uintptr // 32 位平台下为4字节，64 位为8字节
  )

  // 以上在 Go 语言里都被认为是不同的类型
  // 并且编译器不会自动转换
  fmt.Println("int8   : %d", v1)
  fmt.Println("uint8  : %d", v2)
  fmt.Println("int16  : %d", v3)
  fmt.Println("uint16 : %d", v4)
  fmt.Println("int32  : %d", v5)
  fmt.Println("uint32 : %d", v6)
  fmt.Println("int64  : %d", v7)
  fmt.Println("uint64 : %d", v8)
  fmt.Println("int    : %d", v9)
  fmt.Println("uint   : %d", v10)
  fmt.Println("uintptr: %d", v11)

  // ---------------------------
  // Float
  var (
      f1 float32  // 等价与 C 语言的 float 类型
      f2 float64  // 等价与 C 语言的 double 类型
      f3 float32
  )

  f1 = 12.001
  f2 = 12.0
  f3 = float32(f2)    // 强制类型转换

  fmt.Println("f1: ", f1)
  fmt.Println("f2: ", f2)
  fmt.Println("f3: ", f3)

  result := IsEqual(float64(f1), f2)
  fmt.Println("f1 < f2 = ", result)

  // ---------------------------
  // Complex 符合类型
  var complex1 complex64 // 由两个 float32 类型构成的复数类型
  complex1 = 3.2 + 12i
  complex2 := 3.2 + 12i
  complex3 := complex(3.2, 12)
  fmt.Println(complex1, complex2, complex3)

  // ---------------------------
  // String 字符串
  var str1 string

  str1 = "Hello World"
  str2 := str1[0]

  fmt.Printf("The length of \"%s\" is %d \n", str1, len(str1))
  fmt.Printf("The first character of \"%s\" is %c.\n", str1, str2)

  fmt.Println("## 字符串遍历，以字节数组遍历")
  string1_len := len(str1)

  for i := 0; i < string1_len; i++ {
      ch1 := str1[i]
      fmt.Println(i, ch1)
  }

  fmt.Printf("## 以 unicode 字符遍历\n")

  for i, ch2 := range str1 {
      fmt.Println(i, ch2)
  }

  // ---------------------------
  // Array
  // Slice 可以动态增删，数组在定义的时候，必须制定起大小，而切片可以只是指定元素而不指定大小
  var arr1 [5]int = [5]int{1,2,3,4,5} // 数组
  var arr2 []int = arr1[:3]           // 基于 arr1 的切片
  var arr3 = make([]int, 5)           // 创建元素个数为 5 的数组切片
  var arr4 = make([]int, 5, 10)       // 创建 5 个初始值为 0 的数组切片，并预留 10 个元素的存储空间
  var arr5 = []int{1,2,3,4,5}         // 创建包含 5 个元素的数组切片

  fmt.Println("len(arr4): ", len(arr4))
  fmt.Println("cap(arr4): ", cap(arr4))

  PrintArray(arr2)

  fmt.Println("-- arr3")
  PrintArray(arr3)
  arr3 = append(arr3, 6, 7)
  PrintArray(arr3)

  fmt.Println("-- arr4")
  PrintArray(arr4)
  arr4 = append(arr4, arr5...)
  PrintArray(arr4)

  fmt.Println("-- arr5")
  PrintArray(arr5)
  arr5 = append(arr5, 6, 7, 8)
  PrintArray(arr5)
}
