FROM golang:1.20-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go mod download
COPY . .
RUN go build -o /bin/nebula-grid ./cmd/cloudmesh

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /bin/nebula-grid /usr/local/bin/nebula-grid
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/nebula-grid"]
