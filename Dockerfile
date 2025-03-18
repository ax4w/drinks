FROM golang:1.24.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .

RUN templ generate
RUN go build -o server .

FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=build /app/server /app/

EXPOSE 8080

USER appuser

CMD ["./server"]