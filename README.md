# jwt-sample

authenticate server using jwt-go

## ディレクトリ構成

```
.
└──jwt-sample
     ├─cmd
     ├─codes
     ├─config
     ├─domain
     │  ├─model
     │  ├─repository
     │  └─service
     ├─infra
     ├─interface
     │  ├─handler
     │  └─persistence
     ├─proto
     ├─usecase
     └─main.go
```

### cmd
cli から実行するためのコードをここに書きます  
現在はトークン取得まで一度に実行するプログラムを書いています

### codes

ステータスコードをここで定義します  
Golang のエラーの扱い方がまだわかっていないので、とりあえず errors.New を使ってエラーコードを定義しています

### config

環境変数から値の取得をここで行います

### domain

アプリケーションのロジックをここで定義します  
model, repository, service のサブディレクトリを作っています  
repository はデータの永続化層の IF を定義しています  
service では repository を使ったロジックや複数のモデルを使うロジックを書きます


### infra

### interface

外部のシステムや入出力関係をここに書きます  
handler は Web API のリクエスト/レスポンスを, persistence は永続化 (データベース) 関連のコードです

### proto
gRPC のコードをここに書きます

### usecase
アプリケーションで出来ることを書いていきます  
domain に依存しています