FROM golang:1.19.6-alpine3.17 AS build

WORKDIR /go/src/clean_architecture_template
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/clean_architecture_template cmd/clean_architecture_template.go

FROM scratch

ENV GOPROXY=https://proxy.golang.org
ENV GIN_MODE=release

COPY --from=build /go/bin/clean_architecture_template /app

ENTRYPOINT ["/app"]
