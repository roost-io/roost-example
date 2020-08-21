FROM golang:alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
ADD main.go index.html ./
RUN CGO_ENABLED=0 GOOS=linux go build -o resultBoard main.go

FROM alpine:3.9
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
COPY --from=builder /go/resultBoard /app/
COPY --from=builder /go/index.html /app/
ENV BALLOT_ENDPOINT ballot:82
WORKDIR /app
EXPOSE 83
ENTRYPOINT [ "./resultBoard" ]