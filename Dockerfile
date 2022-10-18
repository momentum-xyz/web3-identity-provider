FROM golang:1.18-alpine as build

RUN apk add --update make gcc musl-dev

WORKDIR /build

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN make build


FROM alpine:latest as runtime

WORKDIR /opt/web3-identify-provider
COPY --from=build /build/web3-identity-provider .

ENV WEB3_IDP_BIND_ADDRESS=:4000
EXPOSE 4000
CMD ["/opt/web3-identify-provider/web3-identity-provider"]


