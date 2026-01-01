-- voting

CREATE TYPE target_entity_type AS ENUM ('university', 'professor');

CREATE TABLE option
(
    id          serial PRIMARY KEY,
    name        text,
    weight      int,
    target_id   int                NOT NULL,
    target_type target_entity_type NOT NULL
);


CREATE TABLE vote
(
    user_id   int       NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    option_id int       NOT NULL REFERENCES option (id) ON DELETE CASCADE,
    rate      int       NOT NULL CHECK (rate >= 1 AND rate <= 10),
    rate_time timestamp NOT NULL DEFAULT current_timestamp,

    PRIMARY KEY (user_id, option_id)
);


CREATE TABLE vote_snapshot
(
    id            serial PRIMARY KEY,
    student_count int       not null,
    vote_count    int       NOT NULL,
    option_id     int       NOT NULL,
    rating        smallint  not null,
    time          timestamp not null default current_timestamp
);
CREATE INDEX idx_vote_snapshot_time ON vote_snapshot (time DESC);


-- activities


CREATE TABLE activity
(
    id    serial PRIMARY KEY,
    name  text not null,
    value int  NOT NULL
);

CREATE TABLE activity_history
(
    id          serial PRIMARY KEY,
    user_id     int       not null references users (id) on delete cascade,
    activity_id int       NOT NULL references activity (id) on DELETE cascade,
    time        timestamp not null default current_timestamp
);

CREATE INDEX idx_activity_history_user ON activity_history (user_id, time DESC);