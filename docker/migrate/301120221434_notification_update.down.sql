DROP TABLE
    "public"."user$notification";

DROP TABLE
    "public"."notification";

CREATE TABLE
    "store"."public".notification
    (
        n_id bigserial NOT NULL,
        user_id INTEGER NOT NULL,
        title VARCHAR(1000) NOT NULL,
        MESSAGE text NOT NULL,
        created_date TIMESTAMP WITH TIME zone DEFAULT now() NOT NULL,
        PRIMARY KEY (n_id),
        CONSTRAINT notification_fk1 FOREIGN KEY (user_id) REFERENCES "store"."public"."users"
        ("user_id")
    );


CREATE INDEX
    notification_user_inx
ON
    "public"."notification"
    (
        "user_id"
    );

