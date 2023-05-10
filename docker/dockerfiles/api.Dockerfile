FROM golang:1.20.4 as build

WORKDIR /app

COPY example example

RUN go build example/main.go

FROM alpine:3.14 as prod

WORKDIR /app

COPY --from=build /app/main mengoaingu

CMD [ "mengoaingu" ]