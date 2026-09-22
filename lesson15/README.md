# Lesson 15 標準入力

Python の `input()` 相当が Go では `bufio.Scanner` を使った 4 行になります。スキルチェックはこの定型が書ければ戦えるため、理屈と合わせて丸暗記してください。

## スキルチェック用テンプレート

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024) // 長い行に備えて上限を拡大

	// 1 行目: 整数 N
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	// 2 行目: 空白区切りの N 個の整数
	sc.Scan()
	f := strings.Fields(sc.Text())
	nums := make([]int, n)
	for i := range n {
		nums[i], _ = strconv.Atoi(f[i])
	}

	fmt.Println(nums)
}
```

## ch01 bufio.Scanner で 1 行読む

**標準入力**（キーボードやテストケースからプログラムに渡されるデータ）を読むには bufio.Scanner を使います。Python の input() 相当が、Go では「Scanner を作る → Scan() で 1 行読む → Text() で取り出す」と 4 行かかります。理屈より先に、この定型を丸暗記することがスキルチェック攻略の近道です。Text() が返す文字列に改行は含まれません。🔴 手が覚えるまで何度も書きましょう。

Python ではこう書いた:

```python
s = input()
print(s)
```

Go ではこう書く:

```go
sc := bufio.NewScanner(os.Stdin)
sc.Scan()
s := sc.Text()
fmt.Println(s)
```

## ch02 読んだ文字列を数値へ（strconv.Atoi）

Scanner が読み取った行は必ず文字列なので、数値として計算するには strconv.Atoi で変換します。Atoi は (int, error) の 2 つを返しますが、スキルチェックでは入力形式が保証されているため、n, _ := strconv.Atoi(...) と書いて err を _ で捨てるのが定石です。Python なら int(input()) の 1 行で済む処理が、Go では 4 行の定型 + 変換 1 行になります。これも考えずに書けるまで丸暗記します。🔴 手が覚えるまで書きましょう。

Python ではこう書いた:

```python
n = int(input())
```

Go ではこう書く:

```go
sc := bufio.NewScanner(os.Stdin)
sc.Scan()
n, _ := strconv.Atoi(sc.Text())
```

## ch03 1 行に複数の値（strings.Fields で分割）

1 行に複数の値が空白区切りで並ぶ入力は、strings.Fields で分割します。Fields は文字列を空白で区切って []string を返す関数で、Python の split() に相当します。数値が必要な要素は、取り出してから strconv.Atoi で 1 つずつ変換します。Python の map(int, input().split()) のような一括変換は Go にないため、ここでも行数は増えます。🔴 「Fields で割って Atoi」の流れを手が覚えるまで書きましょう。

Python ではこう書いた:

```python
a, b = map(int, input().split())
```

Go ではこう書く:

```go
sc.Scan()
f := strings.Fields(sc.Text())
a, _ := strconv.Atoi(f[0])
b, _ := strconv.Atoi(f[1])
```

## ch04 複数行の入力（N を読んでから N 行読む）

「1 行目に件数 N、続けて N 行のデータ」という形式はスキルチェックの最頻出パターンです。まず 1 行目を読んで Atoi で N に変換し、for の中で sc.Scan() と sc.Text() を繰り返して N 行を読みます。Scan() は呼ぶたびに次の行へ進むので、Python で input() をループの中で呼んだのと同じ感覚です。input() 相当が 4 行かかる Go でも、この骨組みさえ覚えれば複数行入力は怖くありません。🔴 試験本番で迷わないよう、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
n = int(input())
for i in range(n):
    line = input()
```

Go ではこう書く:

```go
sc.Scan()
n, _ := strconv.Atoi(sc.Text())
for range n {
	sc.Scan()
	line := sc.Text()
}
```

## ch05 スキルチェック定型テンプレート総まとめ

このレッスンの総まとめとして、スキルチェック用の定型テンプレートを丸ごと覚えます。bufio.Scanner を作り、Buffer で読める行の上限を広げ、strings.Fields で行を分割し、strconv.Atoi で数値化する。Python なら input().split() の 1 行で済む処理に、Go では input() 相当だけで 4 行かかります。だからこそ、この一式を考えずに白紙から書き出せるよう丸暗記しておくことが重要です。🔴 1 分で書けるようになるまで繰り返しましょう。

Python ではこう書いた:

```python
n = int(input())
nums = list(map(int, input().split()))
```

Go ではこう書く:

```go
sc := bufio.NewScanner(os.Stdin)
sc.Buffer(make([]byte, 1024*1024), 1024*1024)
sc.Scan()
n, _ := strconv.Atoi(sc.Text())
sc.Scan()
f := strings.Fields(sc.Text())
```
