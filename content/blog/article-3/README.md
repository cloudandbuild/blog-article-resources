# GoのORM、entを使ってGremlinServerに繋げる
これは、GoのORM、entを使ってGremlinServerに繋げるチュートリアルです。

## go generate
ディレクトリを移動して、entのコードを自動生成します
```
go mod download
go generate ./...
```
## GremlinServerを起動する
```
CLOUDSDK_PYTHON=/usr/bin/python
docker-compose build
docker-compose up -d
```

## 実行
Goのプログラムを実行します。
```
go run main.go
```
## GremlinConsoleでQueryして結果を確認する
gremlin-consoleを開く
```
docker-compose exec  gremlin-console /bin/bash -c bin/gremlin.sh
```

gremlin-serverに接続する。
```
:remote connect tinkerpop.server /opt/gremlin/conf/remote.yaml
:remote console
```

queryを実行する。
```
g.V().valueMap(true)
```
