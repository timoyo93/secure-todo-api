FROM node:16-alpine AS build

WORKDIR /app
COPY package*.json ./
RUN npm ci
COPY . .
RUN npm run build

FROM node:16-alpine
WORKDIR /app
COPY --from=build /app/build build/
COPY --from=build /app/node_modules node_modules/
COPY package.json .
EXPOSE 3000
ENV NODE_ENV=production

CMD [ "node", "build" ]