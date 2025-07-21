
CREATE DATABASE greenlight;

CREATE ROLE greenlight WITH LOGIN PASSWORD 'password';

CREATE EXTENSION IF NOT EXISTS citext;

psql --host=localhost --dbname=greenlight --username=greenlight