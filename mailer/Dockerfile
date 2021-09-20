FROM node:14
WORKDIR /usr/src/morpheus/mailer/

COPY package.json .
COPY yarn.lock .
RUN npm install --frozen-lockfile

COPY . .
RUN npm run build

ENTRYPOINT npm start
