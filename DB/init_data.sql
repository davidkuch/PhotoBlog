USE photo_db;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS passwords;
DROP TABLE IF EXISTS galleries;
DROP TABLE IF EXISTS user_gallery_acces;
DROP TABLE IF EXISTS pic_to_gallery;

CREATE TABLE users(
    id VARCHAR(120) PRIMARY KEY,
    name VARCHAR(120),
    family VARCHAR(120),
    email VARCHAR(120)
);

CREATE TABLE passwords(
    id VARCHAR(120) PRIMARY KEY,
    password VARCHAR(120)
);

CREATE TABLE galleries(
    id VARCHAR(120) PRIMARY KEY,
    name VARCHAR(120),
    created DATE,
    last_edit DATE
);

CREATE TABLE user_gallery_acces(
    user_id VARCHAR(120),
    gallery_id VARCHAR(16),
    level INT
);

CREATE TABLE pic_to_gallery(
    gallery_id VARCHAR(120),
    pic_id VARCHAR(120)
);


INSERT INTO users(id, name, family) VALUES ('13', 'david', 'kuch');
INSERT INTO galleries(id,name,created,last_edit) VALUES ('11', 'gal1', '1990-10-11', '1990-10-11');