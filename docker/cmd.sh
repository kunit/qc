#!/bin/bash

cd /var/www/html
git clone https://github.com/kunit/qc-sample-laravel-app.git app
cd app
composer install

cp /var/www/html/.env /var/www/html/app
php artisan migrate

if [ ! -e /var/log/tcpdp ]; then
    mkdir /var/log/tcpdp
fi
