// lesson11 ch02 解答
package main

import "fmt"

func main() {
	x := 7
	y := 12

	// and / or / not をそのまま && / || / ! に置き換えます。
	// ! は直後の式だけを否定するので、対象全体を () で囲むと意図が明確です。
	fmt.Println(x > 5 && y > 10)
	fmt.Println(x > 10 || y > 10)
	fmt.Println(!(x == 7))
}
