
-- +migrate Up
create table if not exists oauth_auth_codes (
  id varchar(60) not null,
  user_id varchar(60) not null,
  client_id varchar(36) not null,
  code varchar(1024) not null,
  scopes varchar(1024) null,
  revoked boolean default false,
  created_at datetime not null,
  expires_at datetime not null,
  foreign key(client_id) references oauth_clients(id),
  foreign key(user_id) references oauth_users(id),
  primary key(id)
);

-- +migrate Down
drop table if exists oauth_auth_codes;
