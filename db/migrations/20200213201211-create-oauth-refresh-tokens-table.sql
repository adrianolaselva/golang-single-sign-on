
-- +migrate Up
create table if not exists oauth_refresh_tokens (
  uuid varchar(256) default null,
  access_token BLOB,
  authentication BLOB,
  revoked boolean default false,
  expires_at datetime not null
);

-- +migrate Down
drop table if exists oauth_refresh_tokens;
