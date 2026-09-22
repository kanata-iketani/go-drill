// lesson23 ch04: init() は main より先に走る
package main

import "fmt"

var config string

// init はパッケージ読み込み時に自動で呼ばれます（自分では呼ばない）
func init() {
	config = "日本語モード"
	fmt.Println("init: 設定を初期化しました")
}

func main() {
	// この時点で config は init によって設定済みです
	fmt.Println("main: 設定 =", config)
}
