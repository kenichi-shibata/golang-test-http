FROM golang:1.12 as BUILD
ADD . /src
RUN cd /src && go build -o golang-http-test

FROM alpine
WORKDIR /app
COPY --from=BUILD /src/golang-http-test /app/
RUN chmod +x /app/golang-http-test
ENTRYPOINT /app/golang-http-test