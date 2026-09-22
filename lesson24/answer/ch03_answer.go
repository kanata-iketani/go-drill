// lesson24 ch03 解答
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var ErrNotFound = errors.New("見つかりません")

func find(name string) error {
	if name == "gopher" {
		return nil
	}
	return fmt.Errorf("検索 %s: %w", name, ErrNotFound)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	name := sc.Text()

	err := find(name)
	if err == nil {
		fmt.Println(name, "が見つかりました")
		return
	}
	// %w で包んだので、メッセージには文脈が足されています。
	// それでも errors.Is なら「原因が ErrNotFound か」を判定できます。
	// エラーメッセージの文字列比較で判定しないのが Go の作法です。
	fmt.Println(err)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("原因は ErrNotFound です")
	}
}
