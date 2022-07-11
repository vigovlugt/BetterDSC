FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build

FROM gcr.io/distroless/base-debian10 as runtime

WORKDIR /

COPY --from=build /app/betterdsc /betterdsc


ENTRYPOINT ["./betterdsc"]