DELETE FROM chips;
VACUUM;
DELETE FROM sensors;
VACUUM;
DELETE FROM sensors_data;
VACUUM;

INSERT INTO chips (id, name) VALUES (1, 'TestChip');

INSERT INTO sensors (name, high_temp, crit_temp, chip_id)
VALUES ('test', 60.2, 100.2, 1),
       ('test1', 61.2, 101.2, 1),
       ('test2', 62.2, 102.2, 1);

INSERT INTO sensors_data (temp, sensor_id)
VALUES (10, (SELECT id FROM sensors WHERE name = 'test' LIMIT 1));

INSERT INTO sensors_data (temp, sensor_id)
VALUES (25, (SELECT id FROM sensors WHERE name = 'test' LIMIT 1));

INSERT INTO sensors_data (temp, sensor_id)
VALUES (28.2, (SELECT id FROM sensors WHERE name = 'test' LIMIT 1));

INSERT INTO sensors_data (temp, sensor_id)
VALUES (28.3, (SELECT id FROM sensors WHERE name = 'test1' LIMIT 1));

INSERT INTO sensors_data (temp, sensor_id)
VALUES (15, (SELECT id FROM sensors WHERE name = 'test1' LIMIT 1));