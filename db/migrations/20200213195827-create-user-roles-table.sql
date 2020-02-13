
-- +migrate Up
create table if not exists user_roles (
    user_uuid varchar(60) not null,
    role varchar(50) not null,
    foreign key(user_uuid) references users(uuid),
    foreign key(role) references roles(name),
    primary key(user_uuid, role)
);

-- +migrate Down
drop table if exists user_roles;