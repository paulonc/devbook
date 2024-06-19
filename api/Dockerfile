FROM golang:1.22.4-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/bin/devbook

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=build /app/bin/devbook /usr/local/bin/devbook
COPY --from=build /app/src/migrations/sql /src/migrations/sql
COPY .env /

ENTRYPOINT ["/usr/local/bin/devbook"]

EXPOSE 5000