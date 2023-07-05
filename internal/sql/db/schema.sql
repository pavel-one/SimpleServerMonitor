CREATE TABLE IF NOT EXISTS chips
(
    id   integer
        constraint chips_pk
            primary key,
    name varchar not null
);

CREATE TABLE IF NOT EXISTS sensors
(
    id        integer
        constraint sensors_pk
            primary key,
    name      varchar       not null,
    high_temp DECIMAL(3, 2) not null,
    crit_temp DECIMAL(3, 2) not null,
    chip_id integer,
    FOREIGN KEY (chip_id) REFERENCES chips (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sensors_data
(
    temp       DECIMAL(3, 2) not null,
    sensor_id  integer       not null,
    created_at DATE DEFAULT (datetime('now', 'localtime')),
    FOREIGN KEY (sensor_id) REFERENCES sensors (id) ON DELETE CASCADE
);