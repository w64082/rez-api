CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS public.workers
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    "name" text COLLATE pg_catalog."default" NOT NULL,
    surname text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    deleted_at timestamp without time zone,
    CONSTRAINT workers_pkey PRIMARY KEY (id)
)

CREATE TABLE IF NOT EXISTS public.places
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    "name" text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    deleted_at timestamp without time zone,
    CONSTRAINT places_pkey PRIMARY KEY (id)
)


