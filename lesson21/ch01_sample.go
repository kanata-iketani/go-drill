// lesson21 ch01: type X struct（class の代わり）
package main

import "fmt"

// Person は名前と年齢をひとまとめにした構造体です
type Person struct {
	Name string
	Age  int
}

func main() {
	// フィールド名: 値 の形で初期化します
	p := Person{Name: "佐藤", Age: 28}
	fmt.Println(p.Name, p.Age)

	// フィールドは後から読み書きできます
	p.Age = 29
	fmt.Println(p.Age)

	// 指定しなかったフィールドはゼロ値（"" や 0）になります
	q := Person{Name: "田中"}
	fmt.Println(q.Name, q.Age)
}
