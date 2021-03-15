# GoのORM、entを使ってGremlinServerに繋げる
これは、GoのORM、entを使ってGremlinServerに繋げるチュートリアルです。

## go generate
ディレクトリを移動して、entのコードを自動生成します
```
cd content/blog/article-3/
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
gremlin> g.V().valueMap(true)
==>{id=3c48f9e7-2ad0-403c-9398-96101b6d2cbf, label=user, name=[bob]}
==>{id=cd3900cf-79fe-4bde-8a55-8c840ce9d8e3, label=user, name=[alice]}
gremlin> g.E()
==>e[e7ca2294-a11d-4ffc-8a98-22ef1196df47][cd3900cf-79fe-4bde-8a55-8c840ce9d8e3-user_friends->3c48f9e7-2ad0-403c-9398-96101b6d2cbf]
```
