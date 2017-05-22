package main

import (
    "fmt"
    "golang.org/x/tour/wc"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
    printMap()
    literalMap()
    mapLiteralsContinued()
    mutatingMaps()
    testWorldCount()
}

func printMap()  {
    var m map[string]Vertex

    // 创建一个 slice
    m = make(map[string]Vertex)

    m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

    m["a"] = Vertex{
		40.68433, -74.39967,
	}

    m["b"] = Vertex{
		40.68433, -74.39967,
	}

    fmt.Println(m["Bell Labs"])
    fmt.Println(m)
}

func literalMap()  {

    // "string" => Vertex

    var m = map[string]Vertex{
    	"Bell Labs": Vertex{ 40.68433, -74.39967, },
    	"Google": Vertex{ 37.42202, -122.08408, },
    }

    fmt.Println(m)
}

func mapLiteralsContinued()  {
    var m = map[string]Vertex{
    	"Bell Labs": {40.68433, -74.39967},
    	"Google":    {37.42202, -122.08408},
    }

    fmt.Println(m)
}

func mutatingMaps()  {
    m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func testWorldCount()  {
    wc.Test(WordCount)
}
