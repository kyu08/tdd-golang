# これは何
golang に関する学びをここに書く

# test を作成する時のルール
- `xxx_test.go` のようなファイル名にする必要がある
- テスト関数は `Test` という単語で始まる必要がある
- テスト関数の1つの引数のみをとる `t *testing.T`

# 文字列を定数に格納するメリット
以下のようにすると Hello が呼び出されるたびに "Hello, " 文字列インスタンスが生成されるのを防ぐことができ、パフォーマンス向上につながる

```go
const helloPrefix = "Hello, "

func Hello(name string) string {
	return helloPrefix + name
}
```

# t.Helper()
を実行すると
- テストやベンチマークからヘルパー関数を呼び出すことができる
- そのメソッドがヘルパーであることをテストスイートに伝えるために必要
  - テストが失敗された時に表示される行番号はヘルパーではなく呼び出された関数の中を示すため、追跡がしやすくなる。

# named retuirn value を用いるメリット
関数のGo Docに表示されるので、コードの意図をより明確にすることができる
ただ、コード自体の可読性が落ちるように感じるのであまり使いたいとは思わなかった

# slice の比較
slice で等号演算子を使うことはできないのでここでは`reflect.DeepEqual`を使う
ただし、 `reflect.DeepEqual` は型安全ではないため注意して使う必要がある

# errcheck による網羅性のチェック
`go install github.com/kisielk/errcheck@latest` して `errcheck ./...` を実行するとエラーをチェックしていない箇所を教えてくれる

# map の注意点
空のマップ変数を以下のように初期化してしまうとゼロ値である `nil` になってしまう。 nilマップに書き込もうとするとランタイムパニックが発生してしまう。
```go
var m map[string]string
```

そのため、以下のようにしてマップを作成しておくべきである。
```go
var m = map[string]string{}

// or
var m = make(map[string]string)
```
