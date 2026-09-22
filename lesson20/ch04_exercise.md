# lesson20 ch04 演習: キーワード引数・デフォルト引数はない

1 行目に名前、2 行目にあいさつの言葉（`-` のときは指定なしとみなす）が与えられます。
`Prefix string` を持つ**オプション構造体** `Options` と、関数 `greet(name string, opt Options)` を定義してください。
greet は `Prefix` が空文字列ならデフォルトの `Hello` を補い、`<Prefix>, <name>!` の形式で出力します。
2 行目が `-` なら `Options{}`、それ以外なら `Options{Prefix: 読んだ値}` を渡してください。

**入力例1**

```text
Taro
-
```

**期待出力1**

```text
Hello, Taro!
```

## テストケース

### ケース 1

入力:

```text
Taro
-
```

期待出力:

```text
Hello, Taro!
```

### ケース 2

入力:

```text
Hana
Hi
```

期待出力:

```text
Hi, Hana!
```

### ケース 3

入力:

```text
Ken
Yo
```

期待出力:

```text
Yo, Ken!
```
