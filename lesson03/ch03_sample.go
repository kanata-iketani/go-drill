// lesson03 ch03: math.Pow（** はない）
package main

import (
	"fmt"
	"math"
)

func main() {
	// Python の 2 ** 10 は math.Pow(2, 10) と書きます
	fmt.Println(math.Pow(2, 10))      // float64 で 1024
	fmt.Println(int(math.Pow(2, 10))) // int にしたいときは変換

	// 2乗くらいなら掛け算で書くほうが Go らしいコードです
	x := 9
	fmt.Println(x * x)

	// 平方根は math.Sqrt です
	fmt.Println(math.Sqrt(2))
}
