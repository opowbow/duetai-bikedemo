FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /london-bike-stations

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /london-bike-stations /london-bike-stations
COPY *.html ./

USER nonroot:nonroot

ENTRYPOINT ["/london-bike-stations"]