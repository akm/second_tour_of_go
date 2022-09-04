# Second Tour of Go

## Goのインストール

### バージョンマネージャーの利用

バージョンマネージャーを利用するとインストールの管理や複数バージョンを利用を簡単に行うことができます。

バージョンマネージャー | 環境 | 複数のGoのバージョンの管理 | 他の他の言語・ツールの管理 | 備考
-------------------|------|-----------------------|-------------------------|-----
[asdf](https://asdf-vm.com/guide/getting-started.html) | macOS/Linux | OK | OK | オススメ
[goenv](https://github.com/syndbg/goenv) | macOS/Linux | OK | NG | anyenvから利用するのが一般的？
[g](https://github.com/voidint/g) | Linux/macOS/Windows | OK | NG | [Multi version management of golang in Windows](https://developpaper.com/multi-version-management-of-golang-in-windows/)
[scoop](https://scoop.sh/) | Windows | NG? | OK |


もしバージョンマネージャーを使うことが難しい場合は後述の [素でインストール](#素でインストール) を行ってインストールしてください。

### 素でインストール

https://go.dev/doc/install の指示に従って Go 1.19 をインストールします。
その環境に一つのバージョンしかいれることができないので、実際の開発時には困ることが多いインストール方法です。


## 取り組み方

### PRベース

各章のPRのコミットを追う方法です。
各コミットのコミットメッセージには実行時の出力などを記しているものもありますので、初心者にはこちらをおすすめします。
### 問題と答えをまとめて確認

各章について以下のように進める方法です。

1. READMEを読む
2. 各問題に回答する
3. 答え合わせする

## 目次

章           | PR | 問題と答え
-------------|------|---------
初めてのテスト | [PR](https://github.com/akm/second_tour_of_go/tree/ch01_first_test_01/ch01_first_test) | [ch01_first_test](https://github.com/akm/second_tour_of_go/tree/main/ch01_first_test)
