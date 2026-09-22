// lesson02 ch02 解答
package main

import "fmt"

func main() {
	// Print は改行しないため、呼び出しを重ねると 1 行につながります。
	// 最後の 1 回だけ Println にして行を閉じるのが定番の形です。
	fmt.Print("Go")
	fmt.Print("lang")
	fmt.Println("!")
}
