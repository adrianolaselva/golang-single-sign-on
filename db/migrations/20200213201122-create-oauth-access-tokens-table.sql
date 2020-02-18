
-- +migrate Up
create table if not exists oauth_access_tokens (
  id varchar(36) not null,
  user_id varchar(256) default null,
  client_id varchar(255) not null,
  access_token varchar(1024) not null,
  scopes varchar(1024) null,
  revoked boolean default false,
  expires_at datetime not null,
  foreign key(client_id) references oauth_clients(id),
  foreign key(user_id) references oauth_users(id),
  primary key(id)
);

-- +migrate Down
drop table if exists oauth_access_tokens;
