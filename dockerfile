FROM golang:latest as build

WORKDIR /build

COPY . .
RUN go mod download
RUN  CGO_ENABLED=0 go build -o main main.go


FROM alpine

# Copy the application executable from the build image
COPY --from=build /build/main /

ENTRYPOINT ["/main"]