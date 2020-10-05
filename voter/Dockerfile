# pull official base image
FROM node:13.12.0-alpine as builder
LABEL maintainer="mgdevstack" \
    vendor="Zettabytes" \
    owner="zbio"

# set working directory
WORKDIR /app

# install app dependencies
COPY package*.json ./
RUN npm install --silent 

# add app
COPY . .

RUN npm run build 

FROM nginx

# map this port to any arbitrary port
EXPOSE 80

# copy all the build files from the previous step
COPY --from=builder /app/build /usr/share/nginx/html

