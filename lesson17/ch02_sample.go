// lesson17 ch02: 存在判定と重複除去
package main

import "fmt"

func main() {
	words := []string{"go", "python", "go", "ruby", "python"}

	// seen（見た要素の記録）を set 代わりの map で持ちます
	seen := map[string]bool{}
	unique := []string{}
	for _, w := range words {
		if !seen[w] { // 初めて見た要素だけ通す
			seen[w] = true
			unique = append(unique, w)
		}
	}

	// Python の set(lst) と違い、元の順序が保たれます
	fmt.Println(unique) // [go python ruby]
}
