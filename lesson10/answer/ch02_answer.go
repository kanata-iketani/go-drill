// lesson10 ch02 解答
package main

import "fmt"

func main() {
	n := 7

	// 条件は上から順に評価されるため、範囲の狭い方（>= 10）から先に書きます。
	// elif ではなく else if で、直前の } と同じ行に続けるのが Go の書き方です。
	if n >= 10 {
		fmt.Println("big")
	} else if n >= 5 {
		fmt.Println("medium")
	} else {
		fmt.Println("small")
	}
}
