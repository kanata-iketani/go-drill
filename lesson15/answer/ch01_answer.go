// lesson15 ch01 解答
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// input() 相当の定型: NewScanner → Scan → Text。
	// Text() は改行を除いた 1 行分の文字列を返すので、そのまま連結に使えます。
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	name := sc.Text()

	fmt.Printf("こんにちは、%sさん\n", name)
}
