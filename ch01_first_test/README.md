# 第1章 初めてのテスト

## セットアップ

1. https://go.dev/doc/install の指示に従って Go 1.19 をインストール
    [asdf](https://asdf-vm.com/)の[golangプラグイン](https://github.com/kennyp/asdf-golang) あるいは [goenv](https://github.com/syndbg/goenv) などを使うことを推奨
1. `go` コマンドのバージョンを確認
    1. `ターミナル` (Mac/Linux) or `コマンドプロンプト` (Windows) を新たに開く
    1. `go version` を実行
    1. インストールしたGoのバージョンが表示されることを確認
1. https://github.com/akm/second_tour_of_go/releases/tag/ch01_first_test_01 からソースコードをダウンロード
1. ダウンロードしたコードを解凍
1. ターミナル上でディレクトリ `ch01_first_test` に移動
    `cd path/to/second_tour_of_go-ch01_first_test_01/ch01_first_test`
    (`path/to` は解凍して作られたディレクトリに置き換えて実行してください )
1. `go test` を実行し、以下のような出力が表示されることを確認
    ```
    PASS
    ok      github.com/akm/second_tour_of_go/ch01_first_test        0.341s
    ```

## :question: 関数Addを作成

足し算を行う関数Addを、main_test.goと同じディレクトリのファイル calc.go に作成してください。

また、テストをAddを使うように変更して、テストがパスすることを確認してください。

:point_up: テストに失敗するようにパラメータや期待する結果を変更すると、どのように出力が変わるのかを確認してください。

## :question: 関数Substract (テストファースト)

引き算を行う関数 Substract を以下の手順で実装してください。

1. 関数Substractの仮実装を作成する（テスト成功）
2. テストTestSubstractを作成する（テスト失敗）
    - TestSubstractでは複数の計算とその値の確認を行うようにしてください
3. Substractを実装する（テスト成功）

このように、このように実装を関数の中身を実装する前にテストを作成する手法を **テストファースト** と呼びます。
今回の実装の中身は非常に簡単なので、テストファーストの恩恵がわかりにくいかもしれませんが、今後もできるだけ
テストファーストでテストと実装を作るようにお願いします。

## :exclamation: testingパッケージのドキュメントを読む

[testingパッケージのドキュメント](https://pkg.go.dev/testing)を開いて以下を確認してください。

- 関数名の規則
- Type T
  - 失敗を知らせるための関数
      - Error
      - Errorf
      - Fail
      - FailNow
      - Fatal
      - Fatalf
  - ログ出力
      - Log
      - Logf

現時点では上記以外はスルーしてOKです。時間があれば、上記を試してみてください。
