FROM golang:alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
ADD main.go  .
RUN CGO_ENABLED=0 GOOS=linux go build -o ballot main.go

FROM alpine:3.9
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
COPY --from=builder /go/ballot /app/
EXPOSE 82
ENTRYPOINT [ "/app/ballot" ]