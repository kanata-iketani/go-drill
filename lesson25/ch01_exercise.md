# lesson25 ch01 演習: 値型のコピー（int・struct・配列は代入で複製）

1 行目に名前と年齢が空白区切りで入力されます。構造体 `Person{Name, Age}` の変数 `original` を作り、`copied := original` と代入してから `copied.Age` に 10 を足してください。そのあと 1 行目に `original: <名前> <年齢>`、2 行目に `copied: <名前> <年齢+10>` を出力し、構造体が値型なのでコピー側だけが変わることを確かめます。

**入力される値**

```text
名前 年齢
```

**入力例1**

```text
田中 20
```

**期待出力1**

```text
original: 田中 20
copied: 田中 30
```

## テストケース

### ケース 1

入力:

```text
田中 20
```

期待出力:

```text
original: 田中 20
copied: 田中 30
```

### ケース 2

入力:

```text
鈴木 35
```

期待出力:

```text
original: 鈴木 35
copied: 鈴木 45
```

### ケース 3

入力:

```text
Gopher 5
```

期待出力:

```text
original: Gopher 5
copied: Gopher 15
```
