FROM golang:1.21.0-alpine3.17 as build
WORKDIR /api
COPY go.* .
RUN go mod tidy
COPY . .
RUN go build -o /api/demo ./cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=build /api/demo .
EXPOSE 1323
CMD ["./demo"]