
-- +migrate Up
create table if not exists oauth_refresh_tokens (
  id varchar(36) not null,
  access_token_id varchar(36) not null,
  refresh_token varchar(1024) not null,
  revoked boolean default false,
  created_at datetime not null,
  expires_at datetime not null,
  foreign key(access_token_id) references oauth_access_tokens(id),
  primary key(id)
);

-- +migrate Down
drop table if exists oauth_refresh_tokens;
