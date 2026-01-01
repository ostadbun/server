-- 2 relation
CREATE TABLE lesson_pre_requisite
(
    lesson_id              int NOT NULL REFERENCES lesson (id) ON DELETE CASCADE,
    prerequisite_lesson_id int NOT NULL REFERENCES lesson (id) ON DELETE CASCADE,

    PRIMARY KEY (lesson_id, prerequisite_lesson_id),

    CHECK (lesson_id <> prerequisite_lesson_id)
);

-- 2 relation
CREATE TABLE lesson_co_requisite
(
    lesson_id              int not null references lesson (id) on delete cascade,
    co_requisite_lesson_id int not null references lesson (id) on delete cascade,

    PRIMARY KEY (lesson_id, co_requisite_lesson_id),
    CHECK (lesson_id <> co_requisite_lesson_id)
);


-- 2 relation
CREATE TABLE major_university
(
    university_id int NOT NULL references university (id) on delete cascade,
    major_id      int NOT NULL references major (id) on delete cascade,
    PRIMARY KEY (major_id, university_id)
);


-- 2 relation
CREATE TABLE professor_university
(

    university_id int NOT NULL references university (id) on delete cascade,
    professor_id  int NOT NULL references professor (id) on delete cascade,
    PRIMARY KEY (professor_id, university_id)
);


-- 2 relation
CREATE TABLE professor_lesson
(
    professor_id int NOT NULL references professor (id) on delete cascade,
    lesson_id    int NOT NULL references lesson (id) on delete cascade,
    PRIMARY KEY (lesson_id, professor_id)
);


-- 2 relation
CREATE TABLE professor_user
(

    professor_id int NOT NULL references professor (id) on delete cascade,
    user_id      int NOT NULL references users (id) on delete cascade,
    PRIMARY KEY (user_id, professor_id)
);


-- 2 relation
CREATE TABLE lesson_major
(

    lesson_id int NOT NULL references lesson (id) on delete cascade,
    major_id  int NOT NULL references major (id) on delete cascade,
    PRIMARY KEY (major_id, lesson_id)
);



CREATE TABLE passed_lesson_professor_user
(
    id            serial PRIMARY KEY,
    professor_id  int NOT NULL references professor (id),
    lesson_id     int NOT NULL references lesson (id),
    university_id int NOT NULL references university (id),
    major_id      int NOT NULL references major (id),
    user_id       int NOT NULL references users (id),

    UNIQUE (user_id, major_id, lesson_id, professor_id, university_id)
);