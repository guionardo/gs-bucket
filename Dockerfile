FROM golang:1.21-alpine as backend-dev

WORKDIR /app
RUN apk add build-base

# Copy go mod and sum files
COPY ./go.mod ./go.sum ./

# Download all required packages
RUN go mod download

# COPY gs-bucket.go .
COPY main.go .
COPY backend ./backend/
COPY domain ./domain/
COPY README.md .

RUN GOOS=linux go build -o gs-bucket -ldflags="-X 'main.Build=$(date +%Y-%m-%dT%H:%M:%S%z)'" .
#RUN CGO_ENABLED=1 GOOS=linux go build -o gs-bucket  -ldflags="-X 'main.Build=$(date +%Y-%m-%dT%H:%M:%S%z)'" .


FROM alpine:latest AS final-build
# Ensure updated CA certificates
RUN apk --no-cache add ca-certificates

# needed for timezones
RUN apk add --no-cache tzdata

WORKDIR /root/
COPY .github/scripts/create_masterkey.sh .
COPY --from=backend-dev /app/gs-bucket .

RUN mkdir /bucket

RUN sh create_masterkey.sh

CMD ["./gs-bucket"]
