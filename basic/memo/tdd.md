# これは何
tdd に関する学びをここに書く

# commit するタイミング
- 最初にテストをパスした時点でいったん commit しておくとよい。リファクタの過程で動作する状態に戻したくなることもあるため。

# TDD の手順
- テストを書く
- コンパイラをパスする
- テストを実行し、失敗することを確認し、エラーメッセージが意味があることを確認する
- テストをパスするように実装を追加・修正する
- リファクタリング