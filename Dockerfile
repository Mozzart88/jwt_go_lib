FROM golang:1.19.0-alpine3.16 as build

RUN apk add --no-cache git

WORKDIR /src

COPY . ./
RUN go build -o jwt main.go 

CMD [ "/src/jwt" ]
