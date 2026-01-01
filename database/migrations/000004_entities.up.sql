CREATE TABLE university
(
    id          serial PRIMARY KEY,
    name        text,
    city        text,
    category    text,
    image_url   text,
    description text
);


CREATE TABLE professor
(
    id                serial PRIMARY KEY,
    name              text,
    education_history jsonb,
    image_url         text,
    description       text
);


CREATE TABLE lesson
(
    id          serial PRIMARY KEY,
    name        text,
    difficulty  int,
    description text
);


CREATE TABLE major
(
    id   serial PRIMARY KEY,
    name text
);



