// lesson24 ch03: %w で包んで errors.Is で判定する
package main

import (
	"errors"
	"fmt"
)

// 「原因」の目印になるエラーです（センチネルエラーと呼びます）
var ErrNotFound = errors.New("見つかりません")

func find(name string) error {
	if name == "gopher" {
		return nil
	}
	// %w で ErrNotFound を包み、文脈（誰を探したか）を足します
	return fmt.Errorf("検索 %s: %w", name, ErrNotFound)
}

func main() {
	err := find("alice")
	fmt.Println(err) // 検索 alice: 見つかりません

	// メッセージを足した後でも、errors.Is なら原因を判定できます
	fmt.Println(errors.Is(err, ErrNotFound)) // true
}
