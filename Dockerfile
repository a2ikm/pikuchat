FROM golang:1.16-buster

RUN mkdir /app
WORKDIR /app

COPY . .

RUN make all

CMD ["air"]
