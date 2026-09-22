Go に例外は存在せず、`try/except` の代わりに関数が `error` を値として返し、`if err != nil` で分岐します。`raise` の代替は `return err` です。
