FROM php:8.2-cli

ENV TZ=America/Sao_Paulo
ENV DEBIAN_FRONTEND noninteractive
ENV LD_LIBRARY_PATH /usr/local/instantclient
ENV ORACLE_HOME /usr/local/instantclient

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
        libmemcached-dev \
        libodbc1

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

## ODBC Oracle Instantclient
COPY ./instantclient_19_3.zip /tmp/instantclient.zip

RUN unzip /tmp/instantclient.zip -d /usr/local/ && \
    ln -s /usr/local/instantclient/sqlplus /usr/bin/sqlplus && \
    rm /tmp/instantclient.zip -d

# OCI
RUN docker-php-ext-configure pdo_oci --with-pdo-oci=instantclient,/usr/local/instantclient && \
    echo 'instantclient,/usr/local/instantclient' | pecl install oci8-3.0.1 && \
    docker-php-ext-install pdo_oci && \
    docker-php-ext-enable oci8 && \
    pecl clear-cache

RUN rm -rf /var/lib/apt/lists/*

RUN curl -L -C - --progress-bar -o /usr/local/bin/composer https://getcomposer.org/composer.phar && \
    chmod 755 /usr/local/bin/composer

RUN composer self-update

RUN useradd -m app && usermod -u 1000 app

COPY --chown=app:www-data . /var/www

RUN chown -R www-data /var/www

USER app

EXPOSE 9000
