FROM golang:1.9

WORKDIR /go/src/app
COPY . .

# dependencies
RUN go get -v github.com/PuerkitoBio/goquery
RUN go get -v github.com/stretchr/testify

EXPOSE 8080
