ALTER TABLE activity_history
DROP
COLUMN id,
ADD CONSTRAINT unique_user_activity UNIQUE (user_id, activity_id);