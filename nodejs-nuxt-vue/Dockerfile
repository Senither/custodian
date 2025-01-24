# Stage 1: Setup the base image
FROM node:22-alpine AS base

ENV NODE_ENV=production
ENV DATABASE_URL="file:/app/database.sqlite"

WORKDIR /app

# Stage 2: Build the assets and prepare the database
FROM base AS build

COPY . .

RUN npm install
RUN npm run db:migrate:deploy
RUN npm run build

# Stage 3: Copy the built assets & run the application
FROM base AS production

COPY --from=build /app/.output /app/database.sqlite /app/

EXPOSE 3000

CMD ["node", "server/index.mjs"]
