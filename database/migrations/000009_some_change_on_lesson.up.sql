ALTER TABLE lesson
    ADD COLUMN is_released boolean DEFAULT false,
    ADD COLUMN registered_by int REFERENCES users(id),
    ADD COLUMN add_date timestamp DEFAULT CURRENT_TIMESTAMP;