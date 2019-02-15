FROM golang:1.11.5-alpine3.8 as build

WORKDIR /go/src/ping

COPY ping.go .
RUN go build


FROM golang:1.11.5-alpine3.8

WORKDIR /go/src/ping

COPY --from=build /go/src/ping/ping .
COPY .version .

EXPOSE 8993

CMD ["./ping"]