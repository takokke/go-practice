FROM golang:1.23

RUN mkdir myapp

WORKDIR /myapp

COPY go.mod /myapp/go.mod
COPY go.sum /myapp/go.sum

# go.modにあるライブラリをダウンロード
RUN go mod download

# 必要に応じて go.mod を整える
RUN go mod tidy

# ソースコードをコピー
COPY . .

CMD ["/bin/bash"]
