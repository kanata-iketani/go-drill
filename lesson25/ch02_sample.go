// lesson25 ch02: スライスとマップは中身を共有する
package main

import "fmt"

func main() {
	// スライスの代入はヘッダ（配列の場所・長さ・容量）のコピーだけです
	a := []int{1, 2, 3}
	b := a
	b[0] = 100
	fmt.Println(a) // [100 2 3] 同じ配列を指しているため a も変わる
	fmt.Println(b) // [100 2 3]

	// マップも内部データへの参照を持ちます
	m1 := map[string]int{"apple": 150}
	m2 := m1
	m2["apple"] = 200
	fmt.Println(m1["apple"]) // 200
}
