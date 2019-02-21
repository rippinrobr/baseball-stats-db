--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1 (Debian 11.1-1.pgdg90+1)
-- Dumped by pg_dump version 11.1 (Debian 11.1-1.pgdg90+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: coms; Type: TABLE; Schema: public; Owner: baseball
--

CREATE TABLE public.coms (
    game_id character(12) NOT NULL,
    idx integer NOT NULL,
    description character varying(128) NOT NULL
);


ALTER TABLE public.coms OWNER TO baseball;

--
-- Name: TABLE coms; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON TABLE public.coms IS 'any comment entries in the play/sub/com area';


--
-- Name: COLUMN coms.idx; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.coms.idx IS 'where the comment belongs in the game events';


--
-- Name: COLUMN coms.description; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.coms.description IS 'a description of the play';


--
-- Name: data; Type: TABLE; Schema: public; Owner: baseball
--

CREATE TABLE public.data (
    game_id character(12) NOT NULL,
    player_id character(8) NOT NULL,
    er integer NOT NULL
);


ALTER TABLE public.data OWNER TO baseball;

--
-- Name: TABLE data; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON TABLE public.data IS 'contains data that was parsed from the data tag';


--
-- Name: games; Type: TABLE; Schema: public; Owner: baseball
--

CREATE TABLE public.games (
    game_id character(12) NOT NULL,
    visteam character(3) NOT NULL,
    hometeam character(3) NOT NULL,
    game_date date NOT NULL,
    number integer DEFAULT 0 NOT NULL,
    starttime time without time zone NOT NULL,
    daynight character varying(16) NOT NULL,
    usedh boolean DEFAULT false NOT NULL,
    pitches character varying(7) NOT NULL,
    umphome character(8),
    ump1b character(8),
    ump2b character(8),
    ump3b character(8),
    umplf character(8),
    umprf character(8),
    fieldcond character varying(7) NOT NULL,
    precip character varying(7) NOT NULL,
    sky character varying(7) NOT NULL,
    temp integer DEFAULT 0 NOT NULL,
    winddir character varying(7) NOT NULL,
    windspeed integer DEFAULT 0 NOT NULL,
    timeofgame integer DEFAULT 0 NOT NULL,
    attendance integer DEFAULT 0 NOT NULL,
    site character(5) NOT NULL,
    wp character(8) NOT NULL,
    lp character(8) NOT NULL,
    save character(8),
    gwrbi character(8),
    edittime timestamp without time zone,
    howscored character varying(32),
    inputprogvers character varying(32),
    inputter character varying(32),
    inputtime timestamp without time zone,
    scorer character varying(32),
    translator character varying(32)
);


ALTER TABLE public.games OWNER TO baseball;

--
-- Name: TABLE games; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON TABLE public.games IS 'Contains the basic information about a game, taken from the INFO tags for the game event';


--
-- Name: COLUMN games.game_id; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.games.game_id IS '12 character unique identifier for each game see retrosheet.org event file description for an indepth overview of the game_id';


--
-- Name: COLUMN games.number; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.games.number IS 'which game in the day this entry represents, 0 is first 1 is second, 2 is third, etc';


--
-- Name: COLUMN games.gwrbi; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.games.gwrbi IS 'game winning RBI';


--
-- Name: plays; Type: TABLE; Schema: public; Owner: baseball
--

CREATE TABLE public.plays (
    game_id character(12) NOT NULL,
    idx integer NOT NULL,
    inning integer NOT NULL,
    team integer NOT NULL,
    player_id character(8) NOT NULL,
    count character varying(16) NOT NULL,
    pitches character varying(32) NOT NULL,
    event character varying(32) NOT NULL
);


ALTER TABLE public.plays OWNER TO baseball;

--
-- Name: TABLE plays; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON TABLE public.plays IS 'contains the plays found for a particular game';


--
-- Name: COLUMN plays.idx; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.plays.idx IS 'the index into the plays in the game';


--
-- Name: COLUMN plays.team; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.plays.team IS '0 for visiting team, 1 for home team';


--
-- Name: COLUMN plays.count; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.plays.count IS 'the count at the time the play occurred, ?? is used when the count is unknown';


--
-- Name: COLUMN plays.pitches; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.plays.pitches IS 'the pitches thrown to the batter, must games will not have this data so this column will be empty';


--
-- Name: COLUMN plays.event; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.plays.event IS 'the result of the play';


--
-- Name: starters; Type: TABLE; Schema: public; Owner: baseball
--

CREATE TABLE public.starters (
    game_id character(12) NOT NULL,
    player_id character(8) NOT NULL,
    name character varying(64) NOT NULL,
    team integer NOT NULL,
    batting_order integer NOT NULL,
    "position" integer NOT NULL
);


ALTER TABLE public.starters OWNER TO baseball;

--
-- Name: TABLE starters; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON TABLE public.starters IS 'the starting lineups for the game';


--
-- Name: COLUMN starters.name; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.starters.name IS 'player''s name';


--
-- Name: COLUMN starters.team; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.starters.team IS '0 for visiting team, 1 for home team';


--
-- Name: COLUMN starters.batting_order; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.starters.batting_order IS 'when the dh is used the pitcher is given the batting order position 0';


--
-- Name: COLUMN starters."position"; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.starters."position" IS 'DH has the position value of 10';


--
-- Name: subs; Type: TABLE; Schema: public; Owner: baseball
--

CREATE TABLE public.subs (
    game_id character(12) NOT NULL,
    player_id character(8) NOT NULL,
    name character varying(64) NOT NULL,
    team integer NOT NULL,
    batting_order integer NOT NULL,
    "position" integer NOT NULL
);


ALTER TABLE public.subs OWNER TO baseball;

--
-- Name: TABLE subs; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON TABLE public.subs IS 'the subs for the game';


--
-- Name: COLUMN subs.name; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.subs.name IS 'player''s name';


--
-- Name: COLUMN subs.team; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.subs.team IS '0 for visiting team, 1 for home team';


--
-- Name: COLUMN subs.batting_order; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.subs.batting_order IS 'when the dh is used the pitcher is given the batting order position 0';


--
-- Name: COLUMN subs."position"; Type: COMMENT; Schema: public; Owner: baseball
--

COMMENT ON COLUMN public.subs."position" IS 'DH has the position value of 10, pinch hitters are 11, and pinch runners are 12';


--
-- Name: coms coms_pk; Type: CONSTRAINT; Schema: public; Owner: baseball
--

ALTER TABLE ONLY public.coms
    ADD CONSTRAINT coms_pk PRIMARY KEY (game_id, idx);


--
-- Name: data data_pk; Type: CONSTRAINT; Schema: public; Owner: baseball
--

ALTER TABLE ONLY public.data
    ADD CONSTRAINT data_pk PRIMARY KEY (game_id, player_id);


--
-- Name: games games_pk; Type: CONSTRAINT; Schema: public; Owner: baseball
--

ALTER TABLE ONLY public.games
    ADD CONSTRAINT games_pk UNIQUE (game_id);


--
-- Name: plays plays_pk; Type: CONSTRAINT; Schema: public; Owner: baseball
--

ALTER TABLE ONLY public.plays
    ADD CONSTRAINT plays_pk PRIMARY KEY (game_id, idx);


--
-- Name: starters starters_pk; Type: CONSTRAINT; Schema: public; Owner: baseball
--

ALTER TABLE ONLY public.starters
    ADD CONSTRAINT starters_pk PRIMARY KEY (game_id, player_id);


--
-- Name: subs subs_pk; Type: CONSTRAINT; Schema: public; Owner: baseball
--

ALTER TABLE ONLY public.subs
    ADD CONSTRAINT subs_pk PRIMARY KEY (game_id, player_id);


--
-- Name: games_game_id_uindex; Type: INDEX; Schema: public; Owner: baseball
--

CREATE UNIQUE INDEX games_game_id_uindex ON public.games USING btree (game_id);


--
-- Name: coms coms_games_game_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: baseball
--

ALTER TABLE ONLY public.coms
    ADD CONSTRAINT coms_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;


--
-- Name: data data_games_game_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: baseball
--

ALTER TABLE ONLY public.data
    ADD CONSTRAINT data_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;


--
-- Name: plays plays_games_game_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: baseball
--

ALTER TABLE ONLY public.plays
    ADD CONSTRAINT plays_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;


--
-- Name: starters starters_games_game_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: baseball
--

ALTER TABLE ONLY public.starters
    ADD CONSTRAINT starters_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;


--
-- Name: subs subs_games_game_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: baseball
--

ALTER TABLE ONLY public.subs
    ADD CONSTRAINT subs_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;


--
-- Name: COLUMN games.game_id; Type: ACL; Schema: public; Owner: baseball
--

GRANT SELECT(game_id),REFERENCES(game_id) ON TABLE public.games TO PUBLIC;


--
-- PostgreSQL database dump complete
--

