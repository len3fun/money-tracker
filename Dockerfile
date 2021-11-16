FROM golang:1.17 AS build

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o .bin/app cmd/main.go

FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/.bin/app .
COPY --from=build /app/configs configs
COPY --from=build /app/schema schema
COPY --from=build /app/.env .
CMD ["./app"]
