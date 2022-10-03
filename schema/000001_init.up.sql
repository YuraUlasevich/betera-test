CREATE TABLE APOD(
    id int not NULL,
    imageurl VARCHAR(300),
    imageb bytea NOT NULL,
    flag int,
    PRIMARY KEY(id)
);