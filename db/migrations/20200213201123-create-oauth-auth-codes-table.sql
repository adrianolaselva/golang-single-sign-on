
-- +migrate Up
create table if not exists oauth_auth_codes (
  uuid varchar(60) not null,
  user_uuid varchar(60) not null,
  client_id varchar(60) not null,
  scopes varchar(2000) null,
  revoked boolean default false,
  expires_at datetime not null
);

-- +migrate Down
drop table if exists oauth_auth_codes;
