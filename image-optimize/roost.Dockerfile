FROM scratch
COPY app /app
ENTRYPOINT [ "/app/main" ]
