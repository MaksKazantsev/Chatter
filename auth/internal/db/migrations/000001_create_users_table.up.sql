CREATE TABLE IF NOT EXISTS users
(
    uuid       varchar(50) UNIQUE,
    username text,
    email    text UNIQUE,
    password text,
    refresh text
    );
CREATE TABLE IF NOT EXISTS user_profiles
(
    uuid varchar(50) UNIQUE,
    username text,
    email text UNIQUE,
    birthday text,
    bio text,
    lastonline time
);