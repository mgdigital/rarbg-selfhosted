FROM golang:alpine3.18 AS build

RUN apk add gcc musl-dev

RUN mkdir /app

COPY . /app

RUN cd /app && go build

FROM alpine:3.18

COPY --from=build /app/rarbg-selfhosted /rarbg-selfhosted

ENTRYPOINT ["/rarbg-selfhosted"]
