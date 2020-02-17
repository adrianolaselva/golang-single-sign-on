
-- +migrate Up
INSERT INTO users (uuid, username, email, password, birthday, activated, created_at, updated_at, deleted_at) VALUES('8d42ee3e-5717-4b65-b0b6-218361f981b3', 'adrianolaselva', 'adrianolaselva@gmail.com', '$2a$14$ZNQlwA.j3Mf7Gyq5GCHC9eZtyA6G9gJmEg3NtDf1MAR2R01JTb/Z.', '1999-12-31', 1, '2020-02-15 11:47:00.0', '2020-02-15 11:47:00.0', NULL);

-- +migrate Down
delete from users;