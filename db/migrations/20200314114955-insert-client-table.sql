
-- +migrate Up
insert into oauth_clients (id, user_id, name, secret, scopes, redirect, revoked, created_at, updated_at)
values
('a9832dab-598c-11ea-a5a2-0242c0a8a000', '8d42ee3e-5717-4b65-b0b6-218361f981b3', 'Aplicativo 1', '8zd9ULma6xNN1wbR7h8er7z7qbERULsjCqD2pzT5', 'user:read,user:write,user:delete', 'https://webhook.site/365c600d-ce97-471a-805e-6076eef7f9aa', false, NOW(), NOW());
insert into oauth_clients (id, user_id, name, secret, scopes, redirect, revoked, created_at, updated_at)
values
('a9832dab-598c-11ea-a5a2-0242c0a8a001', '8d42ee3e-5717-4b65-b0b6-218361f981b3', 'Aplicativo 2', '8zd9ULma6xNN1wbR7h8er7z7qbERULsjCqD2pzT1', 'user:read,user:write,user:delete', 'http://localhost:4200/app/#/auth/callback', false, NOW(), NOW());
insert into oauth_clients (id, user_id, name, secret, scopes, redirect, revoked, created_at, updated_at)
values
('a9832dab-598c-11ea-a5a2-0242c0a8a002', '8d42ee3e-5717-4b65-b0b6-218361f981b3', 'Aplicativo 3', '8zd9ULma6xNN1wbR7h8er7z7qbERULsjCqD2pzT1', 'user:read,user:write,user:delete', 'http://localhost:9099/app/#/auth/callback', false, NOW(), NOW());

-- +migrate Down
delete from oauth_refresh_tokens;
delete from oauth_access_tokens;
delete from oauth_auth_codes;
delete from oauth_clients;
