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
-- Name: gamelogs gamelogs_pkey; Type: CONSTRAINT; Schema: public; Owner: robertrowe
--

ALTER TABLE ONLY gamelogs
    ADD CONSTRAINT gamelogs_pkey PRIMARY KEY (season, game_date_str, game_number, visitors_team, home_team, visitors_season_game_number);


--
-- Name: schedules schedules_pkey; Type: CONSTRAINT; Schema: public; Owner: robertrowe
--

ALTER TABLE ONLY schedules
    ADD CONSTRAINT schedules_pkey PRIMARY KEY (season, game_date_str, game_number, visitors_team, home_team, visitors_season_game_number);


--
-- PostgreSQL database dump complete
--

