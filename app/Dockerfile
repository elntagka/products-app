FROM golang:1.15.5-buster as go-builder
WORKDIR /module
COPY main.go go.mod go.sum /module/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch
COPY --from=go-builder /module/main /main
EXPOSE 8000/tcp
ENTRYPOINT ["/main"]