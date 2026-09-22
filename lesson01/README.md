# Lesson 01 Goとは



## ch01 Goはコンパイルして動かす

Go は**コンパイル**（ソースコードを機械語に変換すること）してから動かす言語です。Python はスクリプトをインタプリタがそのまま実行しましたが、Go はスクリプトではなくビルドして動かします。とはいえ `go run` コマンドがコンパイルと実行をまとめてやってくれるので、学習中の使い心地は Python とほぼ同じです。🟡 まずは「書く → go run で動かす」の流れを理解しましょう。

Python ではこう書いた:

```sh
python3 main.py
```

Go ではこう書く:

```sh
go run main.go
```

## ch02 package main と func main

Python はファイルの先頭から順に実行されました。Go では、プログラムの入口が `func main`（main 関数）と決まっていて、さらにその関数は `package main` というパッケージ（コードの所属を表す宣言）に置く必要があります。この 3 点セット（package main・import・func main）はすべての Go プログラムの骨格です。🔴 何も見ずに書けるまで手を動かしましょう。

Python ではこう書いた:

```python
print("Hello")
```

Go ではこう書く:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}
```

## ch03 静的型付けと gofmt

Go は**静的型付け**（変数の型がコンパイル時に決まっていて、あとから変えられない仕組み）の言語です。Python では変数に何でも入れられましたが、Go で int の変数に文字列を入れるとコンパイルエラーになります。エラーが実行前に見つかるのが利点です。また Go には公式の整形ツール `gofmt` があり、インデントや空白はすべて機械が揃えます。🟡 「型は固定」「整形は gofmt に任せる」の2点を押さえましょう。

Python ではこう書いた:

```python
x = 10
x = "hello"  # 何でも入る
```

Go ではこう書く:

```go
var x int = 10
// x = "hello" // コンパイルエラー: int に文字列は入らない
```
