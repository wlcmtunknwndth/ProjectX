CREATE TABLE public.auth (
                             username character varying(64) NOT NULL,
                             password character varying(64),
                             isadmin boolean DEFAULT false,
                             gender boolean DEFAULT false,
                             age SMALLINT CHECK (age > 0 and age < 130)
);

ALTER TABLE public.auth OWNER TO postgres;


INSERT INTO Auth(username, password, isadmin, gender, age) VALUES ('idkidkidk', 'idkidkidk', true, true, 20);
INSERT INTO Auth(username, password, isAdmin, gender, age) VALUES ('idkidk', 'idkidk', false, false, 18);

CREATE TABLE public.events(
    id BIGSERIAL CHECK (id > 0) PRIMARY KEY,
    price BIGINT CHECK(price > 0 and price < 100000000),
    restrictions BIGINT CHECK ( restrictions > 0 AND restrictions < 120),
    date timestamptz,
    city VARCHAR(32) NOT NULL,
    address VARCHAR(128) NOT NULL,
    name VARCHAR(128) NOT NULL,
    img_path VARCHAR(256),
    description VARCHAR(2048)
);

CREATE TABLE public.features(
    id BIGSERIAL CHECK (id > 0) PRIMARY KEY,
    Name VARCHAR(128) UNIQUE
    Tag VARCHAR(64) UNIQUE
)

CREATE TABLE public.index(
    id BIGSERIAL PRIMARY KEY,
    event_id BIGINT CHECK (event_id > 0)
    features BIGINT[]
)
CREATE TABLE public.cache(
    id BIGINT CHECK (id > 0) PRIMARY KEY
)