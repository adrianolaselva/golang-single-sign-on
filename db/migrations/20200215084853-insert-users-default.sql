
-- +migrate Up
INSERT INTO oauth_users
(id, name, last_name, username, email, password, birthday, activated, created_at, updated_at, deleted_at)
VALUES
('8d42ee3e-5717-4b65-b0b6-218361f981b3','Adriano','Moreira La Selva','adrianolaselva', 'adrianolaselva@gmail.com', '$2a$14$ZNQlwA.j3Mf7Gyq5GCHC9eZtyA6G9gJmEg3NtDf1MAR2R01JTb/Z.', '1987-02-11', 1, NOW(), NOW(), NULL);

-- +migrate Down
delete from oauth_user_roles;
delete from oauth_users;