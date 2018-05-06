-- MySQL dump 10.13  Distrib 5.5.60, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: retrosheet_2017
-- ------------------------------------------------------
-- Server version	5.5.60-0+deb8u1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `gamelogs`
--

DROP TABLE IF EXISTS `gamelogs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `gamelogs` (
  `season` int(11) NOT NULL,
  `game_date_str` varchar(8) NOT NULL,
  `game_month` int(11) NOT NULL,
  `game_month_day` int(11) NOT NULL,
  `game_number` int(11) NOT NULL,
  `week_day` int(11) NOT NULL,
  `visitors_team` varchar(3) NOT NULL,
  `visitors_league` varchar(3) NOT NULL,
  `visitors_season_game_number` int(11) NOT NULL,
  `home_team` varchar(3) NOT NULL,
  `home_league` varchar(3) NOT NULL,
  `home_season_game_number` int(11) NOT NULL,
  `visitors_score` int(11) NOT NULL,
  `home_score` int(11) NOT NULL,
  `game_outs` int(11) NOT NULL,
  `game_type` varchar(1) NOT NULL,
  `completion_information` varchar(128) NOT NULL DEFAULT '',
  `forfeit_information` varchar(128) NOT NULL DEFAULT '',
  `protest_information` varchar(128) NOT NULL DEFAULT '',
  `park_id` varchar(5) NOT NULL DEFAULT '',
  `attendance` int(11) NOT NULL,
  `game_time_in_minutes` int(11) NOT NULL,
  `visitors_line_score` varchar(32) NOT NULL DEFAULT '',
  `home_line_score` varchar(32) NOT NULL DEFAULT '',
  `visitors_ab` int(11) NOT NULL DEFAULT '0',
  `visitors_h` int(11) NOT NULL DEFAULT '0',
  `visitors_doubles` int(11) NOT NULL DEFAULT '0',
  `visitors_triples` int(11) NOT NULL DEFAULT '0',
  `visitors_hr` int(11) NOT NULL DEFAULT '0',
  `visitors_rbi` int(11) NOT NULL DEFAULT '0',
  `visitors_sac_hits` int(11) NOT NULL DEFAULT '0',
  `visitors_sac_flies` int(11) NOT NULL DEFAULT '0',
  `visitors_hbp` int(11) NOT NULL DEFAULT '0',
  `visitors_bb` int(11) NOT NULL DEFAULT '0',
  `visitors_ibb` int(11) NOT NULL DEFAULT '0',
  `visitors_k` int(11) NOT NULL DEFAULT '0',
  `visitors_sb` int(11) NOT NULL DEFAULT '0',
  `visitors_cs` int(11) NOT NULL DEFAULT '0',
  `visitors_gidp` int(11) NOT NULL DEFAULT '0',
  `visitors_catcher_interference` int(11) NOT NULL DEFAULT '0',
  `visitors_lob` int(11) NOT NULL DEFAULT '0',
  `visitors_pitchers_used` int(11) NOT NULL DEFAULT '0',
  `visitors_individual_er` int(11) NOT NULL DEFAULT '0',
  `visitors_team_er` int(11) NOT NULL DEFAULT '0',
  `visitors_wp` int(11) NOT NULL DEFAULT '0',
  `visitors_balks` int(11) NOT NULL DEFAULT '0',
  `visitors_po` int(11) NOT NULL DEFAULT '0',
  `visitors_a` int(11) NOT NULL DEFAULT '0',
  `visitors_e` int(11) NOT NULL DEFAULT '0',
  `visitors_pb` int(11) NOT NULL DEFAULT '0',
  `visitors_dp` int(11) NOT NULL DEFAULT '0',
  `visitors_triple_plays` int(11) NOT NULL DEFAULT '0',
  `home_ab` int(11) NOT NULL DEFAULT '0',
  `home_h` int(11) NOT NULL DEFAULT '0',
  `home_doubles` int(11) NOT NULL DEFAULT '0',
  `home_triples` int(11) NOT NULL DEFAULT '0',
  `home_hr` int(11) NOT NULL DEFAULT '0',
  `home_rbi` int(11) NOT NULL DEFAULT '0',
  `home_sac_hits` int(11) NOT NULL DEFAULT '0',
  `home_sac_flies` int(11) NOT NULL DEFAULT '0',
  `home_hbp` int(11) NOT NULL DEFAULT '0',
  `home_bb` int(11) NOT NULL DEFAULT '0',
  `home_ibb` int(11) NOT NULL DEFAULT '0',
  `home_k` int(11) NOT NULL DEFAULT '0',
  `home_sb` int(11) NOT NULL DEFAULT '0',
  `home_cs` int(11) NOT NULL DEFAULT '0',
  `home_gidp` int(11) NOT NULL DEFAULT '0',
  `home_catcher_interference` int(11) NOT NULL DEFAULT '0',
  `home_lob` int(11) NOT NULL DEFAULT '0',
  `home_pitchers_used` int(11) NOT NULL DEFAULT '0',
  `home_individual_er` int(11) NOT NULL DEFAULT '0',
  `home_team_er` int(11) NOT NULL DEFAULT '0',
  `home_wp` int(11) NOT NULL DEFAULT '0',
  `home_balks` int(11) NOT NULL DEFAULT '0',
  `home_po` int(11) NOT NULL DEFAULT '0',
  `home_a` int(11) NOT NULL DEFAULT '0',
  `home_e` int(11) NOT NULL DEFAULT '0',
  `home_pb` int(11) NOT NULL DEFAULT '0',
  `home_dp` int(11) NOT NULL DEFAULT '0',
  `home_triple_plays` int(11) NOT NULL DEFAULT '0',
  `home_ump_id` varchar(8) NOT NULL DEFAULT '',
  `home_ump_name` varchar(32) NOT NULL DEFAULT '',
  `first_base_ump_id` varchar(8) NOT NULL DEFAULT '',
  `first_base_ump_name` varchar(32) NOT NULL DEFAULT '',
  `second_base_ump_id` varchar(8) NOT NULL DEFAULT '',
  `second_base_ump_name` varchar(32) NOT NULL DEFAULT '',
  `third_base_ump_id` varchar(8) NOT NULL DEFAULT '',
  `third_base_ump_name` varchar(32) NOT NULL DEFAULT '',
  `lf_ump_id` varchar(8) NOT NULL DEFAULT '',
  `lf_ump_name` varchar(32) NOT NULL DEFAULT '',
  `rf_ump_id` varchar(8) NOT NULL DEFAULT '',
  `rf_ump_name` varchar(32) NOT NULL DEFAULT '',
  `visitors_mgr_id` varchar(8) NOT NULL DEFAULT '',
  `visitors_mgr_name` varchar(32) NOT NULL DEFAULT '',
  `home_mgr_id` varchar(8) NOT NULL DEFAULT '',
  `home_mgr_name` varchar(32) NOT NULL DEFAULT '',
  `winning_pitcher_id` varchar(8) NOT NULL DEFAULT '',
  `winning_pitcher_name` varchar(32) NOT NULL DEFAULT '',
  `losing_pitcher_id` varchar(8) NOT NULL DEFAULT '',
  `losing_pitcher_name` varchar(32) NOT NULL DEFAULT '',
  `saving_pitcher_id` varchar(8) NOT NULL DEFAULT '',
  `saving_pitcher_name` varchar(32) NOT NULL DEFAULT '',
  `visitors_starting_pitcher_id` varchar(8) NOT NULL DEFAULT '',
  `visitors_starting_pitcher_name` varchar(32) NOT NULL DEFAULT '',
  `home_starting_pitcher_id` varchar(8) NOT NULL DEFAULT '',
  `home_starting_pitcher_name` varchar(32) NOT NULL DEFAULT '',
  `game_winning_rbi_batter_id` varchar(8) NOT NULL DEFAULT '',
  `game_winning_rbi_batter_name` varchar(32) NOT NULL DEFAULT '',
  `visitor_batter1_id` varchar(8) NOT NULL DEFAULT '',
  `visitor_batter1_name` varchar(32) NOT NULL DEFAULT '',
  `visitor_batter1_pos` int(11) NOT NULL DEFAULT '0',
  `visitor_batter2_id` varchar(8) NOT NULL DEFAULT '',
  `visitor_batter2_name` varchar(32) NOT NULL DEFAULT '',
  `visitor_batter2_pos` int(11) NOT NULL DEFAULT '0',
  `visitor_batter3_id` varchar(8) NOT NULL DEFAULT '',
  `visitor_batter3_name` varchar(32) NOT NULL DEFAULT '',
  `visitor_batter3_pos` int(11) NOT NULL DEFAULT '0',
  `visitor_batter4_id` varchar(8) NOT NULL DEFAULT '',
  `visitor_batter4_name` varchar(32) NOT NULL DEFAULT '',
  `visitor_batter4_pos` int(11) NOT NULL DEFAULT '0',
  `visitor_batter5_id` varchar(8) NOT NULL DEFAULT '',
  `visitor_batter5_name` varchar(32) NOT NULL DEFAULT '',
  `visitor_batter5_pos` int(11) NOT NULL DEFAULT '0',
  `visitor_batter6_id` varchar(8) NOT NULL DEFAULT '',
  `visitor_batter6_name` varchar(32) NOT NULL DEFAULT '',
  `visitor_batter6_pos` int(11) NOT NULL DEFAULT '0',
  `visitor_batter7_id` varchar(8) NOT NULL DEFAULT '',
  `visitor_batter7_name` varchar(32) NOT NULL DEFAULT '',
  `visitor_batter7_pos` int(11) NOT NULL DEFAULT '0',
  `visitor_batter8_id` varchar(8) NOT NULL DEFAULT '',
  `visitor_batter8_name` varchar(32) NOT NULL DEFAULT '',
  `visitor_batter8_pos` int(11) NOT NULL DEFAULT '0',
  `visitor_batter9_id` varchar(8) NOT NULL DEFAULT '',
  `visitor_batter9_name` varchar(32) NOT NULL DEFAULT '',
  `visitor_batter9_pos` int(11) NOT NULL DEFAULT '0',
  `home_batter1_id` varchar(8) NOT NULL DEFAULT '',
  `home_batter1_name` varchar(32) NOT NULL DEFAULT '',
  `home_batter1_pos` int(11) NOT NULL DEFAULT '0',
  `home_batter2_id` varchar(8) NOT NULL DEFAULT '',
  `home_batter2_name` varchar(32) NOT NULL DEFAULT '',
  `home_batter2_pos` int(11) NOT NULL DEFAULT '0',
  `home_batter3_id` varchar(8) NOT NULL DEFAULT '',
  `home_batter3_name` varchar(32) NOT NULL DEFAULT '',
  `home_batter3_pos` int(11) NOT NULL DEFAULT '0',
  `home_batter4_id` varchar(8) NOT NULL DEFAULT '',
  `home_batter4_name` varchar(32) NOT NULL DEFAULT '',
  `home_batter4_pos` int(11) NOT NULL DEFAULT '0',
  `home_batter5_id` varchar(8) NOT NULL DEFAULT '',
  `home_batter5_name` varchar(32) NOT NULL DEFAULT '',
  `home_batter5_pos` int(11) NOT NULL DEFAULT '0',
  `home_batter6_id` varchar(8) NOT NULL DEFAULT '',
  `home_batter6_name` varchar(32) NOT NULL DEFAULT '',
  `home_batter6_pos` int(11) NOT NULL DEFAULT '0',
  `home_batter7_id` varchar(8) NOT NULL DEFAULT '',
  `home_batter7_name` varchar(32) NOT NULL DEFAULT '',
  `home_batter7_pos` int(11) NOT NULL DEFAULT '0',
  `home_batter8_id` varchar(8) NOT NULL DEFAULT '',
  `home_batter8_name` varchar(32) NOT NULL DEFAULT '',
  `home_batter8_pos` int(11) NOT NULL DEFAULT '0',
  `home_batter9_id` varchar(8) NOT NULL DEFAULT '',
  `home_batter9_name` varchar(32) NOT NULL DEFAULT '',
  `home_batter9_pos` int(11) NOT NULL DEFAULT '0',
  `additional_information` varchar(128) NOT NULL DEFAULT '',
  `acquisition_information` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`season`,`game_date_str`,`game_number`,`visitors_team`,`home_team`,`visitors_season_game_number`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `schedules`
--

DROP TABLE IF EXISTS `schedules`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schedules` (
  `season` int(11) NOT NULL,
  `game_date_str` varchar(8) NOT NULL,
  `game_number` int(11) NOT NULL,
  `visitors_team` varchar(3) NOT NULL,
  `home_team` varchar(3) NOT NULL,
  `game_month` int(11) NOT NULL,
  `game_month_day` int(11) NOT NULL,
  `week_day` int(11) NOT NULL,
  `visitors_league` varchar(3) NOT NULL,
  `visitors_season_game_number` int(11) NOT NULL,
  `home_league` varchar(3) NOT NULL,
  `home_season_game_number` int(11) NOT NULL,
  `game_type` varchar(1) NOT NULL,
  `postponed_canceled` varchar(128) NOT NULL DEFAULT '',
  `make_up_date` varchar(128) NOT NULL DEFAULT '',
  PRIMARY KEY (`season`,`game_date_str`,`game_number`,`visitors_team`,`home_team`,`visitors_season_game_number`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-04-29 19:31:43
