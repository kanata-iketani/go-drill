// lesson05 ch02 解答
package main

import "fmt"

func main() {
	// "Go" は 1 バイト × 2 文字、"言語" は 3 バイト × 2 文字なので
	// 2 + 6 = 8 バイトです。len が「文字数」ではなく「バイト数」を
	// 返すことを、日本語入りの文字列で体感するのがこの演習の狙いです。
	fmt.Println(len("Go言語"))
}
