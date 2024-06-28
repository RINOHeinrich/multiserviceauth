CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    login VARCHAR(250) UNIQUE,
    password VARCHAR(250)
);