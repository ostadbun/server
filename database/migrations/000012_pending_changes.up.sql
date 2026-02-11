-- pending_lesson
ALTER TABLE pending_lesson
    ADD COLUMN IF NOT EXISTS name_english TEXT,
    ADD COLUMN IF NOT EXISTS description_english TEXT;

-- pending_university
ALTER TABLE pending_university
    ADD COLUMN IF NOT EXISTS name_english TEXT,
    ADD COLUMN IF NOT EXISTS description_english TEXT;

-- pending_professor
ALTER TABLE pending_professor
    ADD COLUMN IF NOT EXISTS name_english TEXT,
    ADD COLUMN IF NOT EXISTS description_english TEXT;

-- major (جدول اصلی)
ALTER TABLE major
    ADD COLUMN IF NOT EXISTS description_english TEXT;

-- pending_major
ALTER TABLE pending_major
    ADD COLUMN IF NOT EXISTS description TEXT,
    ADD COLUMN IF NOT EXISTS name_english TEXT,
    ADD COLUMN IF NOT EXISTS description_english TEXT;