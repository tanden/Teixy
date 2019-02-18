# support toolのdeploy環境
FROM node:8.11.3-alpine
WORKDIR /support
RUN apk update && \
    npm install -g firebase-tools