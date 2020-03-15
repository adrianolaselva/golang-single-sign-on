
-- +migrate Up
insert into oauth_clients (id, user_id, name, secret, scopes, redirect, revoked, created_at, updated_at)
values
('a9832dab-598c-11ea-a5a2-0242c0a8a000', '8d42ee3e-5717-4b65-b0b6-218361f981b3', 'Aplicativo 1', 'secret', 'user:read user:write user:delete', '', false, NOW(), NOW());

-- +migrate Down
delete from oauth_clients;
