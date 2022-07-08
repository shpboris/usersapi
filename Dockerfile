##
## Build
##
FROM golang:1.18-bullseye AS build


WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o usersapi .
WORKDIR /dist
RUN cp /build/usersapi .

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /dist/usersapi /
EXPOSE 8000
##USER nonroot:nonroot
ENTRYPOINT ["/usersapi"]