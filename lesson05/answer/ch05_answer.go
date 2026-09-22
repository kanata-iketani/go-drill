// lesson05 ch05 解答
package main

import "fmt"

func main() {
	item := "コーヒー"
	price := 480

	// f"{item}は{price}円です" の Go 版です。
	// 表示ではなく「文字列を作って変数に持つ」のが Sprintf の役割で、
	// あとで連結やファイル出力に使い回せるのが Println 直書きとの違いです。
	msg := fmt.Sprintf("%sは%d円です", item, price)
	fmt.Println(msg)
}
