BEGIN;
CREATE TYPE user_role AS ENUM (
    'ADMIN',
    'USER'
);

CREATE TABLE users
(
    id          BIGINT          NOT NULL    GENERATED BY DEFAULT AS IDENTITY,
    name        VARCHAR(100)    NOT NULL,
    phone       VARCHAR(100)    NOT NULL,
    email       VARCHAR(100)    NOT NULL,
    role        user_role       NOT NULL    DEFAULT 'USER',
    password    VARCHAR(100)    NOT NULL,
    
    created_at  TIMESTAMP       NOT NULL    DEFAULT (now() at time zone 'utc'),
    updated_at  TIMESTAMP       NOT NULL    DEFAULT (now() at time zone 'utc'),

    PRIMARY KEY (id)
);

CREATE UNIQUE INDEX user_email_unique ON users (email);
COMMIT;
