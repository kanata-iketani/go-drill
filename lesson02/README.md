# Lesson 02 標準出力

Python の `print()` は `sep=` や `end=` で出力を細かく調整できましたが、Go にその引数はなく、`Print` と `Printf` を使い分けて表現します。

## ch01 fmt.Println

`fmt.Println` は Python の `print()` に一番近い関数です。複数の値をカンマで並べると空白区切りで出力され、最後に自動で改行が付きます。Python と違い、括弧の前に `fmt.` というパッケージ名が必要な点だけ注意してください。🔴 一番よく使う関数なので、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
print("score", 80)
```

Go ではこう書く:

```go
fmt.Println("score", 80)
```

## ch02 fmt.Print（end="" の代わり）

Python では `print("a", end="")` で改行を止めました。Go では改行を付けない `fmt.Print` を使います。`end=` という引数は存在しないので、「改行したい行は `Println`、したくない行は `Print`」と関数そのものを使い分けます。改行だけしたいときは `fmt.Println()` と空で呼びます。🟡 2 つの関数の違いを理屈で押さえましょう。

Python ではこう書いた:

```python
print("Hello", end="")
print("World")
```

Go ではこう書く:

```go
fmt.Print("Hello")
fmt.Println("World")
```

## ch03 fmt.Printf（sep= と書式の代わり）

`fmt.Printf` は**書式指定**（`%d` などの記号で値の埋め込み方を指定する仕組み）を使う出力関数です。`%d` は整数、`%s` は文字列、`%f` は小数（`%.1f` で小数第1位まで）に対応します。Python の `sep="-"` のような区切り指定も、Go では書式の中に直接 `-` を書いて表現します。`Printf` は改行を付けないので、末尾に `\n` を書きます。🔴 スキルチェックでも多用します。

Python ではこう書いた:

```python
print(2026, 9, 20, sep="-")
```

Go ではこう書く:

```go
fmt.Printf("%d-%d-%d\n", 2026, 9, 20)
```

## ch04 コメント

コメント（実行されないメモ書き）は、Python では `#` でしたが、Go では 1 行コメントが `//`、複数行コメントが `/* */` です。Go には Python の docstring はなく、関数の説明も `//` を関数の直前に書くのが慣習です。⚪ 記号が違うことだけ知っておけば十分です。

Python ではこう書いた:

```python
# 1行コメント
```

Go ではこう書く:

```go
// 1行コメント
/* 複数行
   コメント */
```
