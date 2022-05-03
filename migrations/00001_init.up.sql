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

COMMIT;