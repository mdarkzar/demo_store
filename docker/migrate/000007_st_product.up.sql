ALTER TABLE public.product ADD IF NOT EXISTS st_id integer NOT NULL  DEFAULT 1;


ALTER TABLE public.product DROP CONSTRAINT IF EXISTS product_fk2;

ALTER TABLE public.product ADD CONSTRAINT  product_fk2  FOREIGN KEY (st_id)  REFERENCES public."storage"(st_id) ;
