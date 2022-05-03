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

ちなみに `channel`についても同様で、`var ch chan struct{}` とすると `ch` の値はゼロ値である `nil` になってしまうので `make` で初期化しておいた方がよい

# for文内でgoroutine を使う時の注意
以下のような書き方をすると変数が上書きされてしまうことに注意する
```go
for url, _ := range urls {
  go func() {
    someFunc(url)
  }()
}
```
以下のようにすると、各goroutine が独立した変数としての u を持つことができる
```go
for url, _ := range urls {
  go func(u string) {
    someFunc(u)
  }(url)
}
```

# リフレクション
TS でいう `any` は golang では `interface{}`に相当する

渡された引数の型を検査するにはリフレクションを使用する。
実行時にチェックを行う必要があるため、一般的にパフォーマンスが低下する。

## golang にも any があるみたいだけど interface{} とどう違うのか

## context とは
https://zenn.dev/hsaki/books/golang-context/viewer/definition

`Context`型の主な役割
- 処理の締切を伝達
- キャンセル信号の伝播
- リクエストスコープ値の伝達

`複数ゴールーチン間で安全に、そして簡単に情報伝達を行いたい`

### timeout
タイムアウト設定をしていた場合にも、明治的に`cancel`を呼ぶべき(contextリークを防ぐため)
https://pkg.go.dev/context#example-WithDeadline

### Value メソッド
値の伝達のために使う

### Value として与えてもいいデータ・与えるべきではないデータ
#### 与えるべきではないデータ
- 関数の引数
  - 見通しが悪くなるため
- type-unsafe になったら困るもの
- 可変な値
  - 見通しが悪くなるため
- ゴルーチンセーフではない値
  - 例: Value に slice の値を与えてそれに append する処理
  排他処理が取れていないため、ゴルーチンセーフであるといえない

#### Value として与えてもいいデータ
リクエストスコープな値

リクエストスコープとは、「1つのリクエストが処理されている間に共有される」という性質のこと。例をあげると
- ヘッダから抜き出したユーザID
- 認証トークン
- トレースID
などである。

## waitgroup とは
https://qiita.com/ruiu/items/dba58f7b03a9a2ffad65
`sync.WaitGroup` は複数の goroutine の完了を待つための値
以下のように goroutine の中で `wg.Add(1)`すると `wg.Add(1)`が呼ばれる前に `wg.Wait()` が呼ばれるとただしく動作しない。
```go
wg := &sync.WaitGroup{}  // WaitGroupの値を作る
for i := 0; i < 10; i++ { // （例として）10回繰り返す
    go func() {
        wg.Add(1)  // wgをインクリメント
        // ここで何か処理を行う
        wg.Done()  // 完了したのでwgをデクリメント
    }()
}
wg.Wait()  // メインのgoroutineはサブgoroutine 10個が完了するのを待つ
```
なので goroutine の外側で `wg.Add(1)`する
```go
wg := &sync.WaitGroup{}  // WaitGroupの値を作る
for i := 0; i < 10; i++ { // （例として）10回繰り返す
    wg.Add(1)  // wgをインクリメント
    go func() {
        // ここで何か処理を行う
        wg.Done()  // 完了したのでwgをデクリメント
    }()
}
wg.Wait()  // メインのgoroutineはサブgoroutine 10個が完了するのを待つ
```
