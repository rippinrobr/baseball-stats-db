--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.1
-- Dumped by pg_dump version 9.6.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: allstarfull; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE allstarfull (
    "playerID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    "gameNum" integer NOT NULL,
    "gameID" character varying(255),
    "teamID" character varying(255),
    "lgID" character varying(255),
    "GP" integer,
    "startingPos" integer
);


ALTER TABLE allstarfull OWNER TO postgres;

--
-- Name: appearances; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE appearances (
    "yearID" integer NOT NULL,
    "teamID" character varying(255) NOT NULL,
    "lgID" character varying(255),
    "playerID" character varying(255) NOT NULL,
    "G_all" integer NOT NULL,
    "GS" integer,
    "G_batting" integer,
    "G_defense" integer,
    "G_p" integer,
    "G_c" integer,
    "G_1b" integer,
    "G_2b" integer,
    "G_3b" integer,
    "G_ss" integer,
    "G_lf" integer,
    "G_cf" integer,
    "G_rf" integer,
    "G_of" integer,
    "G_dh" integer,
    "G_ph" integer,
    "G_pr" integer
);


ALTER TABLE appearances OWNER TO postgres;

--
-- Name: awardsmanagers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE awardsmanagers (
    "playerID" character varying(255) NOT NULL,
    "awardID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    "lgID" character varying(255) NOT NULL,
    tie character varying(255),
    notes character varying(255)
);


ALTER TABLE awardsmanagers OWNER TO postgres;

--
-- Name: awardsplayers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE awardsplayers (
    "playerID" character varying(255) NOT NULL,
    "awardID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    "lgID" character varying(255) NOT NULL,
    tie character varying(255),
    notes character varying(255)
);


ALTER TABLE awardsplayers OWNER TO postgres;

--
-- Name: awardssharemanagers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE awardssharemanagers (
    "awardID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    "lgID" character varying(255) NOT NULL,
    "playerID" character varying(255) NOT NULL,
    "pointsWon" integer,
    "pointsMax" integer,
    "votesFirst" integer
);


ALTER TABLE awardssharemanagers OWNER TO postgres;

--
-- Name: awardsshareplayers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE awardsshareplayers (
    "awardID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    "lgID" character varying(255) NOT NULL,
    "playerID" character varying(255) NOT NULL,
    "pointsWon" numeric(10,5),
    "pointsMax" integer,
    "votesFirst" numeric(10,5)
);


ALTER TABLE awardsshareplayers OWNER TO postgres;

--
-- Name: batting; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE batting (
    "playerID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    stint integer NOT NULL,
    "teamID" character varying(255) NOT NULL,
    "lgID" character varying(255),
    "G" integer,
    "AB" integer,
    "R" integer,
    "H" integer,
    doubles integer,
    triples integer,
    "HR" integer,
    "RBI" integer,
    "SB" integer,
    "CS" integer,
    "BB" integer,
    "SO" integer,
    "IBB" integer,
    "HBP" integer,
    "SH" integer,
    "SF" integer,
    "GIDP" integer
);


ALTER TABLE batting OWNER TO postgres;

--
-- Name: battingpost; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE battingpost (
    "yearID" integer NOT NULL,
    round character varying(255) NOT NULL,
    "playerID" character varying(255) NOT NULL,
    "teamID" character varying(255) NOT NULL,
    "lgID" character varying(255),
    "G" integer,
    "AB" integer,
    "R" integer,
    "H" integer,
    doubles integer,
    triples integer,
    "HR" integer,
    "RBI" integer,
    "SB" integer,
    "CS" integer,
    "BB" integer,
    "SO" integer,
    "IBB" integer,
    "HBP" integer,
    "SH" integer,
    "SF" integer,
    "GIDP" integer
);


ALTER TABLE battingpost OWNER TO postgres;

--
-- Name: collegeplaying; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE collegeplaying (
    "playerID" character varying(255) NOT NULL,
    "schoolID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL
);


ALTER TABLE collegeplaying OWNER TO postgres;

