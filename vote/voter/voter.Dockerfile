FROM golang:alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o voter main.go

FROM alpine:3.9
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
COPY --from=builder /go/static /app/
COPY --from=builder /go/index.html /app/
COPY --from=builder /go/voter /app/
ENV BALLOT_ENDPOINT ballot:82
WORKDIR /app
EXPOSE 81
ENTRYPOINT [ "./voter" ]