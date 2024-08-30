# GoをDockerで動かす手順

## step1 Dockerfile作成
公式のgolangのDockerイメージをプルする。
```Dockerfile
FROM golang:1.23

RUN mkdir myapp

WORKDIR /myapp

# ソースコードをコピー
COPY . .

# 必要に応じて go.mod を整える
RUN go mod tidy

CMD ["/bin/bash"]

```

## step２　go.modを手書きで作成
go.modを作り、以下をコピー
```
module 任意のモジュール名

go 1.23.0
```

## step3　ビルド&コンテナ実行
イメージのビルド
```
docker build -t 任意のイメージ名 .
```

コンテナ実行&コンテナに入る
```bash
docker run -it --rm -v $(pwd):/myapp my-golang-app
```
--rmをつけておけば、コンテナは終了と同時に削除されるので便利。
-v $(pwd):myappで、現在のホストディレクトリをコンテナ内のmyappにマウントする。


## step4　コンテナの中で、実行可能ファイル作成&実行

```bash
# アプリケーションをビルド
RUN go build -v -o /myapp/main ./main.go

# 実行
./main

# ビルド&実行
go run main.go
```