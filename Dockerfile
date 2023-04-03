FROM golang:1.20.2

WORKDIR /app

COPY go.mod .
RUN go mod download


COPY . ./

RUN go build -o ./MusicAPI cmd/main.go

CMD [ "./MusicAPI" ]