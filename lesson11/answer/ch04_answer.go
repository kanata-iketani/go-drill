// lesson11 ch04 解答
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "golang is fun"

	// Python の "go" in s の代わりです。対象が文字列なので、
	// slices ではなく strings パッケージの Contains を選びます。
	fmt.Println(strings.Contains(s, "go"))
	fmt.Println(strings.Contains(s, "python"))
}
