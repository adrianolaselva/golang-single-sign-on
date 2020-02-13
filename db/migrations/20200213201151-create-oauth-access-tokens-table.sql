
-- +migrate Up
create table if not exists oauth_access_tokens (
  uuid varchar(60) not null,
  user_uuid varchar(60) not null,
  client_id varchar(256) default null,
  access_token BLOB,
  refresh_token varchar(256) default null,
  scopes varchar(2000) null,
  revoked boolean default false,
  expires_at datetime not null
);

-- +migrate Down
drop table if exists oauth_access_tokens;
