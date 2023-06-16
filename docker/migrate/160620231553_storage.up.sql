CREATE TABLE IF NOT EXISTS
    "store"."public".storage
    (
        st_id serial NOT NULL,
        name VARCHAR(255) NOT NULL,
        created_date TIMESTAMP WITH TIME zone DEFAULT now() NOT NULL,
        deleted_date TIMESTAMP WITH TIME zone,
        PRIMARY KEY (st_id),
        CONSTRAINT storage_ix1 UNIQUE (name)
    );