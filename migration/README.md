# migration

## migrateを使ってmigrationを管理する
[golang-migrate/migrate](https://github.com/golang-migrate/migrate)

### migrate.goの中で使っているドライバー
ローカルのsqlファイルの読み込み用のドライバー  
[migrate/source/file/](https://github.com/golang-migrate/migrate/tree/master/source/file)

mysql用のドライバー  
[migrate/database/mysql/](https://github.com/golang-migrate/migrate/tree/master/database/mysql)

### Goからpackageとしてimportして使う場合のDocument
[package migrate](https://godoc.org/github.com/golang-migrate/migrate)

## 主な開発の流れ
### create_sql.shを使って空のsqlファイル（up,down）を作成する
```
% ./create_sql.sh <sql_name>
```
### 空の.sqlファイルにクエリをかく
`<version>_<sql_name>.up.sql`には、テーブル作成クエリや、データのインサート、アップデートのクエリをかく  
`<version>_<sql_name>.down.sql`には、upで書いたクエリを打ち消すクエリをかく（create tableに対して、drop tableなど）

### goのバッチを実行
主なコマンド
### up
```
% go run migrate.go -exec up
```

### down
```
% go run migrate.go -exec down
```

### -f (force option)
```
% go run migrate.go -exec up -f
```

これで新しく追加されたクエリが実行されて、その内容がDBに反映される

