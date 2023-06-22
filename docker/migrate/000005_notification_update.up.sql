DROP TABLE IF EXISTS
     "store"."public".notification;

CREATE TABLE IF NOT EXISTS
    "store"."public".notification
    (
        n_id bigserial NOT NULL,
        title VARCHAR(1000) NOT NULL,
        MESSAGE text NOT NULL,
        created_date TIMESTAMP WITH TIME zone DEFAULT now() NOT NULL,
        PRIMARY KEY (n_id)
    );


CREATE TABLE IF NOT EXISTS
    "store"."public".user$notification
    (
        user_id INTEGER NOT NULL,
        n_id integer not null,
        PRIMARY KEY (user_id, n_id),
        CONSTRAINT usernotification_fk1 FOREIGN KEY (user_id) REFERENCES "store"."public"."users"
        ("user_id"),
         CONSTRAINT usernotification_fk2 FOREIGN KEY (n_id) REFERENCES "store"."public"."notification"
        ("n_id")
    );
