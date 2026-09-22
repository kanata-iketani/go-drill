// lesson17 ch01: Go に set 型はない（map で代用）
package main

import "fmt"

func main() {
	// Python の {"apple", "banana", "orange"} に相当する集合です
	set := map[string]bool{
		"apple":  true,
		"banana": true,
		"orange": true,
	}

	// 所属判定: ないキーはゼロ値 false になるので、そのまま使えます
	fmt.Println(set["apple"])
	fmt.Println(set["grape"])

	// メモリ重視なら値を空構造体にする map[string]struct{} も使われます
	set2 := map[string]struct{}{"apple": {}}
	_, ok := set2["apple"]
	fmt.Println(ok)
}
