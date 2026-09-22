// lesson05 ch04: strings.Repeat（文字列の掛け算の代わり）
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Python の "=" * 20 はこう書きます
	fmt.Println(strings.Repeat("=", 20))

	fmt.Println("メニュー")

	// 複数文字の繰り返しもできます
	fmt.Println(strings.Repeat("-*", 10))

	// 変数と組み合わせて区切り線を作るのが定番の使い方です
	line := strings.Repeat("=", 20)
	fmt.Println(line)
}
