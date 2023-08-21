FROM golang:1.21-alpine
RUN mkdir /build
WORKDIR /build
COPY . .
ENV GOOS=linux CGO_ENABLED=0
RUN set -ex && \
    apk add --no-progress --no-cache \
    gcc \
    musl-dev

RUN go build -o server cmd/main.go

FROM alpine:3.15.0
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=0 /build/server /usr/bin/
COPY .env .

EXPOSE 8080
EXPOSE 5432

CMD [ "server" ]