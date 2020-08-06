FROM golang:1.14-alpine as build
ADD ./src /go/src/app
WORKDIR /go/src/app
RUN go build -v

FROM alpine:3.12.0
COPY --from=build /go/src/app/app /bin/app
ENTRYPOINT ["/bin/app"]