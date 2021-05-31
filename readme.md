# 環境構築
1. コンテナ起動  
```
$ cd /path/to/money-app-ops
$ docker-compose up -d
```
2.コンテナの中に入り、ginのビルドインサーバー起動  

```
$ docker exec -it money-app-ops_app_1 bash
# pwd
/go/src/money-app
# go run main.go
```

(ここで以下のようなdb接続エラーがでた場合)

```
root@c92990514630:/go/src/money-app# go run main.go
2021/03/24 08:42:26 dial tcp: lookup mysql on 127.0.0.11:53: no such host
panic: dial tcp: lookup mysql on 127.0.0.11:53: no such host

goroutine 1 [running]:
money-app/db.Init()
	/go/src/money-app/db/db.go:32 +0x285
main.main()
	/go/src/money-app/main.go:10 +0x25
exit status 2
```

コンテナからいったん出て、dockerを再度起動し直してください。

```
$ docker-compose restart
```


3.(必要であれば)seedデータ作成APIを叩く  
※ 5秒ほどで1000件のデータがmoney_appと言うDBに出来上がります。

```
$ curl --location --request POST 'localhost:8080/sampleData'
```

# エンドポイント
1. ユーザーに残高を追加  
POST:localhost:8080/users/{userId}/balance
   
```
$ curl --location --request POST 'localhost:8080/users/{userId}/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount": 1000,
    "idempotentKey": "冪等生を担保するための36文字のUUID(このkeyが同じリクエストは冪等生が担保されます)"
}'
```

2.全ての顧客に残高を追加  
POST:localhost:8080/balances
```
$ curl --location --request POST 'localhost:8080/balances' \
--header 'Content-Type: application/json' \
--data-raw '{
    "amount": 1000,
    "idempotentKey": "冪等生を担保するための36文字のUUID(このkeyが同じリクエストは冪等生が担保されます)"
}'
```

# 後処理

### コンテナストップ
```
$ docker-compose stop
```

### コンテナやイメージなどすべて削除
```
$ docker-compose down --rmi all --volumes --remove-orphans
```