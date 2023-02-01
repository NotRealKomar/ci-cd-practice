FROM golang:1.20rc3-alpine3.17 as builder
WORKDIR /app
COPY go.* ./
COPY . .
RUN go mod download && go build ./...

FROM alpine:3.17.1
WORKDIR /app

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH="${GOROOT}/bin:${PATH}"

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

COPY --from=builder /usr/local/go/ /usr/local/go/
COPY --from=builder /go/pkg/mod /go/pkg/mod
COPY --from=builder /app .

EXPOSE 3000
CMD go run ./src/main.go
