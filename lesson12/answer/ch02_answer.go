// lesson12 ch02 解答
package main

import "fmt"

func main() {
	// while True: に相当する for {} で回し、
	// 終了条件はループ本体の中の if + break で表します。
	total := 0
	for {
		total += 3
		if total >= 12 {
			break
		}
	}
	fmt.Println(total)
}
