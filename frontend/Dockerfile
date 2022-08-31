FROM node:16-alpine as build
WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM nginx:1.23.1-alpine
COPY default.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist/todo-ui /usr/share/nginx/html
EXPOSE 80
