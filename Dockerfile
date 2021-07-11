FROM golang:1.16-buster AS builder

RUN mkdir /app
WORKDIR /app

COPY . .

RUN make all

CMD ["air"]


FROM gcr.io/distroless/static-debian10

COPY --from=builder --chown=nonroot:nonroot /app/pikuchat /

USER nonroot

CMD ["/pikuchat"]
