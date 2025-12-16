
CREATE TYPE target_table AS ENUM ('univ', 'prof');

CREATE TABLE student (
  id SERIAL PRIMARY KEY,
  wiki JSONB,
  admin_by_id INT NULL,
  university_id INT
);

CREATE TABLE professor (
  id SERIAL PRIMARY KEY,
  name TEXT,
  start_year INT,
  education_history TEXT
);


ALTER TABLE lesson ALTER COLUMN wiki TYPE JSONB USING wiki::JSONB;

CREATE TABLE prerequisite (
  id SERIAL PRIMARY KEY,
  wiki JSONB,
  prerequisite_id INT REFERENCES lesson(id),
  lesson_id INT REFERENCES lesson(id)
);

CREATE TABLE co_requisite (
  id SERIAL PRIMARY KEY,
  wiki JSONB,
  prerequisite_id INT REFERENCES lesson(id),
  lesson_id INT REFERENCES lesson(id)
);

CREATE TABLE major (
  id SERIAL PRIMARY KEY,
  wiki JSONB
);

CREATE TABLE major_university (
  id SERIAL PRIMARY KEY,
  university_id INT,
  major_id INT REFERENCES major(id)
);

CREATE TABLE professor_university (
  id SERIAL PRIMARY KEY,
  professor_id INT REFERENCES professor(id),
  university_id INT
);

CREATE TABLE professor_lessons (
  id SERIAL PRIMARY KEY,
  professor_id INT REFERENCES professor(id),
  lesson_id INT REFERENCES lesson(id)
);

CREATE TABLE professor_student (
  id SERIAL PRIMARY KEY,
  professor_id INT REFERENCES professor(id),
  student_id INT REFERENCES student(id)
);

CREATE TABLE passed_lesson_professor_student (
  id SERIAL PRIMARY KEY,
  student_id INT REFERENCES student(id),
  professor_id INT REFERENCES professor(id),
  lesson_id INT REFERENCES lesson(id),
  semester SMALLINT,
  university_id INT
);

CREATE TABLE options (
  id SERIAL PRIMARY KEY,
  name TEXT,
  weight SMALLINT,
  target_table target_table,
  target_id INT
);

CREATE TABLE vote (
  id SERIAL PRIMARY KEY,
  student_id INT REFERENCES student(id),
  option_id INT REFERENCES options(id),
  rate SMALLINT,
  time INT
);

CREATE TABLE vote_snapshot (
  id SERIAL PRIMARY KEY,
  student_sum INT,
  vote_sum INT,
  option_id INT REFERENCES options(id),
  rate_at_this_time SMALLINT,
  time INT
);

CREATE TABLE lesson_major (
  id SERIAL PRIMARY KEY,
  major_id INT REFERENCES major(id),
  lesson_id INT REFERENCES lesson(id),
  wiki JSONB
);

CREATE TABLE student_activity_history (
  id SERIAL PRIMARY KEY,
  student_id INT REFERENCES student(id),
  activity_id INT,
  time INT
);

CREATE TABLE activity_value (
  id SERIAL PRIMARY KEY,
  name TEXT,
  wiki JSONB,
  value INT
);
