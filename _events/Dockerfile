
FROM node:lts-alpine AS dev_dependencies
WORKDIR /app/
COPY package.json yarn.* tsconfig.json ./
COPY ./src ./src
RUN yarn install --production=false --frozen-lockfile

FROM node:lts-alpine AS dependencies
WORKDIR /app/
COPY package.json yarn.* ./
COPY ./src ./src
RUN yarn install --production=true --frozen-lockfile

FROM node:lts-alpine AS build
WORKDIR /app/
COPY --from=dev_dependencies /app/ .
COPY . .
RUN yarn build

FROM node:lts-alpine AS runtime
USER node
WORKDIR /usr/src/morpheus_events/
ENV PATH /usr/src/morpheus_events/node_modules/.bin:$PATH
ENV NODE_ENV=production
COPY --chown=node:node --from=dependencies /app/node_modules /usr/src/morpheus_events/node_modules/
COPY --from=build --chown=node:node /app/dist /usr/src/morpheus_events/dist/
COPY --from=build --chown=node:node /app/package.json /usr/src/morpheus_events
ENTRYPOINT yarn start