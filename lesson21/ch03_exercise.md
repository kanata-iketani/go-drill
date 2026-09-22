# lesson21 ch03 演習: コンストラクタ関数 NewX

1 行目にプレイヤー名とダメージ量（整数）が空白区切りで与えられます。
`Name string` と `HP int` を持つ構造体 `Player` と、
HP を 100 に初期化して `*Player` を返す**コンストラクタ関数** `NewPlayer(name string) *Player` を定義してください。
NewPlayer で作ったプレイヤーの HP からダメージを引き、`<Name> HP:<HP>` の形式で出力します。

**入力例1**

```text
Hero 30
```

**期待出力1**

```text
Hero HP:70
```

## テストケース

### ケース 1

入力:

```text
Hero 30
```

期待出力:

```text
Hero HP:70
```

### ケース 2

入力:

```text
Slime 0
```

期待出力:

```text
Slime HP:100
```

### ケース 3

入力:

```text
Mage 100
```

期待出力:

```text
Mage HP:0
```