--
-- Name: fielding; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE fielding (
    "playerID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    stint integer NOT NULL,
    "teamID" character varying(255) NOT NULL,
    "lgID" character varying(255),
    "POS" character varying(255) NOT NULL,
    "A" integer,
    "GS" integer,
    "G" integer,
    "PB" integer,
    "InnOuts" integer,
    "ZR" integer,
    "PO" integer,
    "WP" integer,
    "CS" integer,
    "E" integer,
    "DP" integer,
    "SB" integer
);


ALTER TABLE fielding OWNER TO postgres;

--
-- Name: fieldingof; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE fieldingof (
    "playerID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    stint integer NOT NULL,
    "Glf" integer,
    "Gcf" integer,
    "Grf" integer
);


ALTER TABLE fieldingof OWNER TO postgres;

--
-- Name: fieldingofsplit; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE fieldingofsplit (
    "playerID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    stint integer NOT NULL,
    "teamID" character varying(255) NOT NULL,
    "lgID" character varying(255),
    "POS" character varying(255) NOT NULL,
    "G" integer,
    "GS" integer,
    "InnOuts" integer,
    "PO" integer,
    "A" integer,
    "E" integer,
    "DP" integer,
    "PB" integer,
    "WP" integer,
    "SB" integer,
    "CS" integer,
    "ZR" integer
);


ALTER TABLE fieldingofsplit OWNER TO postgres;

--
-- Name: fieldingpost; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE fieldingpost (
    "playerID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    "teamID" character varying(255),
    "lgID" character varying(255),
    round character varying(255) NOT NULL,
    "POS" character varying(255) NOT NULL,
    "G" integer,
    "GS" integer,
    "InnOuts" integer,
    "PO" integer,
    "A" integer,
    "E" integer,
    "DP" integer,
    "TP" integer,
    "PB" integer,
    "SB" integer,
    "CS" integer
);


ALTER TABLE fieldingpost OWNER TO postgres;

--
-- Name: gamelogs; Type: TABLE; Schema: public; Owner: robertrowe
--

