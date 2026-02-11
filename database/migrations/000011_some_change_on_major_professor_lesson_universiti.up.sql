-- ============ بخش ۱: افزودن ستون‌ها (registered_by به‌صورت NULLABLE) ============
ALTER TABLE lesson
    ADD COLUMN IF NOT EXISTS name_english TEXT,
    ADD COLUMN IF NOT EXISTS description_english TEXT,
    ADD COLUMN IF NOT EXISTS add_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,      -- ✅ از اول NOT NULL
    ADD COLUMN IF NOT EXISTS update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,   -- ✅ از اول NOT NULL
    ADD COLUMN IF NOT EXISTS registered_by INTEGER REFERENCES users(id) ON DELETE CASCADE; -- ⚠️ NULLABLE


ALTER TABLE university
    ADD COLUMN IF NOT EXISTS name_english TEXT,
    ADD COLUMN IF NOT EXISTS description_english TEXT,
    ADD COLUMN IF NOT EXISTS add_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ADD COLUMN IF NOT EXISTS update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ADD COLUMN IF NOT EXISTS registered_by INTEGER REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE professor
    ADD COLUMN IF NOT EXISTS name_english TEXT,
    ADD COLUMN IF NOT EXISTS description_english TEXT,
    ADD COLUMN IF NOT EXISTS add_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ADD COLUMN IF NOT EXISTS update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ADD COLUMN IF NOT EXISTS registered_by INTEGER REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE major
    ADD COLUMN IF NOT EXISTS name_english TEXT,
    ADD COLUMN IF NOT EXISTS description_english TEXT,
    ADD COLUMN IF NOT EXISTS add_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ADD COLUMN IF NOT EXISTS update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ADD COLUMN IF NOT EXISTS registered_by INTEGER REFERENCES users(id) ON DELETE CASCADE;

-- ============ بخش ۲: به‌روزرسانی داده‌های قدیمی (فقط تاریخ‌ها — بدون نیاز به کاربر) ============
-- تاریخ‌ها از اول با مقدار پیش‌فرض پر شده‌اند، اما برای اطمینان:
UPDATE lesson SET update_date = NOW() WHERE update_date IS NULL;
UPDATE university SET update_date = NOW() WHERE update_date IS NULL;
UPDATE professor SET update_date = NOW() WHERE update_date IS NULL;
UPDATE major SET update_date = NOW() WHERE update_date IS NULL;

-- ============ بخش ۳: محدودیت‌های یکتا ============
-- حذف محدودیت‌های احتمالی قبلی
ALTER TABLE lesson DROP CONSTRAINT IF EXISTS unique_lesson_name;
ALTER TABLE lesson DROP CONSTRAINT IF EXISTS unique_lesson_name_english;
ALTER TABLE university DROP CONSTRAINT IF EXISTS unique_university_name;
ALTER TABLE university DROP CONSTRAINT IF EXISTS unique_university_name_english;
ALTER TABLE major DROP CONSTRAINT IF EXISTS unique_major_name;
ALTER TABLE major DROP CONSTRAINT IF EXISTS unique_major_name_english;
ALTER TABLE professor DROP CONSTRAINT IF EXISTS unique_professor_name_english;

-- افزودن محدودیت‌های جدید
ALTER TABLE lesson
    ADD CONSTRAINT unique_lesson_name UNIQUE (name),
    ADD CONSTRAINT unique_lesson_name_english UNIQUE (name_english);

ALTER TABLE university
    ADD CONSTRAINT unique_university_name UNIQUE (name),
    ADD CONSTRAINT unique_university_name_english UNIQUE (name_english);

ALTER TABLE major
    ADD CONSTRAINT unique_major_name UNIQUE (name),
    ADD CONSTRAINT unique_major_name_english UNIQUE (name_english);
