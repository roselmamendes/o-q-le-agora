FROM golang:1.9

WORKDIR /go/src/app
COPY . .

# dependencies
RUN go get -v github.com/codegangsta/gin
RUN go get -v github.com/stretchr/testify

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."


EXPOSE 4000
