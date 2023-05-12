FROM node:16.14.0 as build

WORKDIR /app

COPY webui webui

COPY package.json package.json

COPY nuxt.config.ts nuxt.config.ts

COPY windi.config.ts windi.config.ts

COPY yarn.lock yarn.lock

RUN yarn

RUN yarn build

FROM node:16.14.0-alpine as prod

WORKDIR /app

RUN ls -al

COPY --from=build /app/.output .output

CMD [ "node", "/app/.output/server/index.mjs" ]