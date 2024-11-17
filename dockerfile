FROM golang:latest as build

WORKDIR /build

# Copy the Go module files
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN  CGO_ENABLED=0 go build -o main main.go


FROM alpine
EXPOSE 8081

# Copy the application executable from the build image
COPY --from=build /build/main /

ENTRYPOINT ["/main"]