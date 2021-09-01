FROM node:14
WORKDIR /usr/src/morpheus/mailer/

COPY . .

RUN npm install --frozen-lockfile
RUN npm run build

COPY ./src/services/mail/templates/* ./dist/services/mail/templates/

ENTRYPOINT npm start