// lesson12 ch03 解答
package main

import "fmt"

func main() {
	// i++ を本体の先頭に置いてから continue します。
	// continue の後ろに更新を書くと、スキップのたびに i が増えず
	// 無限ループになるため、この順番が重要です。
	i := 0
	for i < 6 {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
