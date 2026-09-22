# Lesson 23 パッケージとモジュール

Go には `if __name__ == "__main__"` がなく、`package main` の `func main` が唯一の入口です。import の書き方、`go mod init` と複数ファイル分割、main より先に走る `init()` を学びます。

## ch01 package と main が入口

Go のソースファイルは必ず `package` 宣言（このファイルが属するパッケージ名の表明）で始まります。実行プログラムは `package main` に属し、その中の `func main` が**唯一の入口**です。Python では import されたときに実行されないよう `if __name__ == "__main__":` というガードを書きましたが、Go にこのガードはありません。実行は常に main から始まり、それ以外の関数は呼ばれるまで動かないからです。🟡 「入口は main だけ」という決まりを理解しましょう。

Python ではこう書いた:

```python
def main():
    print("hello")

if __name__ == "__main__":
    main()  # import 時に走らないためのガード
```

Go ではこう書く:

```go
package main

func main() { // ここが唯一の入口。ガードは不要
	fmt.Println("hello")
}
```

## ch02 import と標準ライブラリ（math・os など）

**標準ライブラリ**（Go に最初から付属するパッケージ群）はインストール不要で、`import` するだけで使えます。複数のパッケージは括弧でまとめて書きます。Python の `from math import sqrt` のように関数だけを取り出す書き方はなく、常に `math.Sqrt` のように**パッケージ名.関数名**で呼びます。`os`（OS まわりの機能）、`time`（時刻）、`math/rand`（乱数）なども同じ形で使います。🟡 まとめ書きの形に慣れましょう。

Python ではこう書いた:

```python
import math
from math import sqrt  # 関数だけ取り出せる
print(sqrt(16))
```

Go ではこう書く:

```go
import (
	"fmt"
	"math"
)

fmt.Println(math.Sqrt(16)) // 常に パッケージ名.関数名
```

## ch03 go mod init と複数ファイル

**モジュール**は依存パッケージとバージョンを管理する単位で、`go mod init 名前` を実行すると管理ファイル `go.mod` が作られます。同じディレクトリで同じ `package main` を宣言したファイル同士なら、別ファイルの関数を **import なしで**そのまま呼べます。Python では別ファイルに分けたら `import` が必要でしたが、Go はパッケージ単位で 1 つの空間を共有します。実行は `go run .` とディレクトリごと指定します。🟡 まず 1 ファイル内の関数分割で感覚をつかみましょう。

Python ではこう書いた:

```python
# helper.py に分けたら import が必要
from helper import add
print(add(1, 2))
```

Go ではこう書く:

```go
// helper.go（同じ package main）に書いた add は
// main.go から import なしで呼べる
fmt.Println(add(1, 2))
```

## ch04 init()（main より先に走る初期化関数）

`init()` は、パッケージが読み込まれたときに **main より先に自動で実行される**初期化関数です。自分で呼び出すコードは書きません。Python でモジュールのトップレベルに書いたコードが import 時に実行されるのと同じ役割です。便利な半面、多用すると処理の流れが追いにくくなるため、実務では設定の初期化など限られた場面で使われます。⚪ 「main より先に走る関数がある」ことだけ知っておきましょう。

Python ではこう書いた:

```python
# module.py のトップレベルコードは import 時に走る
print("読み込みました")
```

Go ではこう書く:

```go
func init() {
	fmt.Println("main より先に実行される")
}
```
