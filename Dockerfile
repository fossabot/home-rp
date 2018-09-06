FROM golang:1.11 as rpdeper
WORKDIR /go/src/github.com/just1689/home-rp
COPY Gopkg.toml Gopkg.lock vendor ./
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -v -vendor-only

FROM rpdeper as builder
WORKDIR /go/src/github.com/just1689/home-rp
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o plancks-cloud .

FROM scratch
WORKDIR /
COPY --from=builder /go/src/github.com/just1689/home-rp/home-rp .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/home-rp"]