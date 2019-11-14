FROM golang:1.10.0-alpine3.7

COPY goapi /app

WORKDIR /app

RUN \
    go build -o goapi
    
EXPOSE 1337

CMD ["./goapi"]