CREATE TABLE gamelogs (
    season integer NOT NULL,
    game_date_str character varying(8) NOT NULL,
    game_month integer NOT NULL,
    game_month_day integer NOT NULL,
    game_number integer NOT NULL,
    week_day integer NOT NULL,
    visitors_team character varying(3) NOT NULL,
    visitors_league character varying(3) NOT NULL,
    visitors_season_game_number integer NOT NULL,
    home_team character varying(3) NOT NULL,
    home_league character varying(3) NOT NULL,
    home_season_game_number integer NOT NULL,
    visitors_score integer NOT NULL,
    home_score integer NOT NULL,
    game_outs integer NOT NULL,
    game_type character varying(1) NOT NULL,
    completion_information character varying(128) DEFAULT ''::character varying NOT NULL,
    forfeit_information character varying(128) DEFAULT ''::character varying NOT NULL,
    protest_information character varying(128) DEFAULT ''::character varying NOT NULL,
    park_id character varying(5) DEFAULT ''::character varying NOT NULL,
    attendance integer NOT NULL,
    game_time_in_minutes integer NOT NULL,
    visitors_line_score character varying(32) DEFAULT ''::character varying NOT NULL,
    home_line_score character varying(32) DEFAULT ''::character varying NOT NULL,
    visitors_ab integer DEFAULT 0 NOT NULL,
    visitors_h integer DEFAULT 0 NOT NULL,
    visitors_doubles integer DEFAULT 0 NOT NULL,
    visitors_triples integer DEFAULT 0 NOT NULL,
    visitors_hr integer DEFAULT 0 NOT NULL,
    visitors_rbi integer DEFAULT 0 NOT NULL,
    visitors_sac_hits integer DEFAULT 0 NOT NULL,
    visitors_sac_flies integer DEFAULT 0 NOT NULL,
    visitors_hbp integer DEFAULT 0 NOT NULL,
    visitors_bb integer DEFAULT 0 NOT NULL,
    visitors_ibb integer DEFAULT 0 NOT NULL,
    visitors_k integer DEFAULT 0 NOT NULL,
    visitors_sb integer DEFAULT 0 NOT NULL,
    visitors_cs integer DEFAULT 0 NOT NULL,
    visitors_gidp integer DEFAULT 0 NOT NULL,
    visitors_catcher_interference integer DEFAULT 0 NOT NULL,
    visitors_lob integer DEFAULT 0 NOT NULL,
    visitors_pitchers_used integer DEFAULT 0 NOT NULL,
    visitors_individual_er integer DEFAULT 0 NOT NULL,
    visitors_team_er integer DEFAULT 0 NOT NULL,
    visitors_wp integer DEFAULT 0 NOT NULL,
    visitors_balks integer DEFAULT 0 NOT NULL,
    visitors_po integer DEFAULT 0 NOT NULL,
    visitors_a integer DEFAULT 0 NOT NULL,
    visitors_e integer DEFAULT 0 NOT NULL,
    visitors_pb integer DEFAULT 0 NOT NULL,
    visitors_dp integer DEFAULT 0 NOT NULL,
    visitors_triple_plays integer DEFAULT 0 NOT NULL,
    home_ab integer DEFAULT 0 NOT NULL,
    home_h integer DEFAULT 0 NOT NULL,
    home_doubles integer DEFAULT 0 NOT NULL,
    home_triples integer DEFAULT 0 NOT NULL,
    home_hr integer DEFAULT 0 NOT NULL,
    home_rbi integer DEFAULT 0 NOT NULL,
    home_sac_hits integer DEFAULT 0 NOT NULL,
    home_sac_flies integer DEFAULT 0 NOT NULL,
    home_hbp integer DEFAULT 0 NOT NULL,
    home_bb integer DEFAULT 0 NOT NULL,
    home_ibb integer DEFAULT 0 NOT NULL,
    home_k integer DEFAULT 0 NOT NULL,
    home_sb integer DEFAULT 0 NOT NULL,
    home_cs integer DEFAULT 0 NOT NULL,
    home_gidp integer DEFAULT 0 NOT NULL,
    home_catcher_interference integer DEFAULT 0 NOT NULL,
    home_lob integer DEFAULT 0 NOT NULL,
    home_pitchers_used integer DEFAULT 0 NOT NULL,
    home_individual_er integer DEFAULT 0 NOT NULL,
    home_team_er integer DEFAULT 0 NOT NULL,
    home_wp integer DEFAULT 0 NOT NULL,
    home_balks integer DEFAULT 0 NOT NULL,
    home_po integer DEFAULT 0 NOT NULL,
    home_a integer DEFAULT 0 NOT NULL,
    home_e integer DEFAULT 0 NOT NULL,
    home_pb integer DEFAULT 0 NOT NULL,
    home_dp integer DEFAULT 0 NOT NULL,
    home_triple_plays integer DEFAULT 0 NOT NULL,
    home_ump_id character varying(8) DEFAULT ''::character varying NOT NULL,
    home_ump_name character varying(32) DEFAULT ''::character varying NOT NULL,
    first_base_ump_id character varying(8) DEFAULT ''::character varying NOT NULL,
    first_base_ump_name character varying(32) DEFAULT ''::character varying NOT NULL,
    second_base_ump_id character varying(8) DEFAULT ''::character varying NOT NULL,
    second_base_ump_name character varying(32) DEFAULT ''::character varying NOT NULL,
    third_base_ump_id character varying(8) DEFAULT ''::character varying NOT NULL,
    third_base_ump_name character varying(32) DEFAULT ''::character varying NOT NULL,
    lf_ump_id character varying(8) DEFAULT ''::character varying NOT NULL,
    lf_ump_name character varying(32) DEFAULT ''::character varying NOT NULL,
    rf_ump_id character varying(8) DEFAULT ''::character varying NOT NULL,
    rf_ump_name character varying(32) DEFAULT ''::character varying NOT NULL,
    visitors_mgr_id character varying(8) DEFAULT ''::character varying NOT NULL,
    visitors_mgr_name character varying(32) DEFAULT ''::character varying NOT NULL,
    home_mgr_id character varying(8) DEFAULT ''::character varying NOT NULL,
    home_mgr_name character varying(32) DEFAULT ''::character varying NOT NULL,
    winning_pitcher_id character varying(8) DEFAULT ''::character varying NOT NULL,
    winning_pitcher_name character varying(32) DEFAULT ''::character varying NOT NULL,
    losing_pitcher_id character varying(8) DEFAULT ''::character varying NOT NULL,
    losing_pitcher_name character varying(32) DEFAULT ''::character varying NOT NULL,
    saving_pitcher_id character varying(8) DEFAULT ''::character varying NOT NULL,
    saving_pitcher_name character varying(32) DEFAULT ''::character varying NOT NULL,
    visitors_starting_pitcher_id character varying(8) DEFAULT ''::character varying NOT NULL,
    visitors_starting_pitcher_name character varying(32) DEFAULT ''::character varying NOT NULL,
    home_starting_pitcher_id character varying(8) DEFAULT ''::character varying NOT NULL,
    home_starting_pitcher_name character varying(32) DEFAULT ''::character varying NOT NULL,
    game_winning_rbi_batter_id character varying(8) DEFAULT ''::character varying NOT NULL,
    game_winning_rbi_batter_name character varying(32) DEFAULT ''::character varying NOT NULL,
    visitor_batter1_id character varying(8) DEFAULT ''::character varying NOT NULL,
    visitor_batter1_name character varying(32) DEFAULT ''::character varying NOT NULL,
    visitor_batter1_pos integer DEFAULT 0 NOT NULL,
    visitor_batter2_id character varying(8) DEFAULT ''::character varying NOT NULL,
    visitor_batter2_name character varying(32) DEFAULT ''::character varying NOT NULL,
    visitor_batter2_pos integer DEFAULT 0 NOT NULL,
    visitor_batter3_id character varying(8) DEFAULT ''::character varying NOT NULL,
    visitor_batter3_name character varying(32) DEFAULT ''::character varying NOT NULL,
    visitor_batter3_pos integer DEFAULT 0 NOT NULL,
    visitor_batter4_id character varying(8) DEFAULT ''::character varying NOT NULL,
    visitor_batter4_name character varying(32) DEFAULT ''::character varying NOT NULL,
    visitor_batter4_pos integer DEFAULT 0 NOT NULL,
    visitor_batter5_id character varying(8) DEFAULT ''::character varying NOT NULL,
    visitor_batter5_name character varying(32) DEFAULT ''::character varying NOT NULL,
    visitor_batter5_pos integer DEFAULT 0 NOT NULL,
    visitor_batter6_id character varying(8) DEFAULT ''::character varying NOT NULL,
    visitor_batter6_name character varying(32) DEFAULT ''::character varying NOT NULL,
    visitor_batter6_pos integer DEFAULT 0 NOT NULL,
    visitor_batter7_id character varying(8) DEFAULT ''::character varying NOT NULL,
    visitor_batter7_name character varying(32) DEFAULT ''::character varying NOT NULL,
    visitor_batter7_pos integer DEFAULT 0 NOT NULL,
    visitor_batter8_id character varying(8) DEFAULT ''::character varying NOT NULL,
    visitor_batter8_name character varying(32) DEFAULT ''::character varying NOT NULL,
    visitor_batter8_pos integer DEFAULT 0 NOT NULL,
    visitor_batter9_id character varying(8) DEFAULT ''::character varying NOT NULL,
    visitor_batter9_name character varying(32) DEFAULT ''::character varying NOT NULL,
    visitor_batter9_pos integer DEFAULT 0 NOT NULL,
    home_batter1_id character varying(8) DEFAULT ''::character varying NOT NULL,
    home_batter1_name character varying(32) DEFAULT ''::character varying NOT NULL,
    home_batter1_pos integer DEFAULT 0 NOT NULL,
    home_batter2_id character varying(8) DEFAULT ''::character varying NOT NULL,
    home_batter2_name character varying(32) DEFAULT ''::character varying NOT NULL,
    home_batter2_pos integer DEFAULT 0 NOT NULL,
    home_batter3_id character varying(8) DEFAULT ''::character varying NOT NULL,
    home_batter3_name character varying(32) DEFAULT ''::character varying NOT NULL,
    home_batter3_pos integer DEFAULT 0 NOT NULL,
    home_batter4_id character varying(8) DEFAULT ''::character varying NOT NULL,
    home_batter4_name character varying(32) DEFAULT ''::character varying NOT NULL,
    home_batter4_pos integer DEFAULT 0 NOT NULL,
    home_batter5_id character varying(8) DEFAULT ''::character varying NOT NULL,
    home_batter5_name character varying(32) DEFAULT ''::character varying NOT NULL,
    home_batter5_pos integer DEFAULT 0 NOT NULL,
    home_batter6_id character varying(8) DEFAULT ''::character varying NOT NULL,
    home_batter6_name character varying(32) DEFAULT ''::character varying NOT NULL,
    home_batter6_pos integer DEFAULT 0 NOT NULL,
    home_batter7_id character varying(8) DEFAULT ''::character varying NOT NULL,
    home_batter7_name character varying(32) DEFAULT ''::character varying NOT NULL,
    home_batter7_pos integer DEFAULT 0 NOT NULL,
    home_batter8_id character varying(8) DEFAULT ''::character varying NOT NULL,
    home_batter8_name character varying(32) DEFAULT ''::character varying NOT NULL,
    home_batter8_pos integer DEFAULT 0 NOT NULL,
    home_batter9_id character varying(8) DEFAULT ''::character varying NOT NULL,
    home_batter9_name character varying(32) DEFAULT ''::character varying NOT NULL,
    home_batter9_pos integer DEFAULT 0 NOT NULL,
    additional_information character varying(128) DEFAULT ''::character varying NOT NULL,
    acquisition_information character varying(64) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE gamelogs OWNER TO robertrowe;

--
-- Name: halloffame; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE halloffame (
    "playerID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    "votedBy" character varying(255) NOT NULL,
    ballots integer,
    needed integer,
    votes integer,
    inducted character varying(255),
    category character varying(255) NOT NULL,
    needed_note character varying(255)
);


ALTER TABLE halloffame OWNER TO postgres;

--
-- Name: homegames; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE homegames (
    yearkey integer NOT NULL,
    parkkey character varying(255) NOT NULL,
    teamkey character varying(255) NOT NULL,
    leaguekey character varying(255),
    attendance integer,
    games integer,
    spanfirst character varying(255),
    spanlast character varying(255),
    openings integer
);


ALTER TABLE homegames OWNER TO postgres;

--
-- Name: managers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE managers (
    "playerID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    "teamID" character varying(255) NOT NULL,
    "lgID" character varying(255) NOT NULL,
    inseason integer NOT NULL,
    "G" integer,
    "W" integer,
    "L" integer,
    rank integer,
    "plyrMgr" character varying(255)
);


ALTER TABLE managers OWNER TO postgres;

--
-- Name: managershalf; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE managershalf (
    "playerID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    "teamID" character varying(255) NOT NULL,
    "lgID" character varying(255) NOT NULL,
    inseason integer,
    half integer NOT NULL,
    "G" integer,
    "W" integer,
    "L" integer,
    rank integer
);


ALTER TABLE managershalf OWNER TO postgres;

--
-- Name: parks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE parks (
    parkkey character varying(255) NOT NULL,
    parkname character varying(255),
    parkalias character varying(255),
    alias character varying(255),
    city character varying(255),
    state character varying(255),
    country character varying(255)
);


ALTER TABLE parks OWNER TO postgres;

--
-- Name: people; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE people (
    "playerID" character varying(9) NOT NULL,
    "birthYear" integer,
    "birthMonth" integer,
    "birthDay" integer,
    "birthCountry" character varying(14),
    "birthState" character varying(22),
    "birthCity" character varying(26),
    "deathYear" integer,
    "deathMonth" integer,
    "deathDay" integer,
    "deathCountry" character varying(14),
    "deathState" character varying(20),
    "deathCity" character varying(28),
    "nameFirst" character varying(12),
    "nameLast" character varying(14),
    "nameGiven" character varying(43),
    weight integer,
    height integer,
    bats character varying(255),
    throws character varying(255),
    debut character varying(255),
    "finalGame" character varying(255),
    "retroID" character varying(255),
    "bbrefID" character varying(9)
);


ALTER TABLE people OWNER TO postgres;

--
-- Name: pitching; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE pitching (
    "playerID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    stint integer NOT NULL,
    "teamID" character varying(255) NOT NULL,
    "lgID" character varying(255),
    "W" integer,
    "L" integer,
    "G" integer,
    "GS" integer,
    "CG" integer,
    "SHO" integer,
    "SV" integer,
    "IPouts" integer,
    "H" integer,
    "ER" integer,
    "HR" integer,
    "BB" integer,
    "SO" integer,
    "BAOpp" real,
    "ERA" real,
    "IBB" integer,
    "WP" integer,
    "HBP" integer,
    "BK" integer,
    "BFP" integer,
    "GF" integer,
    "R" integer,
    "SH" integer,
    "SF" integer,
    "GIDP" integer
);


ALTER TABLE pitching OWNER TO postgres;

--
-- Name: pitchingpost; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE pitchingpost (
    "playerID" character varying(255) NOT NULL,
    "yearID" integer NOT NULL,
    round character varying(255) NOT NULL,
    "teamID" character varying(255),
    "lgID" character varying(255),
    "W" integer,
    "L" integer,
    "G" integer,
    "GS" integer,
    "CG" integer,
    "SHO" integer,
    "SV" integer,
    "IPouts" integer,
    "H" integer,
    "ER" integer,
    "HR" integer,
    "BB" integer,
    "SO" integer,
    "BAOpp" real,
    "ERA" real,
    "IBB" integer,
    "WP" integer,
    "HBP" integer,
    "BK" integer,
    "BFP" integer,
    "GF" integer,
    "R" integer,
    "SH" integer,
    "SF" integer,
    "GIDP" integer
);


ALTER TABLE pitchingpost OWNER TO postgres;

--
-- Name: salaries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE salaries (
    "yearID" integer NOT NULL,
    "teamID" character varying(255) NOT NULL,
    "lgID" character varying(255),
    "playerID" character varying(255) NOT NULL,
    salary integer
);


ALTER TABLE salaries OWNER TO postgres;

--
-- Name: schedules; Type: TABLE; Schema: public; Owner: robertrowe
--

CREATE TABLE schedules (
    season integer NOT NULL,
    game_date_str character varying(8) NOT NULL,
    game_number integer NOT NULL,
    visitors_team character varying(3) NOT NULL,
    home_team character varying(3) NOT NULL,
    game_month integer NOT NULL,
    game_month_day integer NOT NULL,
    week_day integer NOT NULL,
    visitors_league character varying(3) NOT NULL,
    visitors_season_game_number integer NOT NULL,
    home_league character varying(3) NOT NULL,
    home_season_game_number integer NOT NULL,
    game_type character varying(1) NOT NULL,
    postponed_canceled character varying(128) DEFAULT ''::character varying NOT NULL,
    make_up_date character varying(128) DEFAULT ''::character varying NOT NULL
);


ALTER TABLE schedules OWNER TO robertrowe;

--
-- Name: schools; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE schools (
    "schoolID" character varying(255) NOT NULL,
    name_full character varying(255),
    city character varying(255),
    state character varying(255),
    country character varying(255)
);


ALTER TABLE schools OWNER TO postgres;

--
-- Name: seriespost; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE seriespost (
    "yearID" integer NOT NULL,
    round character varying(255) NOT NULL,
    "teamIDwinner" character varying(255),
    "lgIDwinner" character varying(255),
    "teamIDloser" character varying(255),
    "lgIDloser" character varying(255),
    wins integer,
    losses integer,
    ties integer
);


ALTER TABLE seriespost OWNER TO postgres;

--
-- Name: teams; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE teams (
    "yearID" integer NOT NULL,
    "lgID" character varying(255) NOT NULL,
    "teamID" character varying(255) NOT NULL,
    "franchID" character varying(255),
    "divID" character varying(255),
    "Rank" integer,
    "G" integer,
    "Ghome" integer,
    "W" integer,
    "L" integer,
    "DivWin" character varying(255),
    "WCWin" character varying(255),
    "LgWin" character varying(255),
    "WSWin" character varying(255),
    "R" integer,
    "AB" integer,
    "H" integer,
    doubles integer,
    triples integer,
    "HR" integer,
    "BB" integer,
    "SO" integer,
    "SB" integer,
    "CS" integer,
    "HBP" integer,
    "SF" integer,
    "RA" integer,
    "ER" integer,
    "ERA" real,
    "CG" integer,
    "SHO" integer,
    "SV" integer,
    "IPouts" integer,
    "HA" integer,
    "HRA" integer,
    "BBA" integer,
    "SOA" integer,
    "E" integer,
    "DP" integer,
    "FP" real,
    name character varying(255),
    park character varying(255),
    attendance integer,
    "BPF" integer,
    "PPF" integer,
    "teamIDBR" character varying(255),
    "teamIDlahman45" character varying(255),
    "teamIDretro" character varying(255)
);


ALTER TABLE teams OWNER TO postgres;

--
-- Name: teamsfranchises; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE teamsfranchises (
    "franchID" character varying(255) NOT NULL,
    "franchName" character varying(255),
    active character varying(255),
    "NAassoc" character varying(255)
);


ALTER TABLE teamsfranchises OWNER TO postgres;

--
-- Name: teamshalf; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE teamshalf (
    "yearID" integer NOT NULL,
    "lgID" character varying(255),
    "teamID" character varying(255) NOT NULL,
    "Half" integer NOT NULL,
    "divID" character varying(255),
    "DivWin" character varying(255),
    "Rank" integer,
    "G" integer,
    "W" integer,
    "L" integer
);


ALTER TABLE teamshalf OWNER TO postgres;

--
-- Name: allstarfull allstarfull_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY allstarfull
    ADD CONSTRAINT allstarfull_pkey PRIMARY KEY ("playerID", "yearID", "gameNum");


--
-- Name: appearances appearances_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY appearances
    ADD CONSTRAINT appearances_pkey PRIMARY KEY ("yearID", "teamID", "playerID", "G_all");


--
-- Name: awardsmanagers awardsmanagers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY awardsmanagers
    ADD CONSTRAINT awardsmanagers_pkey PRIMARY KEY ("yearID", "lgID", "awardID", "playerID");


--
-- Name: awardsplayers awardsplayers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY awardsplayers
    ADD CONSTRAINT awardsplayers_pkey PRIMARY KEY ("yearID", "lgID", "awardID", "playerID");


--
-- Name: awardssharemanagers awardssharemanagers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY awardssharemanagers
    ADD CONSTRAINT awardssharemanagers_pkey PRIMARY KEY ("awardID", "yearID", "lgID", "playerID");


--
-- Name: awardsshareplayers awardsshareplayers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY awardsshareplayers
    ADD CONSTRAINT awardsshareplayers_pkey PRIMARY KEY ("awardID", "yearID", "lgID", "playerID");


--
-- Name: batting batting_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY batting
    ADD CONSTRAINT batting_pkey PRIMARY KEY ("playerID", "yearID", "teamID", stint);


--
-- Name: battingpost battingpost_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY battingpost
    ADD CONSTRAINT battingpost_pkey PRIMARY KEY ("yearID", round, "playerID", "teamID");


--
-- Name: collegeplaying collegeplaying_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY collegeplaying
    ADD CONSTRAINT collegeplaying_pkey PRIMARY KEY ("playerID", "schoolID", "yearID");


--
-- Name: fielding fielding_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY fielding
    ADD CONSTRAINT fielding_pkey PRIMARY KEY ("playerID", "yearID", "teamID", "POS", stint);


--
-- Name: fieldingof fieldingof_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY fieldingof
    ADD CONSTRAINT fieldingof_pkey PRIMARY KEY ("playerID", "yearID", stint);


--
-- Name: fieldingofsplit fieldingofsplit_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY fieldingofsplit
    ADD CONSTRAINT fieldingofsplit_pkey PRIMARY KEY ("playerID", "yearID", "teamID", "POS", stint);


--
-- Name: fieldingpost fieldingpost_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY fieldingpost
    ADD CONSTRAINT fieldingpost_pkey PRIMARY KEY ("playerID", "yearID", round, "POS");


--
-- Name: gamelogs gamelogs_pkey; Type: CONSTRAINT; Schema: public; Owner: robertrowe
--

ALTER TABLE ONLY gamelogs
    ADD CONSTRAINT gamelogs_pkey PRIMARY KEY (season, game_date_str, game_number, visitors_team, home_team, visitors_season_game_number);


--
-- Name: halloffame halloffame_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY halloffame
    ADD CONSTRAINT halloffame_pkey PRIMARY KEY ("playerID", category, "yearID", "votedBy");


--
-- Name: homegames homegames_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY homegames
    ADD CONSTRAINT homegames_pkey PRIMARY KEY (yearkey, parkkey, teamkey);


--
-- Name: managers managers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY managers
    ADD CONSTRAINT managers_pkey PRIMARY KEY ("yearID", "lgID", "teamID", "playerID", inseason);


--
-- Name: managershalf managershalf_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY managershalf
    ADD CONSTRAINT managershalf_pkey PRIMARY KEY ("playerID", "teamID", "lgID", "yearID", half);


--
-- Name: parks parks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY parks
    ADD CONSTRAINT parks_pkey PRIMARY KEY (parkkey);


--
-- Name: people people_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY people
    ADD CONSTRAINT people_pkey PRIMARY KEY ("playerID");


--
-- Name: pitching pitching_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY pitching
    ADD CONSTRAINT pitching_pkey PRIMARY KEY ("playerID", "yearID", stint, "teamID");


--
-- Name: pitchingpost pitchingpost_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY pitchingpost
    ADD CONSTRAINT pitchingpost_pkey PRIMARY KEY ("playerID", "yearID", round);


--
-- Name: salaries salaries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY salaries
    ADD CONSTRAINT salaries_pkey PRIMARY KEY ("yearID", "teamID", "playerID");


--
-- Name: schedules schedules_pkey; Type: CONSTRAINT; Schema: public; Owner: robertrowe
--

ALTER TABLE ONLY schedules
    ADD CONSTRAINT schedules_pkey PRIMARY KEY (season, game_date_str, game_number, visitors_team, home_team, visitors_season_game_number);


--
-- Name: schools schools_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY schools
    ADD CONSTRAINT schools_pkey PRIMARY KEY ("schoolID");


--
-- Name: seriespost seriespost_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY seriespost
    ADD CONSTRAINT seriespost_pkey PRIMARY KEY ("yearID", round);


--
-- Name: teams teams_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY teams
    ADD CONSTRAINT teams_pkey PRIMARY KEY ("yearID", "lgID", "teamID");


--
-- Name: teamsfranchises teamsfranchises_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY teamsfranchises
    ADD CONSTRAINT teamsfranchises_pkey PRIMARY KEY ("franchID");


--
-- Name: teamshalf teamshalf_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY teamshalf
    ADD CONSTRAINT teamshalf_pkey PRIMARY KEY ("yearID", "teamID", "Half");


--
-- PostgreSQL database dump complete
--

