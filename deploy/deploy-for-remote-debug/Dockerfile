##
## Build
##
FROM golang:1.18-bullseye AS build


WORKDIR /build
RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY ../../go.mod .
COPY ../../go.sum .
RUN go mod download
COPY ../.. .
RUN CGO_ENABLED=0 GOOS=linux go build -gcflags "all=-N -l" -o usersapi .
WORKDIR /dist
RUN cp /build/usersapi .

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /go/bin/dlv /
COPY --from=build /dist/usersapi /
EXPOSE 8000 40000
CMD ["./dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--continue", "--accept-multiclient", "exec" ,"/usersapi"]