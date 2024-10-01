FROM golang:1.23 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/reflect

FROM alpine:latest

COPY --from=build /go/bin/reflect /
ENTRYPOINT ["/reflect"]
CMD ["help"]