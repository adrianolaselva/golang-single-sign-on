
-- +migrate Up
INSERT INTO oauth_roles (id, name) values ('a9832dab-598c-11ea-a5a2-0242c0a8a000', 'ADMINISTRATOR');
INSERT INTO oauth_roles (id, name) values ('a9832dab-598c-11ea-a5a2-0242c0a8a001', 'CUSTOMER');

-- +migrate Down
DELETE FROM oauth_roles;
