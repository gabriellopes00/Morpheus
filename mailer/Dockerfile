FROM node:14
WORKDIR /usr/src/morpheus/mailer/

COPY . .

RUN npm install --frozen-lockfile
RUN npm run build

ENTRYPOINT npm start