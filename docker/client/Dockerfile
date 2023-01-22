FROM node:14.18.0
ENV ROOT=/usr/src/app

WORKDIR ${ROOT}
COPY ./client/package.json ./client/yarn.lock ./
RUN yarn install

COPY ./client ./
CMD ["yarn", "dev"]
