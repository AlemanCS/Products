FROM golang:1.10-alpine
WORKDIR /go/src/products
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/products/app .
CMD ["./app"]
EXPOSE 3000
