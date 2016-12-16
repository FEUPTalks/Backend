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

drop database if exists talk_store;
create database talk_store character set utf8;

use talk_store;

create table picture (
    pictureID int unsigned not null auto_increment primary key,
    speakerpicture longtext not null
) ENGINE=InnoDB;

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
    SpeakerPicture longtext,
    HostName varchar(50) not null,
    HostEmail varchar(50) not null,
    Snack tinyint not null,
    Room varchar(10) not null,
    Other varchar(1000) not null,
    State tinyint unsigned default 1
) ENGINE=InnoDB;

create table talkRegistration (
  Email varchar(50) not null,
  TalkID int unsigned not null,
  primary key(TalkID, Email),
  Name varchar(255) not null,
  IsAttendingSnack boolean,
  WantsToReceiveNotifications boolean
) ENGINE=InnoDB;

create table temporaryTalkRegistration (
  Email varchar(50) not null,
  TalkID int unsigned not null,
  primary key(TalkID, Email),
  Name varchar(255) not null,
  IsAttendingSnack boolean,
  WantsToReceiveNotifications boolean
) ENGINE=InnoDB;

create table talkRegistrationLog (
  LogID int unsigned not null auto_increment primary key,
  Name varchar(255) not null,
  Email varchar(50) not null,
  TalkID int unsigned not null,
  IsAttendingSnack boolean,
  WantsToReceiveNotifications boolean,
  TransactionType tinyint unsigned default 0,
  TransactionDate datetime
) ENGINE=InnoDB;

create table user (
    UserID int unsigned not null auto_increment primary key,
    Email varchar(50) not null,
    Name varchar(50) not null,
    HashCode varchar(256) not null,
    RoleValue tinyint unsigned default 3
) ENGINE=InnoDB;

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

insert into talk (Title, Summary, Date, DateFlex,
Duration, ProponentName, ProponentEmail, SpeakerName, SpeakerBrief, SpeakerAffiliation,
SpeakerPicture, HostName, HostEmail, Snack, Room, Other, State)
values (
    'Test3',
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
    'Outros que tais',
    '5'
);

insert into talk (Title, Summary, Date, DateFlex,
Duration, ProponentName, ProponentEmail, SpeakerName, SpeakerBrief, SpeakerAffiliation,
SpeakerPicture, HostName, HostEmail, Snack, Room,Other, State)
values (
    'Test4',
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
    'Outros que tais',
    '5'
);

insert into user (Email, Name, HashCode, RoleValue)
values (
    'teste@teste.com',
    'Teste Teste',
    '123456789abcdef',
    '3'
);

insert into talkRegistration (Email, TalkID, Name, IsAttendingSnack, WantsToReceiveNotifications)
values (
    'bob_d_girl@hotmale.com',
    '1',
    'Bob Faget',
    '1',
    '1'
);

insert into talkRegistration (Email, TalkID, Name, IsAttendingSnack, WantsToReceiveNotifications)
values (
    'stefania_d_guy@hotmale.com',
    '1',
    'Stefania Dud',
    '0',
    '0'
);

insert into talkRegistration (Email, TalkID, Name, IsAttendingSnack, WantsToReceiveNotifications)
values (
    'dark_snipz@2gud4u.com',
    '2',
    'Dark Snipz',
    '1',
    '1'
);

insert into talkRegistration (Email, TalkID, Name, IsAttendingSnack, WantsToReceiveNotifications)
values (
    'johndoe@default.com',
    '1',
    'John Doe',
    '1',
    '0'
);
