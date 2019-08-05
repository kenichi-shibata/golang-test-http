FROM golang:1.12 as BUILD
ADD . /src
RUN cd /src && CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o golang-http-test ${PWD}

FROM alpine
WORKDIR /app
COPY --from=BUILD /src/golang-http-test /app/
RUN chmod +x /app/golang-http-test
ENTRYPOINT [ "/app/golang-http-test" ]
