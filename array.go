package main

import (
    "fmt"
    "strings"
)

func main()  {
    var a[2]string
    a[0] = "Hello"
    a[1] = "World"

    fmt.Println(a[0], a[1]) // 输出：Hello World
    fmt.Println(a)          // 输出：[Hello World]

    primes := [6]int{2, 3, 5, 7, 11, 13}
    fmt.Println(primes) // 输出：[2 3 5 7 11 13]

    var s []int = primes[1:4]
    fmt.Println(s) // 输出：[3 5 7]

    names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
    fmt.Println(names)

    c := names[0:2]
    d := names[1:3]
    fmt.Println(c, d)

    d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

    sliceLiterals()
    creatingASliceWithMake()
    slicesOfSlices()
    appendingToASlice()
}

func sliceLiterals()  {
    fmt.Println("\n== sliceLiterals")

    q := []int{2, 3, 5, 7, 11, 13}
    fmt.Println(q)

    r := []bool{true, false, true, true, false, true}
    fmt.Println(r)

    ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(ss)
}

func creatingASliceWithMake()  {
    fmt.Println("\n== creatingASliceWithMake")

    a := make([]int, 5)
    printSlice("a", a)

    b := make([]int, 0, 5)
    printSlice("b", b)

    c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func slicesOfSlices()  {
    fmt.Println("\n== slicesOfSlices")

    // Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

    // The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

    for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendingToASlice()  {
    fmt.Println("\n== appendingToASlice")

    var s []int
    printSlice2(s)

    // append works on nil slices.
	s = append(s, 0)
	printSlice2(s)

    // The slice grows as needed.
	s = append(s, 1)
	printSlice2(s)

    // We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice2(s)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
