FROM node:14.18.0 as dev
ENV ROOT=/usr/src/app

WORKDIR ${ROOT}
COPY ./client/package.json ./client/yarn.lock ./
RUN yarn install

COPY ./client ./
CMD ["yarn", "dev"]

FROM node:20-alpine

WORKDIR /app
COPY ./client/package.json ./client/yarn.lock ./
RUN yarn install

COPY ./client ./
RUN yarn build
CMD ["yarn", "start"]
