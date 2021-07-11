FROM golang:1.16-buster

RUN mkdir /app
WORKDIR /app

COPY . .

RUN make install-tools
RUN go build

CMD ["air"]
