CREATE TABLE IF NOT EXISTS users
(
    uuid       varchar(50) UNIQUE,
    username text,
    email    text UNIQUE,
    password text,
    refresh text,
    isverified bool,
    joined time
    );
CREATE TABLE IF NOT EXISTS user_profiles
(
    uuid varchar(50) UNIQUE,
    avatar text,
    username text,
    firstname text,
    secondname text,
    email text UNIQUE,
    birthday text,
    bio text,
    lastonline time
);
CREATE TABLE IF NOT EXISTS codes
(
    code text,
    email text UNIQUE,
    isverified bool
);
CREATE TABLE IF NOT EXISTS friend_reqs
(
    sender text,
    receiver text
);