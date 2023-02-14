FROM php:8.1.0-fpm

ENV TZ='America/Sao_Paulo' \
    LD_LIBRARY_PATH='/usr/local/instantclient/' \
    ORACLE_HOME='/usr/local/instantclient/'

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN rm /etc/apt/preferences.d/no-debian-php && \
    apt-get update -yqq && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    wget \
    curl \
    unzip \
    zip \
    openssl \
    libpq-dev \
    libaio1 \
    libaio-dev \
    zlib1g-dev \
    libpng-dev \
    libjpeg-dev \
    libfreetype6-dev \
    libonig-dev \
    libxml2-dev \
    libzip-dev \
    git

RUN docker-php-ext-install \
    pdo_mysql \
    gd \
    zip \
    mbstring \
    pcntl \
    soap \
    simplexml \
    sockets \
    exif \
    bcmath \
    pecl install redis && \
    docker-php-ext-enable redis

## Instantclient oracle
COPY ./instantclient_19_3.zip /tmp/instantclient.zip

RUN unzip /tmp/instantclient.zip -d /usr/local/ && \
    ln -s /usr/local/instantclient/sqlplus /usr/bin/sqlplus && \
    rm /tmp/instantclient.zip -d

# OCI
RUN docker-php-ext-configure pdo_oci --with-pdo-oci=instantclient,/usr/local/instantclient && \
    echo 'instantclient,/usr/local/instantclient' | pecl install oci8-3.2.1 && \
    docker-php-ext-install pdo_oci && \
    docker-php-ext-enable oci8

# Composer
RUN curl -L -C - --progress-bar -o /usr/local/bin/composer https://getcomposer.org/composer.phar && \
    chmod 755 /usr/local/bin/composer && \
    useradd -m devs && \
    usermod -u 1000 devs

RUN composer self-update

COPY --chown=devs:www-data . /var/www

RUN chown -R www-data /var/www

USER devs

EXPOSE 9000

CMD ["php-fpm"]
