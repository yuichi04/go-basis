package main

import (
	"fmt"
	"go-basis/08_variables/zero_value"
)

type OS int

// 定数は連番と一緒に使用されることが多い
const (
	Mac OS = iota + 1
	Windows
	Linux
)

func main() {
	zero_value.PrintZeroValue()

	fmt.Printf("Mac:%v Windows:%v Linux:%v\n", Mac, Windows, Linux)
}
