FROM golang:1.20.4 as build

WORKDIR /app

COPY . .

RUN make linux

FROM alpine:3.14 as prod

WORKDIR /app

COPY --from=build /app/backend backend
COPY --from=build /app/config.toml config.toml
RUN chmod +x /app/backend

CMD ["/app/backend" ]