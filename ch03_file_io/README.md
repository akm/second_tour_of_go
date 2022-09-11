# ファイル操作

## :exclamation: セットアップ

1. ターミナルあるいはコマンドプロンプトで `ch0_file_io` に移動
2. `go mod init github.com/akm/second_tour_of_go/ch03_file_io` を実行
   - `github.com/akm` の部分は他の文字列に変更しても OK です。
     - Github 等で管理するのであれば、自分の環境に合わせて変更してください

## :question: cat サブコマンド

引数で指定されたファイルの内容を標準出力に出力する `cat` サブコマンドを作成してください。
ただし実行形式のファイル名は `ch03_file_io` とします。
存在しないファイルを読み込もうとした場合などは原因となるエラーを標準エラー出力に出力し、終了コードを 1 としてください。

### 結果

```
./ch03_file_io                            # => (ヘルプを表示)
./ch03_file_io cat                        # => (ヘルプを表示)
./ch03_file_io cat README.md              # => (README.md のファイルの中身を表示)
./ch03_file_io cat README.md > /dev/null  # => (何も出力されない)
./ch03_file_io cat not_found              # => open ./READMEx.md: no such file or directory
./ch03_file_io cat not_found 2> /dev/null # => (何も出力されない)
./ch03_file_io unknown                    # => (ヘルプを表示)
```

PowerShell の場合 `/dev/null` を `$null` に置き換えてください。

### ヒント

標準エラーにメッセージを出力するには [os.Stderr](https://pkg.go.dev/os#pkg-variables) と [fmt.Fprint](https://pkg.go.dev/fmt#Fprint)、[fmt.Fprintf](https://pkg.go.dev/fmt#Fprintf)あるいは[fmt.Fprintln](https://pkg.go.dev/fmt#Fprintln) を使ってください。

### 読み込み方法 1

- [os.ReadFile](https://pkg.go.dev/os#ReadFile)

### 読み込み方法 2

- [os.Open](https://pkg.go.dev/os#Open)
- [io.ReadAll](https://pkg.go.dev/io#ReadAll)
- [\*File.close](https://pkg.go.dev/os#File.Close)

## :question: write サブコマンド

引数で指定されたファイルに、引数で指定された内容と末尾の改行を書き込む `write` サブコマンドを作成してください。
ただし実行形式のファイル名は `ch03_file_io` とします。
存在するファイルに書き込む場合は上書きするものとします。
書き込み等のエラーが発生した場合は標準エラー出力に出力し、終了コードを 1 としてください。

### 結果

```
./ch03_file_io                            # => (ヘルプを表示)
./ch03_file_io cat                        # => (ヘルプを表示)
./ch03_file_io cat README.md              # => (README.md のファイルの中身を表示)
./ch03_file_io cat README.md > /dev/null  # => (何も出力されない)
./ch03_file_io cat not_found              # => open ./READMEx.md: no such file or directory
./ch03_file_io cat not_found 2> /dev/null # => (何も出力されない)
./ch03_file_io write                      # => (ヘルプを表示)
./ch03_file_io write foo.txt              # => (ヘルプを表示)
./ch03_file_io write foo.txt Foo          # => (何も出力しない、終了コード0)
./ch03_file_io unknown                    # => (ヘルプを表示)
```

### ヒント

- [os.Create](https://pkg.go.dev/os#File)
- [fmt.Fprintln](https://pkg.go.dev/fmt#Fprintln)

## :question: write サブコマンドで上書き確認

write サブコマンドを変更し、存在するファイルが指定された場合は以下のようなメッセージを出力し、ユーザーに `y` あるいは `n` を入力してもらうことによって、上書きの是非を確認してください。

### ヒント

- [os.Stat](https://pkg.go.dev/os#Stat)
- [os.IsExist](https://pkg.go.dev/os#IsExist)
- [fmt.Scan](https://pkg.go.dev/fmt#Scan)

### テスト方法

```
$ rm foo.txt
$ go run . write foo.txt fooooooo
$ cat foo.txt
fooooooo
$ go run . write foo.txt baarrrrr
foo.txt already exists. Overwrite? (y/n): n
Quit writing
exit status 1
$ cat foo.txt
fooooooo
$ go run . write foo.txt baarrrrr
foo.txt already exists. Overwrite? (y/n): y
$ cat foo.txt
baarrrrr
```
