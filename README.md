# allysum

## コード化できていない環境設定
### docker
#### docker-compose
AWS Cloud9にはdockerは初めから入っているが、docker-composeはインストールされていない  
インストール方法はドキュメントにあるので参照する  
[Install docker-compose](https://docs.docker.com/compose/install/)  

### direnv
teixy以下で特定の環境変数を有効にするために、direnvを使っている  
設定は各ディレクトリで`.envrc`にそのディレクトリだけで有効な環境変数を記載する（on/offできる）  
インストールはgithubのREADMEを参照する  
[direnv/direnv](https://github.com/direnv/direnv)  

### goenv
任意のディレクトリのgoのバージョンを指定できる  
.go-versionにそのディレクトリで使用するgoのバージョンを記載  
goenvは色々あるが下記のgoenvを使用する  
インストールの方法などはREADMEを参照  
[syndbg/goenv](https://github.com/syndbg/goenv)  

### circleci cli tool
circleciのローカルでテストしたりできるcli tool  
インストールと基本的な使い方はドキュメントを参照  
[Using the CircleCI Local CLI](https://circleci.com/docs/2.0/local-cli/)  
  
しかし、AWS Cloud9にインストールする場合、こける可能性が高い  
対処法はここにまとめておいた  
[CircleCI Local CLIのインストールに失敗する](https://qiita.com/tanden/items/fac327992314ec0fa24a)  

## AWS Cloud9の環境設定

### timezone
デフォルトではUTCなので変更する　　
```
% cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime
# dateコマンドでJSTになっていればok
% date
Sat Feb  9 21:27:53 JST 2019
```

db/create_sql.shから作成するsqlファイルのタイムスタンプが前後する可能性があるので、必ずJSTにセットしてから実行する

## 入れると便利な開発ツール
### tig
```
$ yum install tig
```
で入るはず  
gitのコミット履歴を見やすくしてくれる

## 設定すると便利なalias
```
alias dc=docker-compose
alias ci=circleci
```
