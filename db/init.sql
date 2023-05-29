USE languages_docker;

CREATE TABLE languages (
    id integer not null auto_increment,
    name varchar(200),
    PRIMARY KEY(id)
);

INSERT INTO languages (name) VALUES ('Java');
INSERT INTO languages (name) VALUES ('Kotlin');
INSERT INTO languages (name) VALUES ('Groovy');
INSERT INTO languages (name) VALUES ('Python');
INSERT INTO languages (name) VALUES ('Golang');