FROM node:18 AS base

RUN npm i -g pnpm

FROM base AS build
WORKDIR /app
COPY package.json pnpm-lock.yaml ./
RUN pnpm install
COPY . .
RUN pnpm run build

FROM node:18-alpine
WORKDIR /www
COPY --from=build /app/build .
COPY --from=build /app/package.json .
COPY --from=build /app/node_modules ./node_modules
EXPOSE 3000
CMD ["node", "index.js"]
