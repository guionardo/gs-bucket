FROM golang:1.19-alpine as backend-dev

WORKDIR /app
RUN apk add build-base

# Copy go mod and sum files
COPY ./go.mod ./go.sum ./

# Download all required packages
RUN go mod download

COPY gs-bucket.go .
COPY cmd ./cmd/
COPY pkg ./pkg/
COPY README.md .

RUN GOOS=linux go build -o gs-bucket .
#RUN CGO_ENABLED=1 GOOS=linux go build -o gs-bucket  -ldflags="-X 'main.Build=$(date +%Y-%m-%dT%H:%M:%S%z)'" cmd/gs-bucket.go


FROM alpine:latest as final-build
# FROM cgr.dev/chainguard/static:latest as final-build
# Ensure updated CA certificates
RUN apk --no-cache add ca-certificates

# needed for timezones
RUN apk add --no-cache tzdata

WORKDIR /root/
COPY --from=backend-dev /app/gs-bucket .
CMD ["./gs-bucket", "--data-path", "./data"]