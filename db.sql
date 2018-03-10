-- Database: todo_db

--  DROP DATABASE todo_db;

CREATE DATABASE todo_db
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;
	
-- SCHEMA: public

-- DROP SCHEMA public ;

CREATE SCHEMA public
    AUTHORIZATION postgres;

COMMENT ON SCHEMA public
    IS 'standard public schema';

GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;

-- drop sequence if exists todos_id_seq;
-- create sequence todos_id_seq;
-- drop sequence if exists users_id_seq;
-- create sequence users_id_seq;
-- drop sequence if exists lists_id_seq;
-- create sequence lists_id_seq;
CREATE SEQUENCE public.lists_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.lists_id_seq
    OWNER TO postgres;

CREATE SEQUENCE public.todos_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.todos_id_seq
    OWNER TO postgres;

CREATE SEQUENCE public.users_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1;

ALTER SEQUENCE public.users_id_seq
    OWNER TO postgres;

-- Table: public.users

-- DROP TABLE public.users;

CREATE TABLE public.users
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    login character varying(128) COLLATE pg_catalog."default" NOT NULL,
    password character varying(256) COLLATE pg_catalog."default" NOT NULL,
    email character varying(256) COLLATE pg_catalog."default" NOT NULL,
    register_date timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT users_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to postgres;

-- Index: users_email_uindex

-- DROP INDEX public.users_email_uindex;

CREATE UNIQUE INDEX users_email_uindex
    ON public.users USING btree
    (email COLLATE pg_catalog."default")
    TABLESPACE pg_default;

-- Index: users_id_uindex

-- DROP INDEX public.users_id_uindex;

CREATE UNIQUE INDEX users_id_uindex
    ON public.users USING btree
    (id)
    TABLESPACE pg_default;

-- Index: users_login_uindex

-- DROP INDEX public.users_login_uindex;

CREATE UNIQUE INDEX users_login_uindex
    ON public.users USING btree
    (login COLLATE pg_catalog."default")
    TABLESPACE pg_default;

-- Table: public.lists

-- DROP TABLE public.lists;

CREATE TABLE public.lists
(
    id integer NOT NULL DEFAULT nextval('lists_id_seq'::regclass),
    name character varying(256) COLLATE pg_catalog."default" NOT NULL,
    user_id integer NOT NULL,
    CONSTRAINT lists_pkey PRIMARY KEY (id),
    CONSTRAINT lists_users_id_fk FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.lists
    OWNER to postgres;

-- Index: lists_id_uindex

-- DROP INDEX public.lists_id_uindex;

CREATE UNIQUE INDEX lists_id_uindex
    ON public.lists USING btree
    (id)
    TABLESPACE pg_default;

-- Index: lists_user_id_index

-- DROP INDEX public.lists_user_id_index;

CREATE INDEX lists_user_id_index
    ON public.lists USING btree
    (user_id)
    TABLESPACE pg_default;

-- Table: public.todos

-- DROP TABLE public.todos;

CREATE TABLE public.todos
(
    id integer NOT NULL DEFAULT nextval('todos_id_seq'::regclass),
    title character varying(256) COLLATE pg_catalog."default" NOT NULL,
    description character varying(2048) COLLATE pg_catalog."default",
    list_id integer,
    is_active boolean DEFAULT true,
    user_id integer,
    date_create date NOT NULL DEFAULT now(),
    CONSTRAINT todos_pkey PRIMARY KEY (id),
    CONSTRAINT todos_lists_id_fk FOREIGN KEY (list_id)
        REFERENCES public.lists (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT todos_users_id_fk FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.todos
    OWNER to postgres;

-- Index: todos_id_uindex

-- DROP INDEX public.todos_id_uindex;

CREATE UNIQUE INDEX todos_id_uindex
    ON public.todos USING btree
    (id)
    TABLESPACE pg_default;

-- Index: todos_user

-- DROP INDEX public.todos_user;

CREATE INDEX todos_user
    ON public.todos USING btree
    (user_id)
    TABLESPACE pg_default;

-- Index: todos_user_list

-- DROP INDEX public.todos_user_list;

CREATE INDEX todos_user_list
    ON public.todos USING btree
    (user_id, list_id)
    TABLESPACE pg_default;
