# Turn on maintenance mode
php artisan down || true

# Clear config
php artisan config:clear

# Run database migrations
php artisan migrate --force

# Clear caches
php artisan cache:clear

# Clear and cache routes
php artisan route:cache

# Clear and cache config
php artisan config:cache

# Clear and cache views
php artisan view:cache

# Turn off maintenance mode
php artisan up
