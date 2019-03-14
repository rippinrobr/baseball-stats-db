CREATE TABLE __schema_version (version VARCHAR(8) PRIMARY KEY NOT NULL);
CREATE TABLE games
(
game_id char(12) PRIMARY KEY,
season integer default 0 not null,
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
fieldcond varchar(16) not null,
precip varchar(16) not null,
sky varchar(16) not null,
temp integer default 0 not null,
winddir varchar(16) not null,
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
CREATE TABLE plays (
        game_id character(12) NOT NULL references games(game_id),
        idx integer NOT NULL,
        inning integer NOT NULL,
        team integer NOT NULL,
        player_id character(8) NOT NULL,
        count character varying(16) NOT NULL,
        pitches character varying(64) NOT NULL,
        event character varying(64) NOT NULL,
        PRIMARY KEY (game_id, idx)
      );
CREATE TABLE starters (
game_id character(12) NOT NULL references games(game_id),
player_id character(8) NOT NULL,
name character varying(64) NOT NULL,
team integer NOT NULL,
batting_order integer NOT NULL,
position integer NOT NULL,
PRIMARY KEY (game_id, player_id)
);
CREATE TABLE subs (
game_id character(12) NOT NULL references games(game_id),
"idx" integer NOT NULL,
player_id character(8) NOT NULL,
name character varying(64) NOT NULL,
team integer NOT NULL,
batting_order integer NOT NULL,
position integer NOT NULL,
PRIMARY KEY (game_id, player_id, idx)
);
CREATE TABLE coms (
    game_id character(12) NOT NULL references games(game_id),
    idx integer NOT NULL,
    description character varying(128) NOT NULL,
    PRIMARY KEY (game_id, idx)
);
CREATE TABLE data (
    game_id character(12) NOT NULL references games(game_id),
    player_id character(8) NOT NULL,
    er integer NOT NULL,
    PRIMARY KEY (game_id, player_id)
);
