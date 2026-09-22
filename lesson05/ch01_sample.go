// lesson05 ch01: 文字列リテラルと + 結合
package main

import "fmt"

func main() {
	hello := "Hello"
	world := "World"

	// + で連結できます（Python と同じ）
	fmt.Println(hello + ", " + world + "!")

	// \n で途中に改行を入れられます
	fmt.Println("1行目\n2行目")

	// バッククォートは書いたまま出力されます（\n も文字のまま）
	fmt.Println(`複数行を
そのまま書けます`)
}
