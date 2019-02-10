create schema public;

comment on schema public is 'standard public schema';

alter schema public owner to postgres;

create table games
(
	game_id char(12) not null
		constraint games_pk
			unique,
	visteam char(3) not null,
	hometeam char(3) not null,
	game_date date not null,
	number integer default 0 not null,
	starttime time not null,
	daynight varchar(16) not null,
	usedh boolean default false not null,
	pitches varchar(7) not null,
	umphome char(8),
	ump1b char(8),
	ump2b char(8),
	ump3b char(8),
	umplf char(8),
	umprf char(8),
	fieldcond varchar(7) not null,
	precip varchar(7) not null,
	sky varchar(7) not null,
	temp integer default 0 not null,
	winddir varchar(7) not null,
	windspeed integer default 0 not null,
	timeofgame integer default 0 not null,
	attendance integer default 0 not null,
	site char(5) not null,
	wp char(8) not null,
	lp char(8) not null,
	save char(8),
	gwrbi char(8),
	edittime timestamp,
	howscored varchar(32),
	inputprogvers varchar(32),
	inputter varchar(32),
	inputtime timestamp,
	scorer varchar(32),
	translator varchar(32)
);

comment on table games is 'Contains the basic information about a game, taken from the INFO tags for the game event';

comment on column games.game_id is '12 character unique identifier for each game see retrosheet.org event file description for an indepth overview of the game_id';

comment on column games.number is 'which game in the day this entry represents, 0 is first 1 is second, 2 is third, etc';

comment on column games.gwrbi is 'game winning RBI';

alter table games owner to baseball;

create unique index games_game_id_uindex
	on games (game_id);

create table starters
(
	game_id char(12) not null
		constraint starters_games_game_id_fk
			references games (game_id)
				on delete cascade,
	player_id char(8) not null,
	name varchar(64) not null,
	team integer not null,
	batting_order integer not null,
	position integer not null,
	constraint starters_pk
		primary key (game_id, player_id)
);

comment on table starters is 'the starting lineups for the game';

comment on column starters.name is 'player''s name';

comment on column starters.team is '0 for visiting team, 1 for home team';

comment on column starters.batting_order is 'when the dh is used the pitcher is given the batting order position 0';

comment on column starters.position is 'DH has the position value of 10';

alter table starters owner to baseball;

create table plays
(
	game_id char(12) not null
		constraint plays_games_game_id_fk
			references games (game_id)
				on delete cascade,
	idx integer not null,
	inning integer not null,
	team integer not null,
	player_id char(8) not null,
	count varchar(16) not null,
	pitches varchar(32) not null,
	event varchar(32) not null,
	constraint plays_pk
		primary key (game_id, idx)
);

comment on table plays is 'contains the plays found for a particular game';

comment on column plays.idx is 'the index into the plays in the game';

comment on column plays.team is '0 for visiting team, 1 for home team';

comment on column plays.count is 'the count at the time the play occurred, ?? is used when the count is unknown';

comment on column plays.pitches is 'the pitches thrown to the batter, must games will not have this data so this column will be empty';

comment on column plays.event is 'the result of the play';

alter table plays owner to baseball;

create table coms
(
	game_id char(12) not null
		constraint coms_games_game_id_fk
			references games (game_id)
				on delete cascade,
	idx integer not null,
	description varchar(128) not null,
	constraint coms_pk
		primary key (game_id, idx)
);

comment on table coms is 'any comment entries in the play/sub/com area';

comment on column coms.idx is 'where the comment belongs in the game events';

comment on column coms.description is 'a description of the play';

alter table coms owner to baseball;

create table data
(
	game_id char(12) not null
		constraint data_games_game_id_fk
			references games (game_id)
				on delete cascade,
	player_id char(8) not null,
	er integer not null,
	constraint data_pk
		primary key (game_id, player_id)
);

comment on table data is 'contains data that was parsed from the data tag';

alter table data owner to baseball;

create table subs
(
	game_id char(12) not null
		constraint subs_games_game_id_fk
			references games (game_id)
				on delete cascade,
	player_id char(8) not null,
	name varchar(64) not null,
	team integer not null,
	batting_order integer not null,
	position integer not null,
	constraint subs_pk
		primary key (game_id, player_id)
);

comment on table subs is 'the subs for the game';

comment on column subs.name is 'player''s name';

comment on column subs.team is '0 for visiting team, 1 for home team';

comment on column subs.batting_order is 'when the dh is used the pitcher is given the batting order position 0';

comment on column subs.position is 'DH has the position value of 10, pinch hitters are 11, and pinch runners are 12';

alter table subs owner to baseball;

