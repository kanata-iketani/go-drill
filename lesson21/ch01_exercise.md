# lesson21 ch01 演習: type X struct（class の代わり）

1 行目に名前と年齢（整数）が空白区切りで与えられます。
`Name string` と `Age int` を持つ**構造体** `Person` を定義し、読み取った値で初期化して、
`<Name>は<Age>歳です` の形式で出力してください（フィールド経由で出力すること）。

**入力例1**

```text
Sato 28
```

**期待出力1**

```text
Satoは28歳です
```

## テストケース

### ケース 1

入力:

```text
Sato 28
```

期待出力:

```text
Satoは28歳です
```

### ケース 2

入力:

```text
Tanaka 40
```

期待出力:

```text
Tanakaは40歳です
```

### ケース 3

入力:

```text
Ai 7
```

期待出力:

```text
Aiは7歳です
```
