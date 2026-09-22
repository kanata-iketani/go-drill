# Lesson 06 文字列と数値の相互変換

`str()` `int()` `float()` に当たるのが `strconv` パッケージの関数です。最大の違いは `strconv.Atoi` が `(int, error)` の 2 値を返すことで、この講座ではここで初めて `err` が登場します。

## ch01 strconv.Itoa（数値→文字列）

数値を文字列に変換するには `strconv` パッケージ（文字列と数値の相互変換を集めた標準ライブラリ）の `strconv.Itoa` を使います。Python の `str(42)` に相当する関数です。Go は文字列と int を `+` で直接つなげられないため、連結する前に必ず文字列へ変換する必要があります。Itoa は Integer to ASCII の略です。🔴 数値を文章に埋め込む場面は頻出なので、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
s = str(42)
print("値は" + s)
```

Go ではこう書く:

```go
s := strconv.Itoa(42)
fmt.Println("値は" + s)
```

## ch02 strconv.Atoi と err（文字列→数値）

文字列を数値に変換する `strconv.Atoi` は、Python の `int("42")` に相当します。大きな違いは、Atoi が `(int, error)` の **2 つの値**を返すことです。`error` は「処理が失敗したかどうか」を表す型で、この講座ではここで初めて登場します。成功すると err には `nil`（何もないことを表す値）が入り、失敗すると理由が入ります。`if err != nil { ... }` で失敗を確認するのが Go の定型です。🔴 この形は今後ずっと使うので、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
n = int("42")
```

Go ではこう書く:

```go
n, err := strconv.Atoi("42")
if err != nil {
	fmt.Println("変換できません")
}
```

## ch03 strconv.ParseFloat（文字列→float64）

文字列を小数に変換するには `strconv.ParseFloat` を使います。Python の `float("3.5")` に相当します。第 2 引数には精度を渡し、`float64` を使う Go では常に `64` と書きます。Atoi と同じく戻り値は `(float64, error)` の 2 値で、`if err != nil` のチェックもまったく同じ形です。🟡 「文字列→数値の変換は必ず 2 値で返る」という仕組みを理屈で理解しましょう。

Python ではこう書いた:

```python
x = float("2.5")
```

Go ではこう書く:

```go
x, err := strconv.ParseFloat("2.5", 64)
```

## ch04 strconv.FormatFloat（float64→文字列）

float64 を文字列に変換するには `strconv.FormatFloat` を使います。Python では `str(3.14)` や f文字列の `:.2f` で済んだ変換です。引数は（値・表記の種類・小数桁数・精度）の 4 つで、ふだんは表記 `'f'` と精度 `64` を固定して桁数だけ変えれば十分です。同じことは `fmt.Sprintf("%.2f", x)` でもでき、実務ではそちらもよく使います。⚪ まずはこういう関数があると知っておくだけで十分です。

Python ではこう書いた:

```python
s = f"{3.14159:.2f}"
```

Go ではこう書く:

```go
s := strconv.FormatFloat(3.14159, 'f', 2, 64)
```
