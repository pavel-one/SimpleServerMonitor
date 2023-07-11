CREATE TABLE IF NOT EXISTS memory_data
(
    percent    INTEGER not null,
    free       INTEGER not null,
    total       INTEGER not null,
    created_at DATE DEFAULT (datetime('now'))
);