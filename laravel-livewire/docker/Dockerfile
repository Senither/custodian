# Stage 1: Build Assets
FROM node:18 as assets-builder

WORKDIR /app

COPY package.json package-lock.json ./
COPY tailwind.config.js postcss.config.js vite.config.js ./

COPY resources ./resources
COPY public ./public

RUN npm install
RUN npm run build

# Stage 2: Install PHP dependencies
FROM composer:2 as php-builder

WORKDIR /app

COPY . .
RUN composer install --no-interaction --prefer-dist --optimize-autoloader --no-dev

# Stage 3: Production
FROM php:8.3-fpm-alpine

RUN apk add --no-cache nginx bash

WORKDIR /var/www/html

COPY --from=assets-builder /app/public ./public
COPY --from=php-builder /app ./

COPY docker/nginx.conf /etc/nginx/http.d/default.conf
COPY docker/crontab /etc/crontabs/root

RUN mv .env.docker .env && \
    sh docker/migrate.sh && \
    chown www-data:www-data -R /var/www/html

EXPOSE 80

CMD ["sh", "-c", "php-fpm & nginx -g 'daemon off;' & crond -f"]
