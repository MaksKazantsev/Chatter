CREATE TABLE IF NOT EXISTS users
(
    uuid       varchar(50) UNIQUE,
    username text,
    email    text UNIQUE,
    password text
    );