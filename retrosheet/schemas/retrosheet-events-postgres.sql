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

--
-- Name: diesel_manage_updated_at(regclass); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.diesel_manage_updated_at(_tbl regclass) RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    EXECUTE format('CREATE TRIGGER set_updated_at BEFORE UPDATE ON %s
                    FOR EACH ROW EXECUTE PROCEDURE diesel_set_updated_at()', _tbl);
END;
$$;


--
-- Name: diesel_set_updated_at(); Type: FUNCTION; Schema: public; Owner: -
--

CREATE FUNCTION public.diesel_set_updated_at() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF (
        NEW IS DISTINCT FROM OLD AND
        NEW.updated_at IS NOT DISTINCT FROM OLD.updated_at
    ) THEN
        NEW.updated_at := current_timestamp;
    END IF;
    RETURN NEW;
END;
$$;


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: __schema_version; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.__schema_version (
    version character varying(8) NOT NULL
);


--
-- Name: coms; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.coms (
    game_id character(12) NOT NULL,
    idx integer NOT NULL,
    description character varying(128) NOT NULL
);


--
-- Name: data; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.data (
    game_id character(12) NOT NULL,
    player_id character(8) NOT NULL,
    er integer NOT NULL
);


--
-- Name: games; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.games (
    game_id character(12) NOT NULL,
    season integer DEFAULT 0 NOT NULL,
    visteam character(3) NOT NULL,
    hometeam character(3) NOT NULL,
    game_date date NOT NULL,
    number integer DEFAULT 0 NOT NULL,
    starttime character varying(16) NOT NULL,
    daynight character varying(16) NOT NULL,
    usedh boolean DEFAULT false NOT NULL,
    pitches character varying(7) NOT NULL,
    umphome character varying(16),
    ump1b character varying(16),
    ump2b character varying(16),
    ump3b character varying(16),
    umplf character varying(16),
    umprf character varying(16),
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
    edittime character varying(32),
    howscored character varying(32),
    inputprogvers character varying(32),
    inputter character varying(32),
    inputtime character varying(32),
    scorer character varying(64),
    translator character varying(32)
);


--
-- Name: plays; Type: TABLE; Schema: public; Owner: -
--

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


--
-- Name: starters; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.starters (
    game_id character(12) NOT NULL,
    player_id character(8) NOT NULL,
    name character varying(64) NOT NULL,
    team integer NOT NULL,
    batting_order integer NOT NULL,
    "position" integer NOT NULL
);


--
-- Name: subs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.subs (
    game_id character(12) NOT NULL,
    player_id character(8) NOT NULL,
    idx integer NOT NULL,
    name character varying(64) NOT NULL,
    team integer NOT NULL,
    batting_order integer NOT NULL,
    "position" integer NOT NULL
);


--
-- Name: __schema_version __schema_version_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.__schema_version
    ADD CONSTRAINT __schema_version_pkey PRIMARY KEY (version);


--
-- Name: coms coms_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.coms
    ADD CONSTRAINT coms_pk PRIMARY KEY (game_id, idx);


--
-- Name: data data_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.data
    ADD CONSTRAINT data_pk PRIMARY KEY (game_id, player_id);


--
-- Name: games games_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.games
    ADD CONSTRAINT games_pk UNIQUE (game_id);


--
-- Name: plays plays_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.plays
    ADD CONSTRAINT plays_pk PRIMARY KEY (game_id, idx);


--
-- Name: starters starters_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.starters
    ADD CONSTRAINT starters_pk PRIMARY KEY (game_id, player_id);


--
-- Name: subs subs_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.subs
    ADD CONSTRAINT subs_pk PRIMARY KEY (game_id, player_id, idx);


--
-- Name: games_game_id_uindex; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX games_game_id_uindex ON public.games USING btree (game_id);


--
-- Name: coms coms_games_game_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.coms
    ADD CONSTRAINT coms_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;


--
-- Name: data data_games_game_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.data
    ADD CONSTRAINT data_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;


--
-- Name: plays plays_games_game_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.plays
    ADD CONSTRAINT plays_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;


--
-- Name: starters starters_games_game_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.starters
    ADD CONSTRAINT starters_games_game_id_fk FOREIGN KEY (game_id) REFERENCES public.games(game_id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

