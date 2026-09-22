// lesson20 ch04: キーワード引数・デフォルト引数はない
package main

import "fmt"

// 省略できる設定は構造体にまとめます（オプション構造体）
type Options struct {
	Prefix string
}

func greet(name string, opt Options) {
	// 指定がなければ（ゼロ値 "" なら）デフォルトを補います
	if opt.Prefix == "" {
		opt.Prefix = "Hello"
	}
	fmt.Printf("%s, %s!\n", opt.Prefix, name)
}

func main() {
	greet("Taro", Options{})             // デフォルト引数の代わり
	greet("Hana", Options{Prefix: "Hi"}) // キーワード引数の代わり
}
