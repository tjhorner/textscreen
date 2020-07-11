FROM golang:1.14 AS builder

ENV GO111MODULE=on

RUN apt-get update && apt-get install upx -y

WORKDIR /app/server
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o /textscreen -ldflags "-linkmode external -extldflags -static" -a .

RUN upx /textscreen

FROM gcr.io/distroless/static:8bef63d2c8654ff89358430c7df5778162ab6027

VOLUME [ "/config.yml" ]

COPY --from=builder /textscreen /textscreen
CMD [ "./textscreen" ]