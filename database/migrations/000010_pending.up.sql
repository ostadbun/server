CREATE TABLE pending_university
(
    id               serial PRIMARY KEY,
    name             text    NOT NULL,
    city             text    NOT NULL,
    category         text    NOT NULL,
    image_url        text,
    description      text,
    status           text    NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'approved', 'rejected')),
    submitted_by     integer NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    submitted_at     timestamp        DEFAULT NOW(),
    approved_by      integer REFERENCES users (id) ON DELETE SET NULL,
    approved_at      timestamp,
    rejection_reason text
);

-- درخواست‌های تعلیقی استاد
CREATE TABLE pending_professor
(
    id                serial PRIMARY KEY,
    name              text    NOT NULL,
    education_history jsonb   NOT NULL,
    image_url         text,
    description       text,
    status            text    NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'approved', 'rejected')),
    submitted_by      integer NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    submitted_at      timestamp        DEFAULT NOW(),
    approved_by       integer REFERENCES users (id) ON DELETE SET NULL,
    approved_at       timestamp,
    rejection_reason  text
);

-- درخواست‌های تعلیقی درس
CREATE TABLE pending_lesson
(
    id               serial PRIMARY KEY,
    name             text    NOT NULL,
    difficulty       int     NOT NULL CHECK (difficulty BETWEEN 1 AND 5),
    description      text,
    status           text    NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'approved', 'rejected')),
    submitted_by     integer NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    submitted_at     timestamp        DEFAULT NOW(),
    approved_by      integer REFERENCES users (id) ON DELETE SET NULL,
    approved_at      timestamp,
    rejection_reason text
);

-- درخواست‌های تعلیقی رشته
CREATE TABLE pending_major
(
    id               serial PRIMARY KEY,
    name             text    NOT NULL,
    status           text    NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'approved', 'rejected')),
    submitted_by     integer NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    submitted_at     timestamp        DEFAULT NOW(),
    approved_by      integer REFERENCES users (id) ON DELETE SET NULL,
    approved_at      timestamp,
    rejection_reason text
);