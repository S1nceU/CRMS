# Build stage
FROM golang:1.21.3-alpine AS builderStage
WORKDIR /go/src
ADD . /go/src
RUN cd /go/src
RUN go mod download
RUN go build -o main .

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builderStage /go/src/main /app/
EXPOSE 8080
ENTRYPOINT ./main