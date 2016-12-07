drop user if exists 'lesteamb'@'localhost';
drop user if exists 'lesteamb'@'%';
drop user if exists 'lesteamb'@'127.0.0.1';

create user 'lesteamb'@'localhost' identified by '99RedBalloons';
create user 'lesteamb'@'%' identified by '99RedBalloons';
create user 'lesteamb'@'127.0.0.1' identified by '99RedBalloons';

grant all privileges on talk_store.* to 'lesteamb'@'localhost' with grant option;
grant all privileges on talk_store.* to 'lesteamb'@'%' with grant option;
grant all privileges on talk_store.* to 'lesteamb'@'127.0.0.1' with grant option;

drop database if exists talk_store;
create database talk_store;

use talk_store;

create table picture (
    PictureID int unsigned not null auto_increment primary key,
    filepath varchar(200) not null
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table talk (
    TalkID int unsigned not null auto_increment primary key,
    Title varchar(50) not null,
    Summary varchar(500) not null,
    Date datetime not null,
    DateFlex int not null,
    Duration tinyint unsigned not null,
    ProponentName varchar(500) not null,
    ProponentEmail varchar(500) not null,
    SpeakerName varchar(50) not null,
    SpeakerBrief varchar(50) not null,
    SpeakerAffiliation varchar(50) not null,
    SpeakerPicture int unsigned not null,
    HostName varchar(50) not null,
    HostEmail varchar(50) not null,
    Snack tinyint not null,
    Room varchar(10) not null,
    Other varchar(1000) not null,
    State tinyint unsigned default 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table talkRegistration (
  Email varchar(50) not null,
  TalkID int unsigned not null,
  primary key(TalkID, Email),
  Name varchar(255) not null,
  IsAttendingSnack boolean
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table user (
    UserID int unsigned not null auto_increment primary key,
    Email varchar(50) not null,
    Name varchar(50) not null,
    HashCode varchar(256) not null,
    RoleValue tinyint unsigned default 3
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

alter table talk
add foreign key (SpeakerPicture)
references picture(PictureID);

alter table talk
add constraint chk_proposedDates check (datediff(ProposedEndDate, ProposedInitialDate) >= 0);

alter table talkRegistration
add foreign key (TalkID)
references talk(TalkID);

insert into picture (filepath)
values (
    'test'
);

insert into talk (Title, Summary, Date, DateFlex,
Duration, ProponentName, ProponentEmail, SpeakerName, SpeakerBrief, SpeakerAffiliation,
SpeakerPicture, HostName, HostEmail, Snack, Room,Other)
values (
    'Test',
    'We are testing the talk proposal functionality',
    '2016-11-07 00:00:00',
    '5',
    '3',
    'proponent',
    'proponent@email.com',
    'speaker',
    'É um ganda gajo',
    'harvard',
    '1',
    'host@email.com',
    'host@email.com',
    'Rissóis, panados, aguá e sumos naturais',
    'B219',
    'Outros que tais'
);

insert into talk (Title, Summary, Date, DateFlex,
Duration, ProponentName, ProponentEmail, SpeakerName, SpeakerBrief, SpeakerAffiliation,
SpeakerPicture, HostName, HostEmail, Snack, Room,Other)
values (
    'Test2',
    'We are testing the talk proposal functionality',
    '2016-11-07 00:00:00',
    '5',
    '3',
    'proponent2',
    'proponent2@email.com',
    'speaker2',
    'É um ganda gajo',
    'harvard',
    '1',
    'host2@email.com',
    'host2@email.com',
    'Rissóis, panados, aguá e sumos naturais',
    'B219',
    'Outros que tais'
);

insert into user (Email, Name, HashCode, RoleValue)
values (
    'teste@teste.com',
    'Teste Teste',
    '123456789abcdef',
    '3'
);

insert into talkRegistration (Email, TalkID, Name, IsAttendingSnack)
values (
    'bob_d_girl@hotmale.com',
    '1',
    'Bob Faget',
    '1'
);

insert into talkRegistration (Email, TalkID, Name, IsAttendingSnack)
values (
    'stefania_d_guy@hotmale.com',
    '1',
    'Stefania Dud',
    '0'
);

insert into talkRegistration (Email, TalkID, Name, IsAttendingSnack)
values (
    'dark_snipz@2gud4u.com',
    '2',
    'Dark Snipz',
    '1'
);

insert into talkRegistration (Email, TalkID, Name, IsAttendingSnack)
values (
    'johndoe@default.com',
    '1',
    'John Doe',
    '1'
);
