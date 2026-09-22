// lesson11 ch04: strings.Contains（文字列の in の代わり）
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "golang is fun"

	// "go" in s の代わりに strings.Contains を使います
	fmt.Println(strings.Contains(s, "go"))
	fmt.Println(strings.Contains(s, "python"))

	// 前方一致・後方一致には専用の関数があります
	fmt.Println(strings.HasPrefix(s, "golang"))
	fmt.Println(strings.HasSuffix(s, "fun"))
}
