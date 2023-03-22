FROM golang:1.20.2

WORKDIR app

COPY . ./

RUN go mod tidy

RUN go build -o ./MusicAPI cmd/main.go

CMD [ "/MusicAPI" ]