-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler  version: 0.9.1
-- PostgreSQL version: 9.5
-- Project Site: pgmodeler.io
-- Model Author: ---


-- Database creation must be done outside a multicommand file.
-- These commands were put in this file only as a convenience.
-- -- object: ceri | type: DATABASE --
-- -- DROP DATABASE IF EXISTS ceri;
-- CREATE DATABASE ceri;
-- -- ddl-end --
-- 

-- object: movilidad_academica | type: SCHEMA --
-- DROP SCHEMA IF EXISTS movilidad_academica CASCADE;
CREATE SCHEMA movilidad_academica;
-- ddl-end --
ALTER SCHEMA movilidad_academica OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,movilidad_academica;
-- ddl-end --

-- object: movilidad_academica.tipo_movilidad | type: TABLE --
-- DROP TABLE IF EXISTS movilidad_academica.tipo_movilidad CASCADE;
CREATE TABLE movilidad_academica.tipo_movilidad(
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar,
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	entrante bool NOT NULL,
	CONSTRAINT pk_tipo_movilidad PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE movilidad_academica.tipo_movilidad OWNER TO postgres;
-- ddl-end --

-- object: movilidad_academica.tipo_categoria | type: TABLE --
-- DROP TABLE IF EXISTS movilidad_academica.tipo_categoria CASCADE;
CREATE TABLE movilidad_academica.tipo_categoria(
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar,
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_tipo_categoria PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE movilidad_academica.tipo_categoria OWNER TO postgres;
-- ddl-end --

-- object: movilidad_academica.movilidad | type: TABLE --
-- DROP TABLE IF EXISTS movilidad_academica.movilidad CASCADE;
CREATE TABLE movilidad_academica.movilidad(
	id serial NOT NULL,
	presupuesto numeric(20,7) NOT NULL DEFAULT 0,
	fecha_inicio date NOT NULL,
	fecha_final date NOT NULL,
	persona integer NOT NULL,
	convenio integer NOT NULL,
	id_tipo_movilidad integer,
	id_tipo_categoria integer,
	CONSTRAINT pk_movilidad PRIMARY KEY (id)

);
-- ddl-end --
ALTER TABLE movilidad_academica.movilidad OWNER TO postgres;
-- ddl-end --

-- object: fk_movilidad_tipo_movilidad | type: CONSTRAINT --
-- ALTER TABLE movilidad_academica.movilidad DROP CONSTRAINT IF EXISTS fk_movilidad_tipo_movilidad CASCADE;
ALTER TABLE movilidad_academica.movilidad ADD CONSTRAINT fk_movilidad_tipo_movilidad FOREIGN KEY (id_tipo_movilidad)
REFERENCES movilidad_academica.tipo_movilidad (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_movilidad_tipo_categoria | type: CONSTRAINT --
-- ALTER TABLE movilidad_academica.movilidad DROP CONSTRAINT IF EXISTS fk_movilidad_tipo_categoria CASCADE;
ALTER TABLE movilidad_academica.movilidad ADD CONSTRAINT fk_movilidad_tipo_categoria FOREIGN KEY (id_tipo_categoria)
REFERENCES movilidad_academica.tipo_categoria (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --


