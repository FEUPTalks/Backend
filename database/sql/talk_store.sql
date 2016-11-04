drop user 'lesteamb'@'localhost';
drop user 'lesteamb'@'127.0.0.1';
drop user 'lesteamb'@'::1';

create user 'lesteamb'@'localhost' identified by '99RedBalloons';
grant all privileges on *.* to 'lesteamb'@'localhost' with grant option;
create user 'lesteamb'@'127.0.0.1' identified by '99RedBalloons';
grant all privileges on *.* to 'lesteamb'@'127.0.0.1' with grant option;
create user 'lesteamb'@'::1' identified by '99RedBalloons';
grant all privileges on *.* to 'lesteamb'@'::1' with grant option;

drop database talk_store;
create database talk_store;

use talk_store;

create table talk (
    TalkID int unsigned not null auto_increment primary key,
    Title varchar(50) default null,
    Summary varchar(500) default null,
    ProposedInitialDate datetime default null,
    ProposedEndDate datetime default null,
    DefinitiveDate datetime default null,
    Duration tinyint unsigned default null,
    ProponentName varchar(500) default null,
    ProponentEmail varchar(500) default null,
    ProponentAffiliation varchar(50) default null,
    SpeakerName varchar(50) default null,
    SpeakerBrief varchar(50) default null,
    SpeakerAffiliation varchar(50) default null,
    HostName varchar(50) default null,
    HostEmail varchar(50) default null,
    Snack varchar(255) default null,
    Room varchar(10)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

alter table talk
add constraint chk_proposedDates check (datediff(ProposedEndDate, ProposedInitialDate) >= 0);

insert into talk (Title, Summary, ProposedInitialDate, ProposedEndDate, DefinitiveDate,
Duration, ProponentName, ProponentEmail, ProponentAffiliation, SpeakerName, SpeakerBrief, SpeakerAffiliation,
HostName, HostEmail, Snack, Room)
values (
    'Test',
    'We are testing the talk proposal functionality',
    '2016-11-07T00:00:00Z',
    '2016-11-11T00:00:00Z',
    '2016-11-10T12:00:00Z',
    '3600000000000',
    'proponent',
    'proponent@email.com',
    'feup',
    'speaker',
    'É um ganda gajo',
    'harvard',
    'host@email.com',
    'host@email.com',
    'Rissóis, panados, aguá e sumos naturais',
    'B219'
);