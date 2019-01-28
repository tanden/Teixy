# docker-composeのmysqlコンテナの設定について
## mysqlのバージョン
2019/01の最新安定板の8.0.13のDockerイメージを使用

## コンテナとDB
1つのDBに対して、1のコンテナを設定するのがよさそう。というのは、
rootユーザとは別に、`MYSQL_DATABASE`に書いた名前のDBに対して全ての権限を`MYSQL_USER`と`MYSQL_PASSWORD`で得ることができる。
ただ、複数のDBは設定できないので、1コンテナ1DBを想定しているよう。
（qiitaとかを漁れば、1つのコンテナに2つ以上のDBを設定している人もいるので必要になれば参照する）

## Volumeの設定
mysqlのデータは永続化したい（コンテナを消去しても残しておきたい）ので、volumesを使う。<br>
mysqlの永続化の方法（コンテナとホストのボリュームを共有する方法）は複数あるが、docker-composeのドキュメントがお勧めするvlumesを使う方法で設定する<br>
### docker-composeのドキュメント
[Use volumes](https://docs.docker.com/storage/volumes/)<br>
[Example](https://docs.docker.com/compose/compose-file/#volumes)
