CREATE TABLE university (
  id serial PRIMARY KEY,
  wiki JSONB
);


-- professor <-> university (many-to-many)
ALTER TABLE professor_university
ADD CONSTRAINT fk_professor FOREIGN KEY (professor_id) REFERENCES professor(id),
ADD CONSTRAINT fk_university FOREIGN KEY (university_id) REFERENCES university(id);

-- professor <-> student (many-to-many)
ALTER TABLE professor_student
ADD CONSTRAINT fk_professor_student_prof FOREIGN KEY (professor_id) REFERENCES professor(id),
ADD CONSTRAINT fk_professor_student_student FOREIGN KEY (student_id) REFERENCES student(id);

-- professor <-> lesson (many-to-many)
ALTER TABLE professor_lessons
ADD CONSTRAINT fk_professor_lessons_prof FOREIGN KEY (professor_id) REFERENCES professor(id),
ADD CONSTRAINT fk_professor_lessons_lesson FOREIGN KEY (lesson_id) REFERENCES lesson(id);

-- major <-> university (many-to-many)
ALTER TABLE major_university
ADD CONSTRAINT fk_major_univ_major FOREIGN KEY (major_id) REFERENCES major(id),
ADD CONSTRAINT fk_major_univ_univ FOREIGN KEY (university_id) REFERENCES university(id);

-- lesson <-> major (many-to-many)
ALTER TABLE lesson_major
ADD CONSTRAINT fk_lesson_major_lesson FOREIGN KEY (lesson_id) REFERENCES lesson(id),
ADD CONSTRAINT fk_lesson_major_major FOREIGN KEY (major_id) REFERENCES major(id);

-- passed lessons
ALTER TABLE passed_lesson_professor_student
ADD CONSTRAINT fk_passed_student FOREIGN KEY (student_id) REFERENCES student(id),
ADD CONSTRAINT fk_passed_professor FOREIGN KEY (professor_id) REFERENCES professor(id),
ADD CONSTRAINT fk_passed_lesson FOREIGN KEY (lesson_id) REFERENCES lesson(id),
ADD CONSTRAINT fk_passed_university FOREIGN KEY (university_id) REFERENCES university(id);

-- prerequisites
ALTER TABLE prerequisite
ADD CONSTRAINT fk_prereq_lesson FOREIGN KEY (lesson_id) REFERENCES lesson(id),
ADD CONSTRAINT fk_prereq_prereq FOREIGN KEY (prerequisite_id) REFERENCES lesson(id);

ALTER TABLE co_requisite
ADD CONSTRAINT fk_coreq_lesson FOREIGN KEY (lesson_id) REFERENCES lesson(id),
ADD CONSTRAINT fk_coreq_prereq FOREIGN KEY (prerequisite_id) REFERENCES lesson(id);

-- vote
ALTER TABLE vote
ADD CONSTRAINT fk_vote_student FOREIGN KEY (student_id) REFERENCES student(id),
ADD CONSTRAINT fk_vote_option FOREIGN KEY (option_id) REFERENCES options(id);

-- vote_snapshot
ALTER TABLE vote_snapshot
ADD CONSTRAINT fk_vote_snapshot_option FOREIGN KEY (option_id) REFERENCES options(id);
