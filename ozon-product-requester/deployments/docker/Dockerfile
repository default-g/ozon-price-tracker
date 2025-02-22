
FROM golang:1.24 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/ozon-product-requester/ozon-product-requester.go

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/ozon-product-requester /ozon-product-requester

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/ozon-product-requester"]
