package zero_value

import "fmt"

var (
	i   int
	s   string
	f32 float32
	f64 float64
	a   []int
	b   bool
	m   map[string]int
	p   *int
)

func PrintZeroValue() {
	fmt.Println("int:", i)
	fmt.Println("string:", s)
	fmt.Println("float32:", f32)
	fmt.Println("float64:", f64)
	fmt.Println("[]int:", a)
	fmt.Println("bool:", b)
	fmt.Println("map[string]int:", m)
	fmt.Println("*int:", p)
}
