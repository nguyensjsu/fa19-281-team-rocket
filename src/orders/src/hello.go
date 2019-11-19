package main

import (
	"errors"
	"fmt"
	"math"
)

//func main() {
//fmt.Println("Hello, world")

// x := 5
// y := 4
// sum := x + y
// fmt.Println(sum)

// x := 6
// if x >= 6 {
// 	fmt.Println("greater than 6")
// }

// var a []int
// a = append(a, 5)
// a = append(a, 5)
// a = append(a, 5)
// fmt.Println(a)

// vertices := make(map[string]int)
// vertices["triangle"] = 3
// vertices["square"] = 4
// fmt.Println(vertices)

// for i := 0; i < 5; i++ {
// 	fmt.Println(i)
// }

// arr := []string{"a", "b", "c"}
// for index, value := range arr {
// 	fmt.Println(index, value)
// }
// 	result, err := sqrt(3)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}
// }

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("undefined for negaive no")
	}
	fmt.Println(x)
	return math.Sqrt(x), nil
}
