FROM node:12.18.2

RUN mkdir /app
WORKDIR /app

COPY ./mapviewer/ ./
WORKDIR ./mapviewer
RUN npm install
RUN npm install -g create-react-app
EXPOSE 3000

