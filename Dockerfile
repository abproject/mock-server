FROM golang:1.11.2 AS build
WORKDIR $GOPATH/src/mock-server
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN "./scripts/compile.sh"

FROM scratch
COPY --from=build /go/src/mock-server/.bin/mock-server mock-server
EXPOSE 8000
ENTRYPOINT ["/mock-server"]