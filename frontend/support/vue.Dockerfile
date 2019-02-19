# Vue.jsのビルド環境
FROM node:10.15.0
WORKDIR /support
RUN apt-get update && \
    apt-get upgrade && \
    #vue-cli ver3のinstall
    npm install -g @vue/cli && \
    #yarnを最新にするために入れ直す
    npm uninstall yarn -g && \
    npm install yarn -g