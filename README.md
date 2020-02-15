# golang_sample

### プラグイン
- go-plus
  保存時の自動コンパイル、テストはうざいので設定で切ってもいいかも

## 環境変数関連
```go env``` で見れる。特にGO_ROOT,GO_PATHに注意。

## 並列直列に関して
参考URL https://qiita.com/suin/items/82ecb6f63ff4104d4f5d
- parallel
    - chokuretsu 直列処理(比較のためで新ネタはなし)
    - heiretsu.go 並列処理　プレーンなもので無名関数使用
    - heiretsu2.go 並列処理　heiretsu1と同じだがチェネルの終了にfor文
    - heiretsu3.go 並列処理　heiretsu1と同じだが書き方を修正

    参考URL https://qiita.com/TakaakiFuruse/items/241578174fd2f00aaa8a
    - heiretsu4.go 並列処理　書き方を若干変更(無名関数を使わないためこちらのほうが一般的かも・・・)
    - heiretsu5.go 並列処理　同一処理を並行。これが一般的な使い方かと・・

    他参考URL https://tikasan.hatenablog.com/entry/2017/05/20/170050

## テストについて
別のライブラリを入れたりなどは必要なし

- sample
    - hello.go
    - hello_test.go
    同一階層かつ_testをつけることでtestスクリプトになる

テストコマンド
```
#対象のディレクトリに移動している必要がある
cd ./sample
#全テスト実行
go test  -run ''
#特定文字列を含むメソッドのみ
go test  -run Dummy
```

## インポート(import)に関して

- 標準のライブラリ
- 相対パス
- GO_PATHのものなど(主にライブラリ)

## パッケージ
http://cuto.unirita.co.jp/gostudy/post/go-package/

ルール
- 同一階層ではpackageを統一する必要がある(例外はテストコード)
- フォルダ名とあっていなくてもいい    
