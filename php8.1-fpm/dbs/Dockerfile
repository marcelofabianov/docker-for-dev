FROM php:8.1.0-fpm

ENV TZ=America/Sao_Paulo \
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
    git \
    openssl \
    libpq-dev \
    libaio1 \
    libaio-dev \
    zlib1g-dev \
    libonig-dev \
    libzip-dev \
    libicu-dev \
    libxml2-dev \
    gnupg2 \
    libodbc1 \
    nano \
    openssl

## ODBC Oracle
COPY ./instantclient_19_3.zip /tmp/instantclient.zip

RUN unzip /tmp/instantclient.zip -d /usr/local/ && \
    ln -s /usr/local/instantclient/sqlplus /usr/bin/sqlplus && \
    rm /tmp/instantclient.zip

RUN docker-php-ext-configure intl && \
    docker-php-ext-configure pgsql --with-pgsql=/usr/local/pgsql && \
    docker-php-ext-configure pdo_oci --with-pdo-oci=instantclient,/usr/local/instantclient && \
      echo 'instantclient,/usr/local/instantclient' | pecl install oci8-3.2.1 && \
      pecl install redis && \
      pecl install apcu && \
    docker-php-ext-install \
      pdo_mysql \
      pgsql \
      pdo_pgsql \
      pdo_oci \
      intl \
      zip \
      mbstring \
      pcntl \
      soap \
      simplexml \
      bcmath \
      exif \
      sockets && \
    docker-php-ext-enable \
      redis \
      opcache \
      apcu \
      oci8 && \
    pecl clear-cache

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
