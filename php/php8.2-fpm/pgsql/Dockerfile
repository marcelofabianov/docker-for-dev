FROM php:8.2-fpm

ENV TZ=America/Sao_Paulo
ENV DEBIAN_FRONTEND noninteractive

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN rm /etc/apt/preferences.d/no-debian-php && \
    apt-get update -yqq && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    git \
    wget \
    curl \
    unzip \
    zip \
    openssl \
    libpq-dev \
    libaio1 \
    libaio-dev \
    zlib1g-dev \
    libonig-dev \
    libzip-dev \
    libicu-dev \
    libssl-dev \
    libxml2-dev \
    libmemcached-dev

RUN docker-php-ext-configure pgsql --with-pgsql=/usr/local/pgsql && \
    docker-php-ext-configure intl && \
    docker-php-ext-install \
    intl \
    zip \
    mbstring \
    pcntl \
    sockets \
    exif \
    bcmath \
    soap \
    simplexml \
    pgsql \
    pdo_pgsql

RUN pecl install \
    redis \
    apcu \
    xdebug \
    swoole && \
    docker-php-ext-enable \
    xdebug \
    redis \
    opcache \
    swoole \
    apcu && \
    pecl clear-cache

COPY php.ini /usr/local/etc/php/

RUN curl -L -C - --progress-bar -o /usr/local/bin/composer https://getcomposer.org/composer.phar && \
    chmod 755 /usr/local/bin/composer && \
    useradd -m devs && \
    usermod -u 1000 devs

RUN composer self-update

RUN rm -rf /var/lib/apt/lists/*

COPY --chown=devs:www-data . /var/www

RUN chown -R www-data /var/www

USER devs

EXPOSE 9000
EXPOSE 9003

CMD ["php-fpm"]
