# Vue.jsのビルド環境
FROM node:10.15.0-alpine
WORKDIR /vue
RUN apk update && \
    npm install -g @vue/cli && \
    npm uninstall yarn -g && \
    npm install yarn -g