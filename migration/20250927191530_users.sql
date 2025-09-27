-- +goose Up
CREATE TABLE student (
  id serial PRIMARY KEY,
  name text,
  family text,
  username text,
  permission smallint,
  semester smallint,
  average_degree numeric(4,2),
  majorID int,
  universityID int
);

-- +goose Down
DROP TABLE student;
