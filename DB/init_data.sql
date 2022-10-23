USE photo_db;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS galleries;

CREATE TABLE users(
    id INT PRIMARY KEY,
    name VARCHAR(120),
    family VARCHAR(120)
);

CREATE TABLE galleries(
    id INT PRIMARY KEY,
    name VARCHAR(120),
    created DATE,
    last_edit DATE,
    pics VARCHAR(500)
);


INSERT INTO users(id, name, family) VALUES ('1', 'david', 'kuch');
INSERT INTO galleries(id,name,created,last_edit,pics) VALUES ('11', 'gal1', '1990-10-11', '1990-10-11', '');