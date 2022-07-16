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

CREATE TABLE IF NOT EXISTS public.visits
(
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    id_place uuid NOT NULL,
    id_worker uuid NOT NULL,
    date_start timestamp without time zone NOT NULL,
    date_to timestamp without time zone NOT NULL,
    is_reserved boolean NOT NULL DEFAULT false,
    client_name text COLLATE pg_catalog."default",
    client_surname text COLLATE pg_catalog."default",
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    deleted_at timestamp without time zone,
    CONSTRAINT visits_pkey PRIMARY KEY (id),
    CONSTRAINT fkey_id_place_places FOREIGN KEY (id_place)
    REFERENCES public.places (id) MATCH SIMPLE
                         ON UPDATE CASCADE
                         ON DELETE RESTRICT,
    CONSTRAINT fkey_id_worker_workers FOREIGN KEY (id_worker)
    REFERENCES public.workers (id) MATCH SIMPLE
                         ON UPDATE CASCADE
                         ON DELETE RESTRICT
)
