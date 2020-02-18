
-- +migrate Up
create table if not exists oauth_user_roles (
    user_id varchar(60) not null,
    role_id varchar(50) not null,
    foreign key(user_id) references oauth_users(id),
    foreign key(role_id) references oauth_roles(id),
    primary key(user_id, role_id)
);

-- +migrate Down
drop table if exists oauth_user_roles;