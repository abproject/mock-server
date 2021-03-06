FROM golang:1.11.2 AS build
WORKDIR $GOPATH/src/github.com/abproject/mock-server
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./.bin/mock-server .
RUN ls -lh ./.bin/mock-server | cut -d' ' -f5

FROM scratch
COPY --from=build /go/src/github.com/abproject/mock-server/.bin/mock-server mock-server
EXPOSE 8000
ENTRYPOINT ["/mock-server"]