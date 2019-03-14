--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1 (Debian 11.1-1.pgdg90+1)
-- Dumped by pg_dump version 11.1 (Debian 11.1-1.pgdg90+1)

CREATE TABLE __schema_version (version VARCHAR(8) PRIMARY KEY NOT NULL);
INSERT INTO __schema_version (version) values ('0.1.0');

CREATE TABLE public.coms (
    game_id character(12) NOT NULL,
    idx integer NOT NULL,
    description character varying(128) NOT NULL
);

CREATE TABLE public.data (
    game_id character(12) NOT NULL,
    player_id character(8) NOT NULL,
    er integer NOT NULL
);

CREATE TABLE public.games (
    game_id character(12) NOT NULL,
    season integer DEFAULT 0 NOT NULL,
    visteam character(3) NOT NULL,
    hometeam character(3) NOT NULL,
    game_date date NOT NULL,
    number integer DEFAULT 0 NOT NULL,
    starttime varchar(16) NOT NULL,
    daynight character varying(16) NOT NULL,
    usedh boolean DEFAULT false NOT NULL,
    pitches character varying(7) NOT NULL,
    umphome character(16),
    ump1b character(16),
    ump2b character(16),
    ump3b character(16),
    umplf character(16),
    umprf character(16),
    fieldcond character varying(16) NOT NULL,
    precip character varying(16) NOT NULL,
    sky character varying(16) NOT NULL,
    temp integer DEFAULT 0 NOT NULL,
    winddir character varying(16) NOT NULL,
    windspeed integer DEFAULT 0 NOT NULL,
    timeofgame integer DEFAULT 0 NOT NULL,
    attendance integer DEFAULT 0 NOT NULL,
    site character(5) NOT NULL,
    wp character(8) NOT NULL,
    lp character(8) NOT NULL,
    save character(8),
    gwrbi character(8),
    edittime varchar(32) ,
    howscored character varying(32),
    inputprogvers character varying(32),
    inputter character varying(32),
    inputtime varchar(32),
    scorer character varying(64),
    translator character varying(32)
);


CREATE TABLE public.plays (
    game_id character(12) NOT NULL,
    idx integer NOT NULL,
    inning integer NOT NULL,
    team integer NOT NULL,
    player_id character(8) NOT NULL,
    count character varying(16) NOT NULL,
    pitches character varying(64) NOT NULL,
    event character varying(64) NOT NULL
);


CREATE TABLE public.starters (
    game_id character(12) NOT NULL,
    player_id character(8) NOT NULL,
    name character varying(64) NOT NULL,
    team integer NOT NULL,
    batting_order integer NOT NULL,
    "position" integer NOT NULL
);


CREATE TABLE public.subs (
    game_id character(12) NOT NULL,
    player_id character(8) NOT NULL,
    idx integer NOT NULL,
    name character varying(64) NOT NULL,
    team integer NOT NULL,
    batting_order integer NOT NULL,
    "position" integer NOT NULL
);


ALTER TABLE ONLY public.coms
    ADD CONSTRAINT coms_pk PRIMARY KEY (game_id, idx);

ALTER TABLE ONLY public.data
    ADD CONSTRAINT data_pk PRIMARY KEY (game_id, player_id);

ALTER TABLE ONLY public.games
    ADD CONSTRAINT games_pk UNIQUE (game_id);

ALTER TABLE ONLY public.plays
    ADD CONSTRAINT plays_pk PRIMARY KEY (game_id, idx);

ALTER TABLE ONLY public.starters
    ADD CONSTRAINT starters_pk PRIMARY KEY (game_id, player_id);

ALTER TABLE ONLY public.subs
    ADD CONSTRAINT subs_pk PRIMARY KEY (game_id, player_id, idx);

CREATE UNIQUE INDEX games_game_id_uindex ON public.games USING btree (game_id);

ALTER TABLE ONLY public.coms
    ADD CONSTRAINT coms_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;

ALTER TABLE ONLY public.data
    ADD CONSTRAINT data_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;

ALTER TABLE ONLY public.plays
    ADD CONSTRAINT plays_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;

ALTER TABLE ONLY public.starters
    ADD CONSTRAINT starters_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;
