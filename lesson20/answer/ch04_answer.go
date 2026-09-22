// lesson20 ch04 解答
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Go にデフォルト引数はないので、省略可能な設定は構造体にまとめます。
// 渡さなかったフィールドはゼロ値（"" や 0）になるため、
// 「ゼロ値ならデフォルトを補う」ことで prefix="Hello" の代わりになります。
type Options struct {
	Prefix string
}

func greet(name string, opt Options) {
	if opt.Prefix == "" {
		opt.Prefix = "Hello"
	}
	fmt.Printf("%s, %s!\n", opt.Prefix, name)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	name := strings.Fields(sc.Text())[0]
	sc.Scan()
	prefix := strings.Fields(sc.Text())[0]

	if prefix == "-" {
		greet(name, Options{}) // 指定なし: デフォルトが使われる
	} else {
		greet(name, Options{Prefix: prefix})
	}
}
