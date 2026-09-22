# Lesson 08 スライスと文字列

`join` `split` に当たるのは `strings.Join` / `strings.Split` / `strings.Fields` で、使い方は Python とほぼ同じです。

## ch01 strings.Join（スライス→文字列）

スライスを 1 本の文字列にまとめるには `strings.Join(s, 区切り)` を使います。Python の `",".join(a)` とほぼ同じ機能です。違いは書き方の向きだけで、Python は区切り文字側のメソッドでしたが、Go はスライスを第 1 引数に取る関数です。区切りに `""` を渡せばそのまま連結になります。🔴 出力の整形で毎回使うので、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
print(",".join(a))
```

Go ではこう書く:

```go
fmt.Println(strings.Join(s, ","))
```

## ch02 strings.Split（文字列→スライス）

文字列を区切り文字で分割してスライスにするには `strings.Split(s, 区切り)` を使います。Python の `"a,b,c".split(",")` とほぼ同じで、結果は `[]string` になります。Join と Split は互いに逆の操作で、この 2 つでカンマ区切りデータの読み書きができます。分割した結果はスライスなので `s[i]` や `len` がそのまま使えます。🔴 データ処理の基本なので、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
colors = "赤,青,黄".split(",")
```

Go ではこう書く:

```go
colors := strings.Split("赤,青,黄", ",")
```

## ch03 strings.Fields（空白区切りの分割）

空白区切りの文字列を分割するには `strings.Fields(s)` を使います。Python で引数なしの `split()` を呼んだときとほぼ同じで、連続する空白やタブをまとめて 1 つの区切りとして扱います。`strings.Split(s, " ")` だと連続空白のあいだに空文字列の要素が入ってしまうため、空白区切りには Fields を選ぶのが定石です。🔴 のちの標準入力の読み取りで必ず使うので、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
words = "go  is   fun".split()
```

Go ではこう書く:

```go
words := strings.Fields("go  is   fun")
```
