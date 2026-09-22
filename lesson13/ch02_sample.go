// lesson13 ch02: for i := range 10（range(n) の代わり）
package main

import "fmt"

func main() {
	// Go 1.22 以降の書き方。0 から 2 まで回ります
	for i := range 3 {
		fmt.Println(i)
	}

	// Go 1.21 以前はこの 3 部形式で書いていました（今も使えます）
	for i := 0; i < 3; i++ {
		fmt.Println("old:", i)
	}

	// range(1, 4) のような開始値の指定はないので i+1 で調整します
	for i := range 3 {
		fmt.Println(i + 1)
	}
}
