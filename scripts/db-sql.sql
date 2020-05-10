-- DROP TABLE if exists call_mark;
-- DROP TABLE if exists call;
-- DROP TABLE if exists login;
-- DROP TABLE if exists friend;
-- DROP TABLE if exists profile;
-- DROP TABLE if exists session;


CREATE TABLE IF NOT EXISTS profile
(
        uid             SERIAL PRIMARY KEY,
        name            VARCHAR(30)     UNIQUE  NOT NULL check ( name <> '' ),
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

-- CREATE TABLE IF NOT EXISTS login  -- not used now too
-- (
--         id              SERIAL PRIMARY KEY,
--         sess_id         UUID,
--         user_id         INTEGER REFERENCES profile (uid),
--         user_agent      varchar(128),
--         add_time        TIMESTAMPTZ NOT NULL,
--         ip_addr         VARCHAR(24)
-- );



CREATE TYPE mark AS ENUM ('bad', 'ok', 'good');

CREATE TABLE IF NOT EXISTS call
(
        id              SERIAL PRIMARY KEY,
        caller          INTEGER REFERENCES profile (uid),
        answerer        INTEGER REFERENCES profile (uid),
        start_time      TIMESTAMP NOT NULL,
        end_time        TIMESTAMP,
        result          BOOLEAN     DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS call_mark
(
        id              SERIAL PRIMARY KEY,
        call_id         INTEGER REFERENCES call (id),
        emoted_user_id  INTEGER REFERENCES profile (uid),
        emotion         mark,
        caller          BOOLEAN NOT NULL
);

-- add it!
CREATE TABLE IF NOT EXISTS statistic
(
        id              SERIAL PRIMARY KEY,
        call_id         INTEGER REFERENCES call (id),
        ip              VARCHAR(12),
        country         VARCHAR,
        region          VARCHAR,
        city            VARCHAR
)