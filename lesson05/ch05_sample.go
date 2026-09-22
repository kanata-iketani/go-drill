// lesson05 ch05: fmt.Sprintf（f文字列の代わり）
package main

import "fmt"

func main() {
	name := "佐藤"
	age := 28
	height := 169.5

	// Sprintf は文字列を「作って返す」だけで、画面には出しません
	msg := fmt.Sprintf("%sさんは%d歳です", name, age)
	fmt.Println(msg)

	// 書式は Printf と共通です
	profile := fmt.Sprintf("身長 %.1f cm", height)
	fmt.Println(profile)

	// 作った文字列は普通の文字列として連結にも使えます
	fmt.Println(msg + "（" + profile + "）")
}
