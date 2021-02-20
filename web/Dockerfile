FROM node:latest as build-stage

MAINTAINER drunk_kk

RUN npm install -g cnpm --registry=https://registry.npm.taobao.org

WORKDIR /app

COPY . .

RUN rm -rf node_modules \
    && npm install node-sass \
    && cnpm run build

FROM nginx:1.15.3-alpine as production-stage

COPY ./nginx.conf /etc/nginx/nginx.conf

COPY --from=build-stage /app/dist /usr/share/nginx/html

EXPOSE 8001

CMD ["nginx", "-g", "daemon off;"]
