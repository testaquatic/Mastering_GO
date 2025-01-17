DROP DATABASE IF EXISTS restapi;

CREATE DATABASE restapi;

\ c restapi;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    PASSWORD VARCHAR NOT NULL,
    lastlogin INT,
    admin INT,
    active INT
);

INSERT INTO users (username, PASSWORD, lastlogin, admin, active)
VALUES ('admin', 'admin', 0, 1, 1);