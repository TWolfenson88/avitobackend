-- DROP TABLE if exists login;
-- DROP TABLE if exists friend;
-- DROP TABLE if exists profile;
-- DROP TABLE if exists session;


CREATE TABLE IF NOT EXISTS profile
(
        uid             SERIAL PRIMARY KEY,
        name            VARCHAR(30)     UNIQUE  NOT NULL check ( name <> '' ),
        email           VARCHAR(64),
        ident           VARCHAR(128),
        status          BOOLEAN                 DEFAULT FALSE,
        password        BYTEA                   NOT NULL CHECK ( octet_length(password) <> 0 )
);

CREATE TABLE IF NOT EXISTS friend
(
        uid             SERIAL PRIMARY KEY,
        first           INTEGER REFERENCES profile (uid),
        second          INTEGER REFERENCES profile (uid)
);



-- CREATE TABLE IF NOT EXISTS session -- will be in-memory
-- (
--      sess_id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
--      user_id         INTEGER REFERENCES profile (uid),
--      user_agent      varchar(128),
--      add_time        TIMESTAMPTZ NOT NULL DEFAULT NOW()
-- );

CREATE TABLE IF NOT EXISTS login  -- not used now too
(
        id              SERIAL PRIMARY KEY,
        sess_id         UUID,
        user_id         INTEGER REFERENCES profile (uid),
        user_agent      varchar(128),
        add_time        TIMESTAMPTZ NOT NULL,
        ip_addr         varchar(24)
);