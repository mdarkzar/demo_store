CREATE TABLE IF NOT EXISTS
    "store"."public".users
    (
        user_id serial NOT NULL,
        login VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_date TIMESTAMP without TIME zone DEFAULT now() NOT NULL,
        PRIMARY KEY (user_id),
        CONSTRAINT users_ix1 UNIQUE (login)
    );
