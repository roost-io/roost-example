FROM node:alpine3.10
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
WORKDIR /app
ADD . .
RUN npm install
EXPOSE 3000
ENTRYPOINT ["/usr/local/bin/npm", "start"]
