FROM golang:latest

WORKDIR /firestore-connect

COPY . .

RUN go build -o /go/bin/firestore-connect main.go

CMD ["firebase-connect"]