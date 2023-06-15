-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

BEGIN;

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = public, extensions;
SET default_tablespace = '';
SET default_with_oids = FALSE;

-- EXTENSIONS --

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- TABLES --

CREATE TABLE public.currency
(
    id     SERIAL PRIMARY KEY,
    name   TEXT,
    symbol TEXT
);

CREATE TABLE public.category
(
    id   SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE public.product
(
    id            UUID PRIMARY KEY,
    name          TEXT NOT NULL,
    description   TEXT NOT NULL,
    image_id      UUID,
    price         BIGINT NOT NULL,
    currency_id   INT REFERENCES public.currency (id) NOT NULL,
    rating        INT,
    category_id   INT REFERENCES public.category (id) NOT NULL,
    specification JSONB,
    created_at    TIMESTAMPTZ NOT NULL,
    updated_at    TIMESTAMPTZ,
    CONSTRAINT positive_price CHECK (price > 0),
    CONSTRAINT valid_rating CHECK (rating <= 5)
);

COMMIT;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

COMMIT;