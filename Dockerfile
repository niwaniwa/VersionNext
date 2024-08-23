FROM golang:1.23.0

WORKDIR /app

RUN apt install git -y

# ソースコードのコピー
COPY . .

RUN go mod tidy
RUN go install github.com/air-verse/air@latest

CMD ["air"]