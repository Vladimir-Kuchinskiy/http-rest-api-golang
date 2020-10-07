#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER http_golang WITH ENCRYPTED PASSWORD 'password';
    CREATE DATABASE http_golang;
    GRANT ALL PRIVILEGES ON DATABASE http_golang TO http_golang;

    CREATE USER http_golang_test WITH ENCRYPTED PASSWORD 'password';
    CREATE DATABASE http_golang_test;
    GRANT ALL PRIVILEGES ON DATABASE http_golang_test TO http_golang_test;
EOSQL
