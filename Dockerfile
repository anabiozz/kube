FROM golang as golang
WORKDIR /go/src/echo

COPY server.go server.go

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o echo .

FROM scratch
LABEL maintainer="Alexey Bezrukov <alexeybezrukov2@gmail.com>"
COPY --from=golang /go/src/echo/echo /

EXPOSE 8080

ENTRYPOINT [ "./echo" ]