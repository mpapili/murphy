# syntax=docker/dockerfile:1
FROM docker.io/library/golang:1.22-alpine AS build
WORKDIR /src
RUN apk add --no-cache git ca-certificates
COPY go.mod ./
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/murphy ./cmd/murphy

FROM docker.io/library/alpine:3.20
RUN apk add --no-cache ca-certificates
COPY --from=build /out/murphy /murphy
EXPOSE 8091
ENTRYPOINT ["/murphy"]
