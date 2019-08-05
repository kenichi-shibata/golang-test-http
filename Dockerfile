FROM golang:1.12 as BUILD
ADD . /src
RUN cd /src && go build -o golang-http-test
ENTRYPOINT ./golang-http-test

FROM alpine
WORKDIR /app
COPY --from=build /src/golang-http-test /app/
ENTRYPOINT ./golang-http-test