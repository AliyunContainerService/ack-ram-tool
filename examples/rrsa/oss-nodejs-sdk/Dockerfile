FROM node:19-alpine3.19 as builder
# TARGETPLATFORM

RUN npm set registry https://registry.npmmirror.com

WORKDIR /app
COPY package.json ./
COPY package-lock.json ./

RUN npm install

FROM node:19-alpine3.19

WORKDIR /app
COPY package.json ./
COPY package-lock.json ./
COPY src/index.js ./src/
COPY --from=builder /app/node_modules ./node_modules

CMD npm run demo
