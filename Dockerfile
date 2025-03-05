FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o geo_service cmd/main.go

EXPOSE 8080

CMD [ "./geo_service" ]