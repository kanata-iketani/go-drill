// lesson13 ch02 解答
package main

import "fmt"

func main() {
	// range 3 は 0, 1, 2 と回るので、1回目〜3回目にするには i+1 を使います。
	// Python の range(1, 4) のような開始値の指定は Go にはありません。
	for i := range 3 {
		fmt.Printf("%d回目\n", i+1)
	}
}
