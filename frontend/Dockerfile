FROM node:8 AS builder
WORKDIR /usr/src/app
COPY package.json package-lock.json ./
RUN npm install --silent
COPY . .
RUN npx parcel build src/index.html

FROM nginx
COPY --from=builder /usr/src/app/dist /usr/share/nginx/html