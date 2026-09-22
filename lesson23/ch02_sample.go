// lesson23 ch02: import と標準ライブラリ
package main

// 複数のパッケージは括弧でまとめて import します
import (
	"fmt"
	"math"
	"os"
)

func main() {
	// 常に「パッケージ名.関数名」で呼びます（from math import sqrt はない）
	fmt.Println(math.Sqrt(16))     // 4
	fmt.Println(math.Pow(2, 8))    // 256
	fmt.Println(math.Abs(-3.5))    // 3.5
	fmt.Println(math.Floor(9.876)) // 9

	// os パッケージも同じ形で使います（Args はコマンドライン引数）
	fmt.Println(len(os.Args) > 0) // true
}
