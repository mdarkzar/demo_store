CREATE TABLE
    "store"."public".product
    (
        product_id serial NOT NULL,
        name VARCHAR(255) NOT NULL,
        price DOUBLE PRECISION NOT NULL,
        creator_id INTEGER NOT NULL,
        created_date TIMESTAMP WITH TIME zone DEFAULT now() NOT NULL,
        deleted_date TIMESTAMP WITH TIME zone,
        PRIMARY KEY (product_id),
        CONSTRAINT product_fk1 FOREIGN KEY (creator_id) REFERENCES "store"."public"."users"
        ("user_id")
    );
