# Stage 1: Build the assets and prepare the database
FROM node:22-alpine AS build

ENV NODE_ENV=production
ENV DATABASE_URL="file:/app/database.sqlite"

WORKDIR /app

COPY . .

RUN npm install
RUN npm run db:migrate:deploy
RUN npm run build

# Stage 2: Copy the built assets & run the application
FROM node:22-alpine AS production

ENV NODE_ENV=production
ENV DATABASE_URL="file:/app/database.sqlite"

WORKDIR /app

COPY --from=build /app/.output /app/database.sqlite /app/

EXPOSE 3000

CMD ["node", "server/index.mjs"]
