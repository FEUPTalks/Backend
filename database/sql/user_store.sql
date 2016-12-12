drop user if exists 'lesteamb'@'localhost';
drop user if exists 'lesteamb'@'%';
drop user if exists 'lesteamb'@'127.0.0.1';

create user 'lesteamb'@'localhost' identified by '99RedBalloons';
create user 'lesteamb'@'%' identified by '99RedBalloons';
create user 'lesteamb'@'127.0.0.1' identified by '99RedBalloons';

grant all privileges on talk_store.* to 'lesteamb'@'localhost' with grant option;
grant all privileges on talk_store.* to 'lesteamb'@'%' with grant option;
grant all privileges on talk_store.* to 'lesteamb'@'127.0.0.1' with grant option;
grant all privileges on user_store.* to 'lesteamb'@'localhost' with grant option;
grant all privileges on user_store.* to 'lesteamb'@'%' with grant option;
grant all privileges on user_store.* to 'lesteamb'@'127.0.0.1' with grant option;

drop database if exists user_store;
create database user_store character set utf8;

use user_store;

create table user (
    UUID char(144) not null primary key,
    Email varchar(50) not null,
    Name varchar(50) not null,
    HashCode char(60) not null,
    RoleValue tinyint unsigned default 3
) ENGINE=InnoDB;

alter table User add unique(email);

insert into user (UUID, Email, Name, HashCode, RoleValue)
values (
    '1a1669b5-e83b-4727-95c1-5d42dc8f46c7',
    'admin@teste.com',
    'Admin',
    '$2a$10$2ajWauyJ5.SEfA1GL/wZHu9JwfORBQ5vsln4FLpt4iqwrIypb7EKK',
    '1'
);

insert into user (UUID, Email, Name, HashCode, RoleValue)
values (
    '32f116d8-3bf3-4f91-9dc8-2b85dc0be5fd',
    'employee1@teste.com',
    'Employee1',
    '$2a$10$i/otNHwJlqwi8IFKDYXyi.nREdcrV.XK.jTghunimu0zEyzgCv7vy',
    '2'
);

insert into user (UUID, Email, Name, HashCode, RoleValue)
values (
    'a316481d-122a-4a06-bf03-9244dfdcf433',
    'employee2@teste.com',
    'Employee2',
    '$2a$10$bzFWivaFmaHor/aK.0XPF.oMpeCFSvg6qjQIGGosPQqMBwbqDrYEe',
    '2'
);