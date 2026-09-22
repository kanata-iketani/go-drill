// lesson09 ch01: Go にタプルはない（多値返却で代替）
package main

import "fmt"

// 商と余りを返します。戻り値の型を ( ) で並べるのが多値返却です
func divmod2(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	q, r := divmod2(17, 5)
	fmt.Println(q)
	fmt.Println(r)

	// タプルという 1 つの値ではないため、受け取る変数は個数分必要です
	fmt.Println(divmod2(9, 2))
}
