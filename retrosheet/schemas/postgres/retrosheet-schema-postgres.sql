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
-- Name: __schema_version; Type: TABLE; Schema: public; Owner: baseball
--
CREATE TABLE __schema_version (version VARCHAR(8) PRIMARY KEY NOT NULL);
INSERT INTO __schema_version (version) values ('0.1.0');

--
-- Name: gamelogs; Type: TABLE; Schema: public; Owner: baseball
--

CREATE TABLE public.gamelogs (
    season integer NOT NULL,
    game_date_str varchar(8) NOT NULL,
  game_month integer NOT NULL,
  game_month_day integer NOT NULL,
  game_number integer NOT NULL,
  week_day integer NOT NULL,
  visitors_team varchar(3) NOT NULL,
  visitors_league varchar(3) NOT NULL,
  visitors_season_game_number integer NOT NULL,
  home_team varchar(3) NOT NULL,
  home_league varchar(3) NOT NULL,
  home_season_game_number integer NOT NULL,
  visitors_score integer NOT NULL,
  home_score integer NOT NULL,
  game_outs integer NOT NULL,
  game_type varchar(1) NOT NULL,
  completion_information varchar(128) NOT NULL DEFAULT '',
  forfeit_information varchar(128) NOT NULL DEFAULT '',
  protest_information varchar(128) NOT NULL DEFAULT '',
  park_id varchar(5) NOT NULL DEFAULT '',
  attendance integer NOT NULL,
  game_time_in_minutes integer NOT NULL,
  visitors_line_score varchar(32) NOT NULL DEFAULT '',
  home_line_score varchar(32) NOT NULL DEFAULT '',
  visitors_ab integer NOT NULL DEFAULT '0',
  visitors_h integer NOT NULL DEFAULT '0',
  visitors_doubles integer NOT NULL DEFAULT '0',
  visitors_triples integer NOT NULL DEFAULT '0',
  visitors_hr integer NOT NULL DEFAULT '0',
  visitors_rbi integer NOT NULL DEFAULT '0',
  visitors_sac_hits integer NOT NULL DEFAULT '0',
  visitors_sac_flies integer NOT NULL DEFAULT '0',
  visitors_hbp integer NOT NULL DEFAULT '0',
  visitors_bb integer NOT NULL DEFAULT '0',
  visitors_ibb integer NOT NULL DEFAULT '0',
  visitors_k integer NOT NULL DEFAULT '0',
  visitors_sb integer NOT NULL DEFAULT '0',
  visitors_cs integer NOT NULL DEFAULT '0',
  visitors_gidp integer NOT NULL DEFAULT '0',
  visitors_catcher_interference integer NOT NULL DEFAULT '0',
  visitors_lob integer NOT NULL DEFAULT '0',
  visitors_pitchers_used integer NOT NULL DEFAULT '0',
  visitors_individual_er integer NOT NULL DEFAULT '0',
  visitors_team_er integer NOT NULL DEFAULT '0',
  visitors_wp integer NOT NULL DEFAULT '0',
  visitors_balks integer NOT NULL DEFAULT '0',
  visitors_po integer NOT NULL DEFAULT '0',
  visitors_a integer NOT NULL DEFAULT '0',
  visitors_e integer NOT NULL DEFAULT '0',
  visitors_pb integer NOT NULL DEFAULT '0',
  visitors_dp integer NOT NULL DEFAULT '0',
  visitors_triple_plays integer NOT NULL DEFAULT '0',
  home_ab integer NOT NULL DEFAULT '0',
  home_h integer NOT NULL DEFAULT '0',
  home_doubles integer NOT NULL DEFAULT '0',
  home_triples integer NOT NULL DEFAULT '0',
  home_hr integer NOT NULL DEFAULT '0',
  home_rbi integer NOT NULL DEFAULT '0',
  home_sac_hits integer NOT NULL DEFAULT '0',
  home_sac_flies integer NOT NULL DEFAULT '0',
  home_hbp integer NOT NULL DEFAULT '0',
  home_bb integer NOT NULL DEFAULT '0',
  home_ibb integer NOT NULL DEFAULT '0',
  home_k integer NOT NULL DEFAULT '0',
  home_sb integer NOT NULL DEFAULT '0',
  home_cs integer NOT NULL DEFAULT '0',
  home_gidp integer NOT NULL DEFAULT '0',
  home_catcher_interference integer NOT NULL DEFAULT '0',
  home_lob integer NOT NULL DEFAULT '0',
  home_pitchers_used integer NOT NULL DEFAULT '0',
  home_individual_er integer NOT NULL DEFAULT '0',
  home_team_er integer NOT NULL DEFAULT '0',
  home_wp integer NOT NULL DEFAULT '0',
  home_balks integer NOT NULL DEFAULT '0',
  home_po integer NOT NULL DEFAULT '0',
  home_a integer NOT NULL DEFAULT '0',
  home_e integer NOT NULL DEFAULT '0',
  home_pb integer NOT NULL DEFAULT '0',
  home_dp integer NOT NULL DEFAULT '0',
  home_triple_plays integer NOT NULL DEFAULT '0',
  home_ump_id varchar(8) NOT NULL DEFAULT '',
  home_ump_name varchar(32) NOT NULL DEFAULT '',
  first_base_ump_id varchar(8) NOT NULL DEFAULT '',
  first_base_ump_name varchar(32) NOT NULL DEFAULT '',
  second_base_ump_id varchar(8) NOT NULL DEFAULT '',
  second_base_ump_name varchar(32) NOT NULL DEFAULT '',
  third_base_ump_id varchar(8) NOT NULL DEFAULT '',
  third_base_ump_name varchar(32) NOT NULL DEFAULT '',
  lf_ump_id varchar(8) NOT NULL DEFAULT '',
  lf_ump_name varchar(32) NOT NULL DEFAULT '',
  rf_ump_id varchar(8) NOT NULL DEFAULT '',
  rf_ump_name varchar(32) NOT NULL DEFAULT '',
  visitors_mgr_id varchar(8) NOT NULL DEFAULT '',
  visitors_mgr_name varchar(32) NOT NULL DEFAULT '',
  home_mgr_id varchar(8) NOT NULL DEFAULT '',
  home_mgr_name varchar(32) NOT NULL DEFAULT '',
  winning_pitcher_id varchar(8) NOT NULL DEFAULT '',
  winning_pitcher_name varchar(32) NOT NULL DEFAULT '',
  losing_pitcher_id varchar(8) NOT NULL DEFAULT '',
  losing_pitcher_name varchar(32) NOT NULL DEFAULT '',
  saving_pitcher_id varchar(8) NOT NULL DEFAULT '',
  saving_pitcher_name varchar(32) NOT NULL DEFAULT '',
  visitors_starting_pitcher_id varchar(8) NOT NULL DEFAULT '',
  visitors_starting_pitcher_name varchar(32) NOT NULL DEFAULT '',
  home_starting_pitcher_id varchar(8) NOT NULL DEFAULT '',
  home_starting_pitcher_name varchar(32) NOT NULL DEFAULT '',
  game_winning_rbi_batter_id varchar(8) NOT NULL DEFAULT '',
  game_winning_rbi_batter_name varchar(32) NOT NULL DEFAULT '',
  visitor_batter1_id varchar(8) NOT NULL DEFAULT '',
  visitor_batter1_name varchar(32) NOT NULL DEFAULT '',
  visitor_batter1_pos integer NOT NULL DEFAULT '0',
  visitor_batter2_id varchar(8) NOT NULL DEFAULT '',
  visitor_batter2_name varchar(32) NOT NULL DEFAULT '',
  visitor_batter2_pos integer NOT NULL DEFAULT '0',
  visitor_batter3_id varchar(8) NOT NULL DEFAULT '',
  visitor_batter3_name varchar(32) NOT NULL DEFAULT '',
  visitor_batter3_pos integer NOT NULL DEFAULT '0',
  visitor_batter4_id varchar(8) NOT NULL DEFAULT '',
  visitor_batter4_name varchar(32) NOT NULL DEFAULT '',
  visitor_batter4_pos integer NOT NULL DEFAULT '0',
  visitor_batter5_id varchar(8) NOT NULL DEFAULT '',
  visitor_batter5_name varchar(32) NOT NULL DEFAULT '',
  visitor_batter5_pos integer NOT NULL DEFAULT '0',
  visitor_batter6_id varchar(8) NOT NULL DEFAULT '',
  visitor_batter6_name varchar(32) NOT NULL DEFAULT '',
  visitor_batter6_pos integer NOT NULL DEFAULT '0',
  visitor_batter7_id varchar(8) NOT NULL DEFAULT '',
  visitor_batter7_name varchar(32) NOT NULL DEFAULT '',
  visitor_batter7_pos integer NOT NULL DEFAULT '0',
  visitor_batter8_id varchar(8) NOT NULL DEFAULT '',
  visitor_batter8_name varchar(32) NOT NULL DEFAULT '',
  visitor_batter8_pos integer NOT NULL DEFAULT '0',
  visitor_batter9_id varchar(8) NOT NULL DEFAULT '',
  visitor_batter9_name varchar(32) NOT NULL DEFAULT '',
  visitor_batter9_pos integer NOT NULL DEFAULT '0',
  home_batter1_id varchar(8) NOT NULL DEFAULT '',
  home_batter1_name varchar(32) NOT NULL DEFAULT '',
  home_batter1_pos integer NOT NULL DEFAULT '0',
  home_batter2_id varchar(8) NOT NULL DEFAULT '',
  home_batter2_name varchar(32) NOT NULL DEFAULT '',
  home_batter2_pos integer NOT NULL DEFAULT '0',
  home_batter3_id varchar(8) NOT NULL DEFAULT '',
  home_batter3_name varchar(32) NOT NULL DEFAULT '',
  home_batter3_pos integer NOT NULL DEFAULT '0',
  home_batter4_id varchar(8) NOT NULL DEFAULT '',
  home_batter4_name varchar(32) NOT NULL DEFAULT '',
  home_batter4_pos integer NOT NULL DEFAULT '0',
  home_batter5_id varchar(8) NOT NULL DEFAULT '',
  home_batter5_name varchar(32) NOT NULL DEFAULT '',
  home_batter5_pos integer NOT NULL DEFAULT '0',
  home_batter6_id varchar(8) NOT NULL DEFAULT '',
  home_batter6_name varchar(32) NOT NULL DEFAULT '',
  home_batter6_pos integer NOT NULL DEFAULT '0',
  home_batter7_id varchar(8) NOT NULL DEFAULT '',
  home_batter7_name varchar(32) NOT NULL DEFAULT '',
  home_batter7_pos integer NOT NULL DEFAULT '0',
  home_batter8_id varchar(8) NOT NULL DEFAULT '',
  home_batter8_name varchar(32) NOT NULL DEFAULT '',
  home_batter8_pos integer NOT NULL DEFAULT '0',
  home_batter9_id varchar(8) NOT NULL DEFAULT '',
  home_batter9_name varchar(32) NOT NULL DEFAULT '',
  home_batter9_pos integer NOT NULL DEFAULT '0',
  additional_information varchar(128) NOT NULL DEFAULT '',
  acquisition_information varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (season,game_date_str,game_number,visitors_team,home_team,visitors_season_game_number)
);


ALTER TABLE public.gamelogs OWNER TO baseball;

--
-- Name: schedules; Type: TABLE; Schema: public; Owner: baseball
--

CREATE TABLE public.schedules (
  season integer NOT NULL,
  game_date_str varchar(8) NOT NULL,
  game_number integer NOT NULL,
  visitors_team varchar(3) NOT NULL,
  home_team varchar(3) NOT NULL,
  game_month integer NOT NULL,
  game_month_day integer NOT NULL,
  week_day integer NOT NULL,
  visitors_league varchar(3) NOT NULL,
  visitors_season_game_number integer NOT NULL,
  home_league varchar(3) NOT NULL,
  home_season_game_number integer NOT NULL,
  game_type varchar(1) NOT NULL,
  postponed_canceled varchar(128) NOT NULL DEFAULT '',
  make_up_date varchar(128) NOT NULL DEFAULT '',
  PRIMARY KEY (season,game_date_str,game_number,visitors_team,home_team,visitors_season_game_number)

);


ALTER TABLE public.schedules OWNER TO baseball;

--
-- PostgreSQL database dump complete
--

