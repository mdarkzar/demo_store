FROM node:16.20-alpine3.16 as build-front

WORKDIR /app
COPY frontend/package.json /app

RUN yarn install 

COPY frontend /app

RUN yarn build

FROM alpine:latest
RUN apk add --no-cache bash

COPY --from=build-front /app/dist/assets /app/assets
COPY --from=build-front /app/dist/index.html /app/index.html
COPY docker/site/entrypoint.sh /app/entrypoint.sh

CMD "/app/entrypoint.sh"