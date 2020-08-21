FROM node:10-slim
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"
WORKDIR /app
COPY . .
RUN npm install
EXPOSE 3000
ENTRYPOINT ["npm", "start"]
