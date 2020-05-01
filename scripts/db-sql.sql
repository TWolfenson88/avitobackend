
CREATE TABLE IF NOT EXISTS profile
(
        uid             SERIAL PRIMARY KEY,
        name            VARCHAR(30)        NOT NULL check ( name <> '' ),
        email           VARCHAR(64),
        ident           VARCHAR(128),
        status          BOOLEAN,
        password        BYTEA              NOT NULL CHECK ( octet_length(password) <> 0 )
);

CREATE TABLE IF NOT EXISTS friend
(
        uid             SERIAL PRIMARY KEY,
        first           INTEGER REFERENCES profile (uid),
        second          INTEGER REFERENCES profile (uid)
);



postgres=> CREATE TABLE IF NOT EXISTS login
(
        id              SERIAL PRIMARY KEY DEFAULT,
        sess_id         UUID,
        user_id         INTEGER REFERENCES profile (uid),
        user_agent      varchar(128),
        add_time        TIMESTAMPTZ NOT NULL,
        ip_addr         varchar(24)
);