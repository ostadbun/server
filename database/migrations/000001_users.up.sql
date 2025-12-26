CREATE TABLE users
(
    id            serial PRIMARY KEY,
    name          text,
    email         text UNIQUE,
    image_url     text[],
    admin_by      int,
    university_id int,
    major_id      int

);

CREATE INDEX idx_users_email ON users (email);