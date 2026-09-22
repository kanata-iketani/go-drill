// lesson16 ch04: delete と len
package main

import "fmt"

func main() {
	stock := map[string]int{"pen": 120, "note": 180, "clip": 50}
	fmt.Println(len(stock)) // 3

	// delete でキーごと削除します
	delete(stock, "note")
	fmt.Println(len(stock)) // 2

	// 存在しないキーを delete してもエラーになりません（何も起きない）
	delete(stock, "eraser")
	fmt.Println(len(stock)) // 2
}
