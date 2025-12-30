ALTER TABLE users
    ADD CONSTRAINT fk_users_admin_by
        FOREIGN KEY (admin_by)
            REFERENCES users (id)
            ON DELETE SET NULL;