// lesson15 ch01: bufio.Scanner で 1 行読む
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Python の input() に相当する処理は Go では複数行かかります
	sc := bufio.NewScanner(os.Stdin) // 1. 標準入力を読む Scanner を作る
	sc.Scan()                        // 2. 1 行読み込む
	line := sc.Text()                // 3. 読んだ行を文字列で取り出す（改行は含まない）

	fmt.Println("入力:", line)
}
