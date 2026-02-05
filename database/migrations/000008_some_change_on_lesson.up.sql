ALTER TABLE lesson
    ADD COLUMN name_english TEXT,
    ADD COLUMN description_english TEXT,
    ADD CONSTRAINT unique_name_english UNIQUE (name_english),
    ADD CONSTRAINT unique_name UNIQUE (name);