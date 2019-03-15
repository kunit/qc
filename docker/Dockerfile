FROM php:7.3-apache

COPY ./apache/000-default.conf /etc/apache2/sites-available/000-default.conf
COPY ./php/php.ini /usr/local/etc/php/php.ini

ENV DOCKERIZE_VERSION v0.6.1

RUN apt-get update \
  && apt-get install -y curl wget git libzip-dev libicu-dev zlib1g-dev mysql-client \
  && docker-php-ext-install pdo_mysql zip intl mbstring opcache \
  && pecl install xdebug-2.7.0 \
  && docker-php-ext-enable xdebug \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* \
  && curl -sSL https://getcomposer.org/installer | php \
  && mv composer.phar /usr/local/bin/composer \
  && a2enmod rewrite \
  && wget https://github.com/k1LoW/tcpdp/releases/download/v0.17.0/tcpdp_v0.17.0_linux_amd64_static.xenial.tar.gz -P /tmp \
  && cd /tmp; tar xvf tcpdp_v0.17.0_linux_amd64_static.xenial.tar.gz \
  && cp /tmp/tcpdp_v0.17.0_linux_amd64_static.xenial/tcpdp /usr/local/bin \
  && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
  && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz

