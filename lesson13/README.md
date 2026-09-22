# Lesson 13 for と range

`range(n)` は Go 1.22 以降の `for i := range n` に、`enumerate` は最初から index と value を返す `for i, v := range s` に対応します。使わない値は `_` で捨てます。

## ch01 for i, v := range スライス（enumerate 相当）

スライスの全要素を順に処理するには `for i, v := range s {}` を使います。`range` はインデックスと値の 2 つを最初から返すため、Python で `enumerate` を使って書いていた形が、Go ではそのまま標準の書き方になります。値だけを取り出す `for v in s:` に対応する専用の形はなく、ch03 で学ぶ `_` でインデックスを捨てて表現します。🔴 スライス処理の基本形なので、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
for i, v in enumerate(fruits):
    print(i, v)
```

Go ではこう書く:

```go
for i, v := range fruits {
	fmt.Println(i, v)
}
```

## ch02 for i := range 10（range(n) の代わり）

Python の `for i in range(10):` に相当するのが `for i := range 10 {}` で、0 から 9 まで順に回ります。これは **Go 1.22 以降**で使える構文です。それより前の Go には `for i := 0; i < 10; i++ {}` という 3 部形式しかなく、古い記事や既存コードでは今もよく見ます。また `range(1, 11)` のような開始値の指定はできないため、`i+1` を使って調整します。🔴 回数指定ループの基本形として、手が覚えるまで書きましょう。

Python ではこう書いた:

```python
for i in range(3):
    print(i)
```

Go ではこう書く:

```go
for i := range 3 {
	fmt.Println(i)
}
```

## ch03 _ で値を捨てる

`range` の 1 つ目の戻り値はインデックスと決まっているため、値だけが欲しいときは 1 つ目を**ブランク識別子** `_`（受け取った値を捨てるための特別な名前）で受けて `for _, v := range s` と書きます。Go では宣言したのに使わない変数がコンパイルエラーになるので、「受け取るが使わない」を `_` で明示する必要があるのです。インデックスだけなら 2 つ目を省略して `for i := range s` と書けます。Python の `_` はただの慣習でしたが、Go では言語の仕組みです。🟡 なぜ `_` が必要なのか、理屈を理解しましょう。

Python ではこう書いた:

```python
for v in scores:
    print(v)
```

Go ではこう書く:

```go
for _, v := range scores {
	fmt.Println(v)
}
```